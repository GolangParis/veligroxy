package controllers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ReadVelib() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// id := vars["id"]

		log.WithFields(log.Fields{
			//"id": id,
		}).Info("Reading velib")
	}
}
