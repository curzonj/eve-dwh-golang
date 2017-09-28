package web

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

func (h *handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		finished := time.Now()

		h.clients.Logger.WithFields(log.Fields{
			"at":      "httpRequest",
			"method":  r.Method,
			"elapsed": finished.Sub(start).Seconds(),
			"path":    r.URL.RequestURI(),
		}).Info()
	})
}
