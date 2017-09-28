package web

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/gorilla/sessions"
	"github.com/pborman/uuid"
)

func (h *handler) EveOauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := h.store.Get(r, "session")
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
	token, err := h.clients.ESIAuthenticator.TokenExchange(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc, err := h.clients.ESIAuthenticator.TokenSource(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verify the client (returns clientID)
	characterInfo, err := h.clients.ESIAuthenticator.Verify(tokSrc)
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
	var character model.UserCharacter
	characterExists := true

	err = h.clients.DB.Get(&character, "select * from user_characters where id = $1", characterInfo.CharacterID)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		characterExists = false
	}

	val := session.Values["user_id"]
	if passedUserID, ok := val.(string); ok && passedUserID != "" {
		err = h.clients.DB.Get(&userID, "select id from users where id = $1", passedUserID)
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
			_, err := h.clients.DB.Exec("insert into users (id) values ($1)", userID)
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
				h.clients.Logger.Error("user_id mismatch at login")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if !characterExists {
		character = model.UserCharacter{
			UserID:      userID,
			ID:          characterInfo.CharacterID,
			Name:        characterInfo.CharacterName,
			OwnerHash:   characterInfo.CharacterOwnerHash,
			OauthScopes: characterInfo.Scopes,
			OauthToken:  jsonToken,
		}

		_, err = h.clients.DB.NamedExec("insert into user_characters (user_id, id, name, owner_hash, oauth_scopes, oauth_token) values (:user_id, :id, :name, :owner_hash, :oauth_scopes, :oauth_token)", &character)
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

func (h *handler) RedirectToSSO(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
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
	url := h.clients.ESIAuthenticator.AuthorizeURL(state, true, h.clients.ESIScopes)

	// Send the user to the URL
	http.Redirect(w, r, url, 302)
}

func (h *handler) AuthenticationRequirement(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := h.store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var userID string

		val := session.Values["user_id"]
		if passedUserID, ok := val.(string); ok && passedUserID != "" {
			err = h.clients.DB.Get(&userID, "select id from users where id = $1", passedUserID)
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

			h.RedirectToSSO(session, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
