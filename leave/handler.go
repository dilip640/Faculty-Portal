package leave

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dilip640/Faculty-Portal/auth"
	log "github.com/sirupsen/logrus"
)

// HandleLeave for leave application
func HandleLeave(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	if r.Method == http.MethodPost {
		var response string
		var err error
		var reqStruct leaveRequest
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(bodyBytes, &reqStruct)

		if applyLeaveData := reqStruct.ApplyForLeave; applyLeaveData != nil {
			i, _ := strconv.Atoi(*applyLeaveData.NoOfDays)
			err = requestLeave(i, *applyLeaveData.StartDate, *applyLeaveData.Comment, userName)
			if err != nil {
				response = "You Have Already an active application!"
			}
		}

		if err != nil {
			log.Error(err)
			fmt.Fprint(w, response)
			return
		}
		fmt.Fprint(w, "ok")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

type leaveRequest struct {
	ApplyForLeave *applyLeaveRequest `json:"applyForLeave"`
}

type applyLeaveRequest struct {
	NoOfDays  *string `json:"no_of_days"`
	StartDate *string `json:"start_date"`
	Comment   *string `json:"comment"`
}
