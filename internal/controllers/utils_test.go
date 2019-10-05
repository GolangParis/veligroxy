package controllers

import (
	"testing"
)

func TestBuildQueryUrl(t *testing.T) {
	queryParams := map[string][]string{
		"dataset":            []string{"velib-disponibilite-en-temps-reel"},
		"facet":              []string{"overflowactivation", "creditcard", "kioskstate", "station_state"},
		"geofilter.distance": []string{"1,2,3"},
	}

	response := buildQueryParams("?", queryParams)

	expected := "?dataset=velib-disponibilite-en-temps-reel&facet=overflowactivation&facet=creditcard&facet=kioskstate&facet=station_state&geofilter.distance=1,2,3"

	if response != expected {
		t.Errorf("Query build was incorrect, got: <%s>, want: <%s>", response, expected)
	}
}
