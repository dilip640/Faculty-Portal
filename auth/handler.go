package auth

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HandleLogin handle the greeeting
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.tpl", "templates/base.tpl")
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	err = t.ExecuteTemplate(w, "base", struct{}{})
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
