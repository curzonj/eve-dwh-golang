package web

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/antihax/goesi"
	"github.com/curzonj/eve-dwh-golang/model"
	"github.com/pborman/uuid"
)

func exchangeToken(a *goesi.SSOAuthenticator, r *http.Request) (*goesi.VerifyResponse, string, error) {
	session := session(r)

	// get our code and state
	code := r.FormValue("code")
	state := r.FormValue("state")
	sessionState := session.Values["state"]

	delete(session.Values, "state")

	// Verify the state matches our randomly generated string from earlier.
	if sessionState != state {
		return nil, "", errors.New("Invalid state")
	}

	// Exchange the code for an Access and Refresh token.
	token, err := a.TokenExchange(code)
	if err != nil {
		return nil, "", err
	}

	// Obtain a token source (automaticlly pulls refresh as needed)
	tokSrc, err := a.TokenSource(token)
	if err != nil {
		return nil, "", err
	}

	// Verify the client (returns clientID)
	characterInfo, err := a.Verify(tokSrc)
	if err != nil {
		return nil, "", err
	}

	jsonToken, err := goesi.TokenToJSON(token)
	if err != nil {
		return nil, "", err
	}

	return characterInfo, jsonToken, nil
}

func (h *handler) eveOauthCallback(w http.ResponseWriter, r *http.Request) error {
	session := session(r)

	characterInfo, jsonToken, err := exchangeToken(h.clients.ESIAuthenticator, r)
	if err != nil {
		return err
	}

	var character model.UserCharacter
	characterExists := true

	err = h.clients.DB.Get(&character, "select * from user_characters where id = $1 limit 1", characterInfo.CharacterID)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}

		characterExists = false
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok {
		if characterExists {
			userID = character.UserID
		} else {
			userID = uuid.NewUUID().String()
			_, err := h.clients.DB.Exec("insert into users (id) values ($1)", userID)
			if err != nil {
				return err
			}
		}

		session.Values["user_id"] = userID
	} else {
		if characterExists {
			if character.UserID != userID {
				http.Error(w, "Another users owns that character", http.StatusUnauthorized)
				return nil
			}

			// TODO update the available scopes stored in the db for this character
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
			return err
		}
	}

	redirectTo := session.Values["redirect_to"].(string)
	if redirectTo == "" {
		redirectTo = "/"
	}

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	http.Redirect(w, r, redirectTo, 302)
	return nil
}

func (h *handler) redirectToSSO(w http.ResponseWriter, r *http.Request) error {
	session := session(r)
	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Save the state on the session
	session.Values["state"] = state
	err := session.Save(r, w)
	if err != nil {
		return err
	}

	// Generate the SSO URL with the state string
	url := h.clients.ESIAuthenticator.AuthorizeURL(state, false, nil)

	// Send the user to the URL
	http.Redirect(w, r, url, 302)
	return nil
}

func (h *handler) logoutSession(w http.ResponseWriter, r *http.Request) error {
	// Get a session. Get() always returns a session, even if empty.
	session := session(r)
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/", 302)
	return nil
}

func (h *handler) authenticationRequirement(next http.Handler) http.Handler {
	return http.HandlerFunc(wrapErrors(func(w http.ResponseWriter, r *http.Request) error {
		session := session(r)
		_, ok := session.Values["user_id"].(string)

		if !ok {
			session.Values["redirect_to"] = r.URL.RequestURI()
			err := session.Save(r, w)
			if err != nil {
				return err
			}

			h.redirectToSSO(w, r)
			return nil
		}

		next.ServeHTTP(w, r)
		return nil
	}))
}
