package leave

import (
	"database/sql"
	"errors"

	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/util"
)

func requestLeave(noOfDays int, startDate, comment, empID string) error {
	currLeave, err := storage.GetRemainingLeaves(empID, util.DateTToYear(startDate))
	if err != nil {
		return err
	}

	nxtLeave, err := storage.GetRemainingLeaves(empID, util.DateTToYear(startDate)+1)
	if err != nil {
		return err
	}

	var applier string

	_, err = storage.GetCCFacultyDetails(empID)
	if err == nil {
		applier = "cc_faculty"
	} else if _, err = storage.GetHodDetails(empID); err == nil {
		applier = "hod"
	} else if _, err = storage.GetFacultyDetails(empID); err == nil {
		applier = "faculty"
	} else {
		return errors.New("Nhi milegi")
	}

	route, err := storage.GetRouteStatusTo(applier, applier)
	if err != nil {
		return err
	}

	if noOfDays <= currLeave+nxtLeave {
		borrowedDays := 0
		if noOfDays > currLeave {
			borrowedDays = noOfDays - currLeave
		}
		err = storage.CreateLeaveApplication(empID, noOfDays, borrowedDays, startDate, applier,
			route.RouteTo, "INITIALIZED", comment)
		return err
	}

	return errors.New("Itni chhutti nhi milegi")
}

// GetActiveLeaveReqs returns leave requests
func GetActiveLeaveReqs(empID string) ([]*storage.LeaveApplication, error) {
	var leaveApplications []*storage.LeaveApplication

	hodDetails, err := storage.GetHodDetails(empID)
	if err == nil {
		leaveApplications, err = storage.GetActiveHodRequests(hodDetails.DeptID)
		return leaveApplications, err
	}

	ccFacultyDetails, err := storage.GetCCFacultyDetails(empID)
	if err == nil {
		leaveApplications, err = storage.GetActiveCCFRequests(ccFacultyDetails.Post.Name)
		return leaveApplications, err
	}

	return leaveApplications, err

}

// ValidateComment validate comment and add to application
func ValidateComment(leaveCommentHistory storage.LeaveCommentHistory, borrowAlowed bool) error {
	leaveApplication, err := storage.GetLeaveApplication(leaveCommentHistory.LeaveID)
	if err != nil {
		return err
	}

	// For position of added comment.
	leaveCommentHistory.Position = leaveApplication.RouteStatus

	if leaveCommentHistory.Status == "send_back" {
		err = storage.CommentAndChangeLeaveStatus(leaveApplication.Applier, "PENDING", leaveCommentHistory)
		if err != nil {
			return err
		}

	} else if leaveCommentHistory.Status == "approve" {
		if !borrowAlowed && (leaveApplication.BorrowedDays > 0) {
			return errors.New("Please allow borrow leave also")
		}

		routeStatus, err := storage.GetRouteStatusTo(leaveApplication.Applier, leaveApplication.RouteStatus)
		if err == sql.ErrNoRows {
			err = storage.CommentAndChangeLeaveStatus(leaveApplication.RouteStatus, "APPROVED", leaveCommentHistory)
			if err != nil {
				return err
			}
			err := storage.DeductLeave(leaveApplication.EmpID, util.DateTToYear(leaveApplication.StartDate),
				leaveApplication.NumOfDays)

			return err
		} else if err != nil {
			return err
		}

		err = storage.CommentAndChangeLeaveStatus(routeStatus.RouteTo, "PENDING", leaveCommentHistory)
		if err != nil {
			return err
		}
	} else if leaveCommentHistory.Status == "disapprove" {
		err = storage.CommentAndChangeLeaveStatus(leaveApplication.RouteStatus, "DISAPPROVED", leaveCommentHistory)
		if err != nil {
			return err
		}
	} else if leaveCommentHistory.Status == "add_comment" {
		routeStatus, err := storage.GetLatestRoute(leaveApplication.LeaveID)
		if err != nil {
			return err
		}

		err = storage.CommentAndChangeLeaveStatus(routeStatus, "PENDING", leaveCommentHistory)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid Opertion")
	}

	return nil
}
