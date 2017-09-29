package web

import (
	"context"
	"net/http"

	"github.com/curzonj/eve-dwh-golang/types"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func (h *handler) buildSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get a session. Get() always returns a session, even if empty.
		session, err := h.store.Get(r, "session")
		defer gcontext.Clear(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c := r.Context()
		c = context.WithValue(c, types.ContextSessionKey, session)
		r = r.WithContext(c)

		next.ServeHTTP(w, r)
	})
}

func session(r *http.Request) *sessions.Session {
	return r.Context().Value(types.ContextSessionKey).(*sessions.Session)
}
