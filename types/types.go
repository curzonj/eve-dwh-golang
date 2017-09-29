package types

import (
	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/antihax/goesi/esi"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type contextKey int

const (
	ContextSessionKey contextKey = iota // 0
	ContextLoggerKey                    // 1
)

type Clients struct {
	ESIClient        *esi.APIClient
	ESIAuthenticator *goesi.SSOAuthenticator
	DB               *sqlx.DB
	Redis            *redis.Pool
	Logger           log.FieldLogger
}
