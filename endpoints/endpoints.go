package endpoints

import (
	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/dashboard"
	"github.com/gorilla/mux"
)

// Router gorilla router
var Router = mux.NewRouter()

// SetupRoutes setup the all routes
func SetupRoutes() {
	Router.HandleFunc("/", dashboard.HandleHome)
	Router.HandleFunc("/login", auth.HandleLogin)
	Router.HandleFunc("/logout", auth.HandleLogout)
	Router.HandleFunc("/register", auth.HandleRegister)
	Router.HandleFunc("/profile", dashboard.HandleProfile)
}
