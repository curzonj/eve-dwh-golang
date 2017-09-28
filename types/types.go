package types

import (
	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/jmoiron/sqlx"
)

type Clients struct {
	ESIClient        *esi.APIClient
	ESIAuthenticator *goesi.SSOAuthenticator
	DB               *sqlx.DB
	Logger           log.FieldLogger
}
