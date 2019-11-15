package dashboard

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// HandleProfile handle the greeeting
func HandleProfile(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	self := true
	params := mux.Vars(r)
	if params["id"] != "" {
		self = false
		if userName == params["id"] {
			self = true
		}
		userName = params["id"]
	}
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	data := struct {
		CVDetail  storage.CVDetail
		Style     cssparam
		Self      bool
		NumLeaves int
	}{Self: self}

	cvdetail, err := storage.GetCVDetails(userName)
	if err == nil {
		data.CVDetail = cvdetail
	} else {
		log.Error(err)
	}

	numLeaves, err := storage.GetRemainingLeaves(userName)
	if err == nil {
		data.NumLeaves = numLeaves
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
		return
	}

	data := struct {
		Error       string
		Departments []*storage.Department
	}{}

	if r.Method == http.MethodPost {
		startDate := r.FormValue("start_date")
		dept := r.FormValue("dept")

		if startDate != "" && dept != "" {
			err := storage.InsertFaculty(userName, startDate, dept)
			if err == nil {
				storage.CreateCV(userName)
				http.Redirect(w, r, "/profile", 302)
				return
			}

			data.Error = "Faculty Exists!"
			log.Error(err)

		} else {
			data.Error = "Please enter all the details."
		}
	}
	depts, err := storage.GetAllDepartments()

	if err == nil {
		data.Departments = depts
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/faculty/register.html", "templates/base.html")
}

// HandleRegisterCCFaculty for update and register faculty
func HandleRegisterCCFaculty(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	data := struct {
		Error string
		Posts []*storage.Post
	}{}

	if r.Method == http.MethodPost {
		startDate := r.FormValue("start_date")
		postID := r.FormValue("post")

		if startDate != "" && postID != "" {
			err := storage.InsertCCFaculty(userName, startDate, postID)
			if err == nil {
				http.Redirect(w, r, "/profile", 302)
				return
			}

			data.Error = "Cross Cutting Faculty Exists!"
			log.Error(err)

		} else {
			data.Error = "Please enter all the details."
		}
	}
	posts, err := storage.GetAllPosts()

	if err == nil {
		data.Posts = posts
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/ccfaculty/register.html", "templates/base.html")
}

type cssparam struct {
	Edit  bool
	Class string
}
