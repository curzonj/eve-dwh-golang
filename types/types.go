package types

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/jmoiron/sqlx"
)

type contextKey int

const (
	ContextSessionKey contextKey = iota // 0
	ContextLoggerKey                    // 1
)

type Clients struct {
	HTTPBreakerClient *http.Client
	EVEBreakerClient  *goesi.APIClient
	EVERetryClient    *goesi.APIClient
	ESIAuthenticator  *goesi.SSOAuthenticator
	DB                *sqlx.DB
	Logger            log.FieldLogger
}
