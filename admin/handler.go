package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
		return
	}

	_, err := storage.GetAdmin(userName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		var err error
		var reqStruct adminEditRequest
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &reqStruct)

		if deptName := reqStruct.AddDept; deptName != nil {
			err = storage.InsertDepartment(*deptName)
		} else if deptID := reqStruct.DeleteDept; deptID != nil {
			i, _ := strconv.Atoi(*deptID)

			err = storage.DeleteDepartment(i)
		} else if postName := reqStruct.AddPost; postName != nil {
			err = storage.InsertPost(*postName)
		} else if postID := reqStruct.DeletePost; postID != nil {
			i, _ := strconv.Atoi(*postID)

			err = storage.DeletePost(i)
		} else if hod := reqStruct.AssignHOD; hod != nil {
			err = storage.InsertHOD(*hod)
		} else if routID := reqStruct.DeleteRoute; routID != nil {
			i, _ := strconv.Atoi(*routID)

			err = storage.DeleteRoute(i)
		} else if route := reqStruct.AddRoute; route != nil {
			err = storage.InsertRoute(*route)
		}

		if err != nil {
			log.Error(err)
			fmt.Fprint(w, "error")
			return
		}
		fmt.Fprint(w, "ok")
		return
	}

	data := struct {
		Departments []*storage.Department
		Posts       []*storage.Post
		Routes      []*storage.Route
		Hods        []*storage.HODDetail
	}{}

	depts, err := storage.GetAllDepartments()
	if err == nil {
		data.Departments = depts
	} else {
		log.Error(err)
	}

	posts, err := storage.GetAllPosts()
	if err == nil {
		data.Posts = posts
	} else {
		log.Error(err)
	}

	routes, err := storage.GetAllRoutes()
	if err == nil {
		data.Routes = routes
	} else {
		log.Error(err)
	}

	hods, err := storage.GetAllHOD()
	if err == nil {
		data.Hods = hods
	} else {
		log.Error(err)
	}

	templatemanager.Render(w, auth.GetUserName(r), data, "base",
		"templates/admin/index.html", "templates/base.html")
}

type adminEditRequest struct {
	DeleteDept  *string        `json:"deleteDept"`
	AddDept     *string        `json:"addDept"`
	DeletePost  *string        `json:"deletePost"`
	AddPost     *string        `json:"addPost"`
	AssignHOD   *storage.HOD   `json:"assignHOD"`
	DeleteRoute *string        `json:"deleteRoute"`
	AddRoute    *storage.Route `json:"addRoute"`
}
