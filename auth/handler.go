package auth

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/templatemanager"
)

// HandleLogin handle the greeeting
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if GetUserName(r) != "" {
		http.Redirect(w, r, "/profile", 302)
	}

	name := r.FormValue("name")
	pass := r.FormValue("password")
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, w)
		http.Redirect(w, r, "/profile", 302)
	}

	templatemanager.Render(w, GetUserName(r), struct{}{}, "base", "templates/login.tpl", "templates/base.tpl")
}

// HandleRegister handle the greeeting
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if GetUserName(r) != "" {
		http.Redirect(w, r, "/profile", 302)
	}

	name := r.FormValue("name")
	pass := r.FormValue("password")
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, w)
		http.Redirect(w, r, "/profile", 302)
	}

	templatemanager.Render(w, GetUserName(r), struct{}{}, "base", "templates/register.tpl", "templates/base.tpl")
}

// HandleLogout logout the user
func HandleLogout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}
