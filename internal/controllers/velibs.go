package controllers

import (
	"encoding/json"
	_ "encoding/json"
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

	return func(w http.ResponseWriter, r *http.Request) {
		route := "https://opendata.paris.fr/api/records/1.0/search"

		// Assemblage route + query string avec encodage des caractères spéciaux
		params := fmt.Sprintf("?dataset=velib-disponibilite-en-temps-reel&facet=overflowactivation&facet=creditcard&facet=kioskstate&facet=station_state&geofilter.distance=%s,%s,%s",
			lat, long, radius)
		url := fmt.Sprintf("%s/%s", route, html.EscapeString(params))

		payload := models.VelibStatus{}
		err := getJson(url, &payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info("Reading velib status was OK")

		response, _ := json.Marshal(&payload)
		w.Write(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
