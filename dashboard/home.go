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
		ActiveLeaveReqs        []*storage.LeaveApplication
	}{}

	if userName != "" {
		activeLeaveApplication, err := storage.GetActiveApplication(userName)
		if err == nil {
			data.ActiveLeaveApplication = activeLeaveApplication
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
