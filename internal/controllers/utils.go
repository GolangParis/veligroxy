package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func buildQueryParams(prefix string, params map[string][]string) string {
	var b strings.Builder
	b.WriteString(prefix)

	for k, v := range params {
		for _, element := range v {
			b.WriteString(fmt.Sprintf("%s=%s&", k, element))
		}
	}

	return strings.TrimSuffix(b.String(), "&")
}
