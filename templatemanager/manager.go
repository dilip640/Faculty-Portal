package templatemanager

import (
	"html/template"
	"net/http"

	"github.com/dilip640/Faculty-Portal/storage"
	log "github.com/sirupsen/logrus"
)

// Render renders to page
func Render(w http.ResponseWriter, auth string, input interface{}, name string, filenames ...string) {
	t, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := struct {
		User     string
		Data     interface{}
		Faculty  storage.Faculty
		Employee storage.Employee
	}{User: auth, Data: input}

	faculty, err := storage.GetFacultyDetails(auth)
	if err == nil {
		data.Faculty = faculty
	} else {
		log.Error(err)
	}

	employee, err := storage.GetEmployeeDetails(auth)
	if err == nil {
		data.Employee = employee
	} else {
		log.Error(err)
	}

	err = t.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
