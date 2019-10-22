package endpoints

import (
	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/gorilla/mux"
)

// Router gorilla router
var Router = mux.NewRouter()

// SetupRoutes setup the all routes
func SetupRoutes() {
	Router.HandleFunc("/", auth.HandleLogin)
}
