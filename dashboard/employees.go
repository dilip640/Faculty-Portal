package dashboard

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	log "github.com/sirupsen/logrus"
)

// HandleEmployeeDetails handles all the details of employees.
func HandleEmployeeDetails(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Faculties   []*storage.FacultyDetails
		CCFaculties []*storage.CCFacultyDetails
		Hods        []*storage.HODDetail
	}{}

	faculties, err := storage.GetAllFacultyDetails()
	if err == nil {
		data.Faculties = faculties
	} else {
		log.Error(err)
	}

	ccFaculties, err := storage.GetAllccFacultyDetails()
	if err == nil {
		data.CCFaculties = ccFaculties
	} else {
		log.Error(err)
	}

	hods, err := storage.GetAllHOD()
	if err == nil {
		data.Hods = hods
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, "", data, "base",
		"templates/employeeDetails.html", "templates/base.html")
}
