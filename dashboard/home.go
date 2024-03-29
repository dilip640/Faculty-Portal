package dashboard

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/auth"
	"github.com/dilip640/Faculty-Portal/leave"
	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	log "github.com/sirupsen/logrus"
)

// HandleHome handle the greeeting
func HandleHome(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)

	data := struct {
		ActiveLeaveApplication storage.LeaveApplication
		PastLeaveApplications  []*storage.LeaveApplication
		ActiveLeaveReqs        []*storage.LeaveApplication
		NumLeaves              int
	}{}

	if userName != "" {

		numLeaves, err := storage.GetRemainingLeaves(userName)
		if err == nil {
			data.NumLeaves = numLeaves
		} else {
			log.Error(err)
		}

		activeLeaveApplication, err := storage.GetActiveApplication(userName)
		if err == nil {
			data.ActiveLeaveApplication = activeLeaveApplication
		} else {
			log.Error(err)
		}

		pastLeaveApplications, err := storage.GetPastApplications(userName)
		if err == nil {
			data.PastLeaveApplications = pastLeaveApplications
		} else {
			log.Error(err)
		}

		activeLeaveReqs, err := leave.GetActiveLeaveReqs(userName)
		if err == nil {
			data.ActiveLeaveReqs = activeLeaveReqs
		} else {
			log.Error(err)
		}
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/index.html", "templates/base.html")
}
