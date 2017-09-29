package web

import (
	"context"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/curzonj/eve-dwh-golang/types"
	"github.com/pborman/uuid"
)

func (h *handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		l := h.clients.Logger.WithField("req", uuid.NewUUID().String())

		c := r.Context()
		c = context.WithValue(c, types.ContextLoggerKey, l)
		r = r.WithContext(c)

		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		finished := time.Now()

		l.WithFields(log.Fields{
			"at":      "httpRequest",
			"method":  r.Method,
			"status":  lrw.statusCode,
			"elapsed": finished.Sub(start).Seconds(),
			"path":    r.URL.RequestURI(),
		}).Info()
	})
}

func logger(r *http.Request) log.FieldLogger {
	return r.Context().Value(types.ContextLoggerKey).(log.FieldLogger)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
