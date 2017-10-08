package web

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/a-h/hsts"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
)

type Cfg struct {
	Port        string `env:"PORT,required"`
	Secret      string `env:"SECRET,required"`
	SSLRequired bool   `env:"SSL_REQUIRED,default=true"`
}

type handler struct {
	clients types.Clients
	store   sessions.Store
}

func RunHandler(clients types.Clients, cfg Cfg) {
	h := &handler{
		clients: clients,
		store:   sessions.NewCookieStore([]byte(cfg.Secret)),
	}

	h.run(cfg)
}

func (h *handler) run(cfg Cfg) {
	logger := h.clients.Logger.WithField("fn", "runWebHandler")
	logger.WithFields(log.Fields{
		"at":   "start",
		"port": cfg.Port,
	}).Info()

	r := chi.NewRouter()
	r.Use(h.logRequest)
	r.Use(h.buildSession)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/auth/eveonline/callback", wrapErrors(h.eveOauthCallback))
	r.Get("/auth/login", wrapErrors(h.redirectToSSO))
	r.Get("/auth/logout", wrapErrors(h.logoutSession))

	r.Route("/", func(r chi.Router) {
		r.Use(h.authenticationRequirement)

		r.Get("/industry", wrapErrors(h.industryJobs))

		r.Mount("/", http.FileServer(http.Dir("public")))
	})

	handler := http.Handler(r)

	if cfg.SSLRequired {
		logger.Info("SSL enforcement enabled")
		handler = hsts.NewHandler(handler)
	}

	http.ListenAndServe(":"+cfg.Port, handler)
}
