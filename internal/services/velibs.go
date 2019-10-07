package services

import (
	"fmt"
	"html"

	"github.com/GolangParis/veligroxy/internal/models"
)

// GetVelibStatus query the external opendata endpoint
func GetVelibStatus(point models.SearchPoint) (models.VelibStatus, error) {
	route := "https://opendata.paris.fr/api/records/1.0/search"

	params := map[string][]string{
		"dataset":            []string{"velib-disponibilite-en-temps-reel"},
		"facet":              []string{"overflowactivation", "creditcard", "kioskstate", "station_state"},
		"geofilter.distance": []string{point.ToString()},
	}

	queryParams := BuildQueryParams("?", params)

	url := fmt.Sprintf("%s/%s", route, html.EscapeString(queryParams))

	payload := models.VelibStatus{}
	err := GetJson(url, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
