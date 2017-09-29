package web

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
)

type Cfg struct {
	Port   string `env:"PORT,required"`
	Secret string `env:"SECRET,default=totally-insecure"`
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

	h.run(cfg.Port)
}

func (h *handler) run(port string) {
	logger := h.clients.Logger.WithField("fn", "runWebHandler")
	logger.WithFields(log.Fields{
		"at":   "start",
		"port": port,
	}).Info()

	r := chi.NewRouter()
	r.Use(h.logRequest)
	r.Use(h.buildSession)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Get("/auth/eveonline/callback", wrapErrors(h.eveOauthCallback))
	r.Get("/auth/login", wrapErrors(h.redirectToSSO))
	r.Get("/auth/logout", wrapErrors(h.logoutSession))

	r.Route("/u", func(r chi.Router) {
		r.Use(h.authenticationRequirement)

		r.Get("/", wrapErrors(func(w http.ResponseWriter, r *http.Request) error {
			// Get a session. Get() always returns a session, even if empty.
			session := session(r)
			userID := session.Values["user_id"].(string)

			var character model.UserCharacter
			err := h.clients.DB.Get(&character, "select * from user_characters where user_id = $1 limit 1", userID)
			if err != nil {
				return err
			}

			tokSrc, err := character.TokenSource(h.clients.ESIAuthenticator)
			if err != nil {
				return err
			}

			ctx := context.WithValue(r.Context(), goesi.ContextOAuth2, tokSrc)
			data, _, err := h.clients.ESIClient.ClonesApi.GetCharactersCharacterIdImplants(ctx, int32(character.ID), nil)
			if err != nil {
				return err
			}

			w.Write([]byte(fmt.Sprintf("welcome %s: %+v", userID, data)))
			return nil
		}))
	})

	http.ListenAndServe(":"+port, r)
}
