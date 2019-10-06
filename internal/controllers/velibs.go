package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GolangParis/veligroxy/internal/models"
	"github.com/GolangParis/veligroxy/internal/services"
	log "github.com/sirupsen/logrus"
)

func QueryVelibStatus() http.HandlerFunc {
	point := models.SearchPoint{Lat: "48.853169", Long: "2.402782", Radius: "100"}

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := services.GetVelibStatus(point)
		if err != nil {
			log.Warn("Error reading velib status")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info("Reading velib status was OK")

		response, _ := json.Marshal(&payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
