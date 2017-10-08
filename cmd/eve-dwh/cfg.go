package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/curzonj/eve-dwh-golang/poller"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/curzonj/eve-dwh-golang/web"
	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

var cfg struct {
	Poller poller.Cfg
	Web    web.Cfg

	DatabaseURL string `env:"DATABASE_URL,required"`

	ESI struct {
		UserAgent        string `env:"USER_AGENT,required"`
		OauthClientID    string `env:"OAUTH_CLIENT_ID,required"`
		OauthSecretKey   string `env:"OAUTH_SECRET_KEY,required"`
		OauthRedirectURL string `env:"OAUTH_REDIRECT_URL,required"`
		ScopesString     string `env:"ESI_SCOPES,required"`
	}
}

var clients types.Clients

func connectToDatabase() {
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		// TODO could add something like https://github.com/heroku/x/blob/master/hredis/redigo/redigo.go#L15
		log.Fatal(err)
	}

	clients.DB = db
}

func buildClients() {
	connectToDatabase()
	buildESIClient()
}

func loadEnvironment() {
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	clients.Logger = log.New()
}
