package dashboard

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/templatemanager"
)

// HandleHome handle the greeeting
func HandleHome(w http.ResponseWriter, r *http.Request) {
	templatemanager.Render(w, auth.GetUserName(r), struct{}{}, "base", "templates/index.tpl", "templates/base.tpl")
}

// HandleProfile handle the greeeting
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	if auth.GetUserName(r) == "" {
		http.Redirect(w, r, "/login", 302)
	}

	templatemanager.Render(w, auth.GetUserName(r), struct{}{}, "base", "templates/profile.tpl", "templates/base.tpl")
}
