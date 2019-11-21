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
		Faculties       []*storage.FacultyDetails
		CCFaculties     []*storage.CCFacultyDetails
		Hods            []*storage.HODDetail
		Director        storage.CCFacultyDetails
		PastFaculties   []*storage.FacultyDetails
		PastCCFaculties []*storage.CCFacultyDetails
		PastHods        []*storage.HODDetail
	}{}

	faculties, err := storage.GetAllFacultyDetails()
	if err == nil {
		data.Faculties = faculties
	} else {
		log.Error(err)
	}

	pastfaculties, err := storage.GetAllPastFacultyDetails()
	if err == nil {
		data.PastFaculties = pastfaculties
	} else {
		log.Error(err)
	}

	ccFaculties, err := storage.GetAllccFacultyDetails()
	if err == nil {
		data.CCFaculties = ccFaculties
	} else {
		log.Error(err)
	}

	pastccFaculties, err := storage.GetAllPastccFacultyDetails()
	if err == nil {
		data.PastCCFaculties = pastccFaculties
	} else {
		log.Error(err)
	}

	hods, err := storage.GetAllHOD()
	if err == nil {
		data.Hods = hods
	} else {
		log.Error(err)
	}

	pasthods, err := storage.GetAllPastHOD()
	if err == nil {
		data.PastHods = pasthods
	} else {
		log.Error(err)
	}

	dir, err := storage.GetDirector()
	if err == nil {
		data.Director = dir
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, "", data, "base",
		"templates/employeeDetails.html", "templates/base.html")
}
