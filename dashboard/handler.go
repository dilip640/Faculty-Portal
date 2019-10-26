package dashboard

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// HandleHome handle the greeeting
func HandleHome(w http.ResponseWriter, r *http.Request) {
	templatemanager.Render(w, auth.GetUserName(r), struct{}{}, "base",
		"templates/index.tpl", "templates/base.tpl")
}

// HandleProfile handle the greeeting
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	params := mux.Vars(r)
	if params["id"] != "" {
		userName = params["id"]
	}
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Faculty  storage.Faculty
		Employee storage.Employee
	}{}

	faculty, err := storage.GetFacultyDetails(userName)
	if err == nil {
		faculty.FacultyConvert()
		data.Faculty = faculty
	} else {
		log.Error(err)
	}

	employee, err := storage.GetEmployeeDetails(userName)
	if err == nil {
		data.Employee = employee
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/profile.tpl", "templates/base.tpl")
}

// HandleUpdateRegisterFaculty for update and register faculty
func HandleUpdateRegisterFaculty(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Error string
	}{}

	startDate := r.FormValue("start_date")
	endDate := r.FormValue("end_date")
	dept := r.FormValue("dept")

	if startDate != "" && endDate != "" && dept != "" {
		data.Error = "Enter Correct Details!"
		err := storage.InsertUpdateFaculty(userName, startDate, endDate, dept)
		if err == nil {
			http.Redirect(w, r, "/profile", 302)
		} else {
			if err.Error() == "pq: Faculty Exists!" {
				data.Error = "Faculty Exists!"
			}
			log.Error(err)
		}
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/faculty/register.tpl", "templates/base.tpl")
}
