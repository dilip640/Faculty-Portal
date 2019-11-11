package admin

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	log "github.com/sirupsen/logrus"
)

// HandleAdmin handle the greeeting
func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
	}

	data := struct {
		Departments []*storage.Department
	}{}
	depts, err := storage.GetAllDepartments()
	if err == nil {
		data.Departments = depts
	} else {
		log.Error(err)
	}
	templatemanager.Render(w, auth.GetUserName(r), data, "base",
		"templates/admin/index.html", "templates/base.html")
}
