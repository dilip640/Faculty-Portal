package personalcv

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
