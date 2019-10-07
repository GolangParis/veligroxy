package routes

import (
	"github.com/GolangParis/veligroxy/internal/controllers"
	"github.com/GolangParis/veligroxy/internal/diagnostics"
	"github.com/gorilla/mux"
)

// BusinessRoutes returns a new router with endpoints providing a business functionality
func BusinessRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/velibs", controllers.QueryVelibStatus()).Methods("GET")
	return r
}

// DiagnosticsRoutes returns a new router with the technical
func DiagnosticsRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", diagnostics.Health()).Methods("GET")
	r.HandleFunc("/ready", diagnostics.Ready()).Methods("GET")
	return r
}
