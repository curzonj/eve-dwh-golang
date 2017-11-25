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

		requestID := uuid.NewUUID().String()

		l := h.clients.Logger.WithFields(log.Fields{
			"req": requestID,
		})

		c := r.Context()
		c = context.WithValue(c, types.ContextLoggerKey, l)
		r = r.WithContext(c)

		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		l.WithFields(log.Fields{
			"at":      "httpRequest",
			"method":  r.Method,
			"path":    r.URL.RequestURI(),
			"status":  lrw.statusCode,
			"elapsed": time.Now().Sub(start).Seconds(),
		}).Info()
	})
}

func logger(ctx context.Context) log.FieldLogger {
	return ctx.Value(types.ContextLoggerKey).(log.FieldLogger)
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
