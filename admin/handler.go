package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	if r.Method == http.MethodPost {
		var err error
		var reqStruct deptEditRequest
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &reqStruct)

		if deptName := reqStruct.AddDept; deptName != nil {
			err = storage.InsertDepartment(*deptName)
		} else if deptID := reqStruct.DeleteDept; deptID != nil {
			err = storage.DeleteDepartment(*deptID)
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

type deptEditRequest struct {
	DeleteDept *string `json:"deleteDept"`
	AddDept    *string `json:"addDept"`
}
