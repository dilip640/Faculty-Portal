package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		"templates/index.html", "templates/base.html")
}

// HandleProfile handle the greeeting
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	self := true
	params := mux.Vars(r)
	if params["id"] != "" {
		userName = params["id"]
		self = false
	}
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Faculty  storage.Faculty
		Employee storage.Employee
		CVDetail storage.CVDetail
		Style    cssparam
		Self     bool
	}{Self: self}

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

	cvdetail, err := storage.GetCVDetails(userName)
	if err == nil {
		data.CVDetail = cvdetail
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/profile.html", "templates/faculty/cv_template.html", "templates/base.html")
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
		"templates/faculty/register.html", "templates/base.html")
}

// HandleCVEdit for edit cv
func HandleCVEdit(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	if r.Method == http.MethodPost {
		var err error
		var reqStruct cvEditRequest
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &reqStruct)

		log.Print(err, reqStruct)

		if bio := reqStruct.Biography; bio != nil {
			err = storage.SaveBio(userName, *bio)
		} else if aboutme := reqStruct.AboutMe; aboutme != nil {
			err = storage.SaveAboutme(userName, *aboutme)
		} else if project := reqStruct.Project; project != nil {
			err = storage.AddProject(userName, *project)
		}
		if err != nil {
			log.Error(err)
			fmt.Fprint(w, struct{ status string }{"error"})
			return
		}
		fmt.Fprint(w, struct{ status string }{"ok"})
		return
	}

	data := struct {
		CVDetail storage.CVDetail
		Style    cssparam
		Self     bool
	}{
		Style: cssparam{true, "card card-body bg-light"},
	}

	cvdetail, err := storage.GetCVDetails(userName)
	if err == nil {
		data.CVDetail = cvdetail
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/faculty/editCV.html", "templates/faculty/cv_template.html", "templates/base.html")
}

type cssparam struct {
	Edit  bool
	Class string
}

type cvEditRequest struct {
	Biography *string            `json:"biography"`
	AboutMe   *string            `json:"aboutme"`
	Project   *storage.CVProject `json:"project"`
}
