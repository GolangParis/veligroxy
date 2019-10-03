package controllers

import (
	"fmt"
	"html"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ReadVelib() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lat := "48.853169"
		long := "2.402782"
		radius := "100"

		route := "https://opendata.paris.fr/api/records/1.0/search"

		// Assemblage route + query string avec encodage des caractères spéciaux
		params := fmt.Sprintf("?dataset=velib-disponibilite-en-temps-reel&facet=overflowactivation&facet=creditcard&facet=kioskstate&facet=station_state&geofilter.distance=%s,%s,%s",
			lat, long, radius)
		url := fmt.Sprintf("%s/%s", route, html.EscapeString(params))

		_, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.WithFields(log.Fields{
			//"id": id,
		}).Info("Reading velib")

		w.WriteHeader(http.StatusOK)
	}
}
