package model

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/antihax/goesi"
	"golang.org/x/oauth2"
)

type UserCharacter struct {
	UserID         string         `db:"user_id"`
	ID             int64          `db:"id"`
	Name           string         `db:"name"`
	OwnerHash      string         `db:"owner_hash"`
	OauthScopes    string         `db:"oauth_scopes"`
	OauthToken     string         `db:"oauth_token"`
	EVEAccountName sql.NullString `db:"eve_account"`
	Flags          sql.NullString `db:"flags"`
	XMLKeyID       sql.NullInt64  `db:"xml_key_id"`
	XMLVCode       sql.NullString `db:"xml_vcode"`

	db *DAO
}

func (d *DAO) GetAllCharacters() ([]*UserCharacter, error) {
	var characters []*UserCharacter
	err := d.Select(&characters, "select * from user_characters")
	if err != nil {
		return nil, err
	}

	for _, c := range characters {
		c.db = d
	}

	return characters, nil
}

func (d *DAO) GetUserCharacterByID(id int64) (*UserCharacter, error) {
	var c UserCharacter
	err := d.Get(&c, "select * from user_characters where id = $1 limit 1", id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		return nil, nil
	}

	c.db = d
	return &c, nil
}

func (d *DAO) GetUserCharactersByUserID(id string) ([]*UserCharacter, error) {
	var characters []*UserCharacter
	err := d.Select(&characters, "select * from user_characters where user_id = $1", id)
	if err != nil {
		return nil, err
	}

	for _, c := range characters {
		c.db = d
	}

	return characters, nil
}

func (u *UserCharacter) HasFlag(s string) bool {
	/*	for _, f := range u.Flags {
		if f == s {
			return true
		}
	} */

	return false
}

func (u *UserCharacter) TokenSource(a *goesi.SSOAuthenticator) (oauth2.TokenSource, error) {
	var storedToken oauth2.Token
	if err := json.Unmarshal([]byte(u.OauthToken), &storedToken); err != nil {
		return nil, err
	}

	tokSrc, err := a.TokenSource((*goesi.CRESTToken)(&storedToken))
	if err != nil {
		return nil, err
	}

	if !storedToken.Valid() {
		newToken, err := tokSrc.Token()
		if err != nil {
			return nil, err
		}

		if newToken.Valid() {
			bytes, err := json.Marshal(newToken)
			if err != nil {
				return nil, err
			}

			_, err = u.db.Exec("update user_characters set oauth_token = $2 where id = $1", u.ID, string(bytes))
			if err != nil {
				return nil, err
			}
		}
	}

	return tokSrc, nil
}

func (u *UserCharacter) TokenSourceContext(ctx context.Context, a *goesi.SSOAuthenticator) (context.Context, error) {
	tokSrc, err := u.TokenSource(a)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, goesi.ContextOAuth2, tokSrc)
	return ctx, nil
}
