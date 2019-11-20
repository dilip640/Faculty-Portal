package leave

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
			err = requestLeave(applyLeaveData.NoOfDays, applyLeaveData.StartDate, applyLeaveData.Comment, userName)
			if err != nil {
				response = "You Have Already an active application!"
			}
		} else if commentLeaveReq := reqStruct.CommentLeaveReq; commentLeaveReq != nil {
			leaveCommentHistory := storage.LeaveCommentHistory{}

			leaveCommentHistory.LeaveID = commentLeaveReq.LeaveID
			leaveCommentHistory.SignedBy = userName
			leaveCommentHistory.Comment = commentLeaveReq.Comment
			leaveCommentHistory.Status = commentLeaveReq.Action

			err = ValidateComment(leaveCommentHistory, reqStruct.CommentLeaveReq.BorrowedAllowed)
			if err != nil {
				response = "Something went wrong!"
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

// HandleLeaveRequest for leave request
func HandleLeaveRequest(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	data := struct {
		ActiveLeaveReqs []*storage.LeaveApplication
	}{}

	if userName != "" {

		activeLeaveReqs, err := GetActiveLeaveReqs(userName)
		if err == nil {
			data.ActiveLeaveReqs = activeLeaveReqs
		} else {
			log.Error(err)
		}
	}

	templatemanager.Render(w, userName, data, "base",
		"templates/leave_request.html", "templates/base.html")
}

type leaveRequest struct {
	ApplyForLeave   *applyLeaveRequest `json:"applyForLeave"`
	CommentLeaveReq *commentLeaveReq   `json:"commentLeaveReq"`
}

type applyLeaveRequest struct {
	NoOfDays  int    `json:"no_of_days"`
	StartDate string `json:"start_date"`
	Comment   string `json:"comment"`
}

type commentLeaveReq struct {
	Comment         string `json:"comment"`
	LeaveID         int    `json:"leave_id"`
	Action          string `json:"action"`
	BorrowedAllowed bool   `json:"borrow_approved"`
}
