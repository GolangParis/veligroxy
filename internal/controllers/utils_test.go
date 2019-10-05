package controllers

import (
	"reflect"
	"strings"
	"testing"
)

func TestBuildQueryUrlWithElements(t *testing.T) {
	// Given
	queryParams := map[string][]string{
		"dataset":            []string{"velib-disponibilite-en-temps-reel"},
		"facet":              []string{"overflowactivation", "creditcard", "kioskstate", "station_state"},
		"geofilter.distance": []string{"1,2,3"},
	}

	// When
	response := buildQueryParams("?", queryParams)
	responseTokens := strings.Split(response, "&")

	// Then
	expected := "?dataset=velib-disponibilite-en-temps-reel&facet=overflowactivation&facet=creditcard&facet=kioskstate&facet=station_state&geofilter.distance=1,2,3"
	expectedTokens := strings.Split(expected, "&")

	if reflect.DeepEqual(responseTokens, expectedTokens) != true {
		t.Errorf("Query build was incorrect, got: <%s>, want: <%s>", responseTokens, expectedTokens)
	}
}

func TestBuildQueryUrlWhenEmpty(t *testing.T) {
	// Given
	emptyQueryParams := map[string][]string{}

	// When
	response := buildQueryParams("?", emptyQueryParams)
	responseTokens := strings.Split(response, "&")

	// Then
	expectedTokens := []string{"?"}

	if reflect.DeepEqual(responseTokens, expectedTokens) != true {
		t.Errorf("Query build was incorrect, got: <%s>, want: <%s>", responseTokens, expectedTokens)
	}
}
