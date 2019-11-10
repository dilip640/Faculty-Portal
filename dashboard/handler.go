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
		CVDetail storage.CVDetail
		Style    cssparam
		Self     bool
	}{Self: self}

	cvdetail, err := storage.GetCVDetails(userName)
	if err == nil {
		data.CVDetail = cvdetail
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/profile.html", "templates/faculty/cv_template.html", "templates/base.html")
}

// HandleRegisterFaculty for update and register faculty
func HandleRegisterFaculty(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Faculty  storage.Faculty
		Employee storage.Employee
		Error    string
	}{}

	startDate := r.FormValue("start_date")
	dept := r.FormValue("dept")

	if startDate != "" && dept != "" {
		data.Error = "Enter Correct Details!"
		err := storage.InsertFaculty(userName, startDate, dept)
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

// HandleUpdateFaculty for update and register faculty
func HandleUpdateFaculty(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Error string
	}{}

	faculty, err := storage.GetFacultyDetails(userName)
	if err != nil {
		log.Error(err)
	}

	endDate := r.FormValue("end_date")
	dept := r.FormValue("dept")

	if endDate != "" && dept != "" {
		data.Error = "Enter Correct Details!"
		err := storage.UpdateFaculty(userName, faculty.StartDate, endDate, dept)
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
		"templates/faculty/update.html", "templates/base.html")
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
		} else if dproject := reqStruct.Deleteproject; dproject != nil {
			err = storage.DeleteProject(userName, *dproject)
		} else if prize := reqStruct.Prize; prize != nil {
			err = storage.AddPrize(userName, *prize)
		} else if dprize := reqStruct.Deleteprize; dprize != nil {
			err = storage.DeletePrize(userName, *dprize)
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
		Faculty  storage.Faculty
		Employee storage.Employee
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
	Biography     *string            `json:"biography"`
	AboutMe       *string            `json:"aboutme"`
	Project       *storage.CVProject `json:"project"`
	Deleteproject *string            `json:"deleteproject"`
	Prize         *storage.CVPrize   `json:"prize"`
	Deleteprize   *string            `json:"deleteprize"`
}
