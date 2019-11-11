package endpoints

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/admin"
	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/dashboard"
	"github.com/dilip640/Faculty-Portal/personalcv"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Router gorilla router
var Router = mux.NewRouter()

func wrapHandlerWithLogging(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("--> %s %s", req.Method, req.URL.Path)
		f(w, req)
	})
}

// SetupRoutes setup the all routes
func SetupRoutes() {
	Router.HandleFunc("/", wrapHandlerWithLogging(dashboard.HandleHome))
	Router.HandleFunc("/login", wrapHandlerWithLogging(auth.HandleLogin))
	Router.HandleFunc("/logout", wrapHandlerWithLogging(auth.HandleLogout))
	Router.HandleFunc("/register", wrapHandlerWithLogging(auth.HandleRegister))
	Router.HandleFunc("/profile", wrapHandlerWithLogging(dashboard.HandleProfile))
	Router.HandleFunc("/profile/{id}", wrapHandlerWithLogging(dashboard.HandleProfile))
	Router.HandleFunc("/faculty/register",
		wrapHandlerWithLogging(dashboard.HandleRegisterFaculty))
	Router.HandleFunc("/ccfaculty/register",
		wrapHandlerWithLogging(dashboard.HandleRegisterCCFaculty))
	Router.HandleFunc("/faculty/editcv",
		wrapHandlerWithLogging(personalcv.HandleCVEdit))

	// Admin routes
	Router.HandleFunc("/admin",
		wrapHandlerWithLogging(admin.HandleAdmin))
}
