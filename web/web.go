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
	gcontext "github.com/gorilla/context"
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/auth/eveonline/callback", h.EveOauthCallbackHandler)
	r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		// Get a session. Get() always returns a session, even if empty.
		session, err := h.store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.RedirectToSSO(session, w, r)
	})

	r.Get("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		// Get a session. Get() always returns a session, even if empty.
		session, err := h.store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Options.MaxAge = -1

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		http.Redirect(w, r, "/", 302)
	})

	r.Route("/u", func(r chi.Router) {
		r.Use(h.AuthenticationRequirement)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// Get a session. Get() always returns a session, even if empty.
			session, err := h.store.Get(r, "session")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			val, _ := session.Values["user_id"].(string)

			var character model.UserCharacter
			err = h.clients.DB.Get(&character, "select * from user_characters where user_id = $1 limit 1", val)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			refreshToken, err := goesi.TokenFromJSON(character.OauthToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tokSrc, err := h.clients.ESIAuthenticator.TokenSource(refreshToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(context.TODO(), goesi.ContextOAuth2, tokSrc)
			data, _, err := h.clients.ESIClient.ClonesApi.GetCharactersCharacterIdImplants(ctx, int32(character.ID), nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write([]byte(fmt.Sprintf("welcome %s: %+v", val, data)))
		})
	})

	http.ListenAndServe(":"+port, gcontext.ClearHandler(r))
}
