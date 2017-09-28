package model

type UserCharacter struct {
	UserID      string `db:"user_id"`
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	OwnerHash   string `db:"owner_hash"`
	OauthScopes string `db:"oauth_scopes"`
	OauthToken  string `db:"oauth_token"`
}
