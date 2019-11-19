package templatemanager

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Render renders to page
func Render(w http.ResponseWriter, auth string, input interface{}, name string, filenames ...string) {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"html": func(value interface{}) template.HTML {
			return template.HTML(fmt.Sprint(value))
		},
	}).ParseFiles(filenames...))

	data := struct {
		User      string
		Data      interface{}
		Faculty   storage.Faculty
		CCFaculty storage.CCFaculty
		Employee  storage.Employee
		Hod       storage.HOD
	}{User: auth, Data: input}

	faculty, err := storage.GetFacultyDetails(auth)
	if err == nil {
		data.Faculty = faculty
	} else {
		log.Error(err)
	}

	ccFaculty, err := storage.GetCCFacultyDetails(auth)
	if err == nil {
		data.CCFaculty = ccFaculty
	} else {
		log.Error(err)
	}

	employee, err := storage.GetEmployeeDetails(auth)
	if err == nil {
		data.Employee = employee
	} else {
		log.Error(err)
	}

	hod, err := storage.GetHodDetails(auth)
	if err == nil {
		data.Hod = hod
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

// ServeStatic to serve static files
func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"css": staticDirectory + "/css/",
		"js":  staticDirectory + "/js/",
		"img": staticDirectory + "/img/",
	}
	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
}
