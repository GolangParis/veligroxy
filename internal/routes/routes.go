package routes

import (
	"github.com/GolangParis/veligroxy/internal/controllers"
	"github.com/GolangParis/veligroxy/internal/diagnostics"
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
