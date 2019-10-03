package routes

import (
	"github.com/fagossa/golang-rest-api/internal/controllers"
	"github.com/fagossa/golang-rest-api/internal/diagnostics"
	"github.com/gorilla/mux"
)

func BusinessRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/velibs", controllers.ReadVelib()).Methods("GET")
	return r
}

func DiagnosticsRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", diagnostics.Health()).Methods("GET")
	r.HandleFunc("/ready", diagnostics.Ready()).Methods("GET")
	return r
}
