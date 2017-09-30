package model

import (
	"encoding/json"

	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/types"
	"golang.org/x/oauth2"
)

type UserCharacter struct {
	UserID      string `db:"user_id"`
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	OwnerHash   string `db:"owner_hash"`
	OauthScopes string `db:"oauth_scopes"`
	OauthToken  string `db:"oauth_token"`
}

func (u *UserCharacter) TokenSource(c types.Clients) (goesi.CRESTTokenSource, error) {
	a := c.ESIAuthenticator

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

			_, err = c.DB.Exec("update user_characters set oauth_token = $2 where id = $1", u.ID, string(bytes))
			if err != nil {
				return nil, err
			}
		}
	}

	return tokSrc, nil
}
