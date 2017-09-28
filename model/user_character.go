package model

import "github.com/antihax/goesi"

type UserCharacter struct {
	UserID      string `db:"user_id"`
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	OwnerHash   string `db:"owner_hash"`
	OauthScopes string `db:"oauth_scopes"`
	OauthToken  string `db:"oauth_token"`
}

func (u *UserCharacter) TokenSource(a *goesi.SSOAuthenticator) (goesi.CRESTTokenSource, error) {
	refreshToken, err := goesi.TokenFromJSON(u.OauthToken)
	if err != nil {
		return nil, err
	}

	return a.TokenSource(refreshToken)
}
