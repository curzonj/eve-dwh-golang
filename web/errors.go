package web

import "net/http"

func wrapErrors(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			logger(r.Context()).Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
