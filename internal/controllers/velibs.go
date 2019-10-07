package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GolangParis/veligroxy/internal/models"
	"github.com/GolangParis/veligroxy/internal/services"
	log "github.com/sirupsen/logrus"
)

// QueryVelibStatus provide the endpoint to Velib status
func QueryVelibStatus() http.HandlerFunc {
	point := models.SearchPoint{
		Lat:    "48.853169",
		Long:   "2.402782",
		Radius: "100"}

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := services.GetVelibStatus(point)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Warn("Error reading velib status")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stations := payload.ExtractStations()

		log.WithFields(log.Fields{
			"Lat":     point.Lat,
			"Long":    point.Long,
			"Radius":  point.Radius,
			"Results": len(payload.Records),
		}).Info("Success velib status was OK")

		response, _ := json.Marshal(&stations)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
