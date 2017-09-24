package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/antihax/goesi"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"github.com/pborman/uuid"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func RedirectToSSO(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Save the state on the session
	session.Values["state"] = state
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate the SSO URL with the state string
	url := globals.esiAuthenticator.AuthorizeURL(state, true, cfg.ESIScopes)

	// Send the user to the URL
	http.Redirect(w, r, url, 302)
}

func AuthenticationRequirement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var userID string

		val := session.Values["user_id"]
		if passedUserID, ok := val.(string); ok && passedUserID != "" {
			err = globals.db.Get(&userID, "select id from users where id = $1", passedUserID)
			if err != nil {
				if err == sql.ErrNoRows || strings.HasPrefix(err.Error(), "pq: invalid input syntax for uuid:") {
					session.Values["user_id"] = ""
					err = session.Save(r, w)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				} else {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		if userID == "" {
			session.Values["redirect_to"] = r.URL.RequestURI()
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			RedirectToSSO(session, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type UserCharacter struct {
	UserID      string `db:"user_id"`
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	OwnerHash   string `db:"owner_hash"`
	OauthScopes string `db:"oauth_scopes"`
	OauthToken  string `db:"oauth_token"`
}

func EveOauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get our code and state
	code := r.FormValue("code")
	state := r.FormValue("state")

	// Verify the state matches our randomly generated string from earlier.
	if session.Values["state"] != state {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exchange the code for an Access and Refresh token.
	token, err := globals.esiAuthenticator.TokenExchange(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc, err := globals.esiAuthenticator.TokenSource(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verify the client (returns clientID)
	characterInfo, err := globals.esiAuthenticator.Verify(tokSrc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonToken, err := goesi.TokenToJSON(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userID string
	var character UserCharacter
	characterExists := true

	err = globals.db.Get(&character, "select * from user_characters where id = $1", characterInfo.CharacterID)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		characterExists = false
	}

	val := session.Values["user_id"]
	if passedUserID, ok := val.(string); ok && passedUserID != "" {
		err = globals.db.Get(&userID, "select id from users where id = $1", passedUserID)
		if err != nil {
			if err == sql.ErrNoRows || strings.HasPrefix(err.Error(), "pq: invalid input syntax for uuid:") {
				session.Values["user_id"] = ""
				err = session.Save(r, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if userID == "" {
		if characterExists {
			userID = character.UserID
		} else {
			userID = uuid.NewUUID().String()
			_, err := globals.db.Exec("insert into users (id) values ($1)", userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		session.Values["user_id"] = userID
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if characterExists {
			// TODO update scopes
			if character.UserID != userID {
				globals.logger.Error("user_id mismatch at login")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if !characterExists {
		character = UserCharacter{
			UserID:      userID,
			ID:          characterInfo.CharacterID,
			Name:        characterInfo.CharacterName,
			OwnerHash:   characterInfo.CharacterOwnerHash,
			OauthScopes: characterInfo.Scopes,
			OauthToken:  jsonToken,
		}

		_, err = globals.db.NamedExec("insert into user_characters (user_id, id, name, owner_hash, oauth_scopes, oauth_token) values (:user_id, :id, :name, :owner_hash, :oauth_scopes, :oauth_token)", &character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	redirectTo := session.Values["redirect_to"].(string)
	if redirectTo == "" {
		redirectTo = "/"
	}

	http.Redirect(w, r, redirectTo, 302)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		finished := time.Now()

		globals.logger.WithFields(log.Fields{
			"at":      "httpRequest",
			"method":  r.Method,
			"elapsed": finished.Sub(start).Seconds(),
			"path":    r.URL.RequestURI(),
		}).Info()
	})
}

func runWebHandler() {
	logger := globals.logger.WithField("fn", "runWebHandler")
	logger.WithFields(log.Fields{
		"at":   "start",
		"port": cfg.Port,
	}).Info()

	r := chi.NewRouter()
	r.Use(logRequest)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/auth/eveonline/callback", EveOauthCallbackHandler)
	r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		// Get a session. Get() always returns a session, even if empty.
		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		RedirectToSSO(session, w, r)
	})

	r.Get("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		// Get a session. Get() always returns a session, even if empty.
		session, err := store.Get(r, "session")
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
		r.Use(AuthenticationRequirement)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// Get a session. Get() always returns a session, even if empty.
			session, err := store.Get(r, "session")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			val, _ := session.Values["user_id"].(string)
			w.Write([]byte("welcome " + val))
		})
	})

	http.ListenAndServe(":"+cfg.Port, r)
}
