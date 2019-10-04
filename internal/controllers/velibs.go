package controllers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/GolangParis/veligroxy/internal/models"
	log "github.com/sirupsen/logrus"
)

func ReadVelib() http.HandlerFunc {
	lat := "48.853169"
	long := "2.402782"
	radius := "100"

	queryParams := map[string][]string{
		"dataset":            []string{"velib-disponibilite-en-temps-reel"},
		"facet":              []string{"overflowactivation", "creditcard", "kioskstate", "station_state"},
		"geofilter.distance": []string{fmt.Sprintf("%s,%s,%s", lat, long, radius)},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		route := "https://opendata.paris.fr/api/records/1.0/search"

		queryParams := buildQueryParams("?", queryParams)

		url := fmt.Sprintf("%s/%s", route, html.EscapeString(queryParams))

		payload := models.VelibStatus{}
		err := getJson(url, &payload)
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
