package templatemanager

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Render renders to page
func Render(w http.ResponseWriter, auth string, input interface{}, name string, filenames ...string) {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	err = t.ExecuteTemplate(w, name, struct {
		User string
		data interface{}
	}{auth, input})
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
