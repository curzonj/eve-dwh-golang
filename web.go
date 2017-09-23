package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
)

func runWebHandler() {
	logger := globals.logger.WithField("fn", "runWebHandler")
	logger.WithFields(log.Fields{
		"at":   "start",
		"port": cfg.Port,
	}).Info()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":"+cfg.Port, r)
}
