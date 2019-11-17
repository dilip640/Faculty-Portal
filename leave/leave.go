package leave

import (
	"database/sql"
	"errors"

	"github.com/dilip640/Faculty-Portal/storage"
)

func requestLeave(noOfDays int, startDate, comment, empID string) error {
	currLeave, err := storage.GetRemainingLeaves(empID)
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

	if noOfDays <= currLeave {
		route, err := storage.GetRouteStatusTo(applier, applier)
		if err != nil {
			return err
		}
		err = storage.CreateLeaveApplication(empID, noOfDays, startDate, applier, route.RouteTo, "INITIALIZED", comment)
	}
	return err
}

// GetActiveLeaveReqs returns leave requests
func GetActiveLeaveReqs(empID string) ([]*storage.LeaveApplication, error) {
	var leaveApplications []*storage.LeaveApplication

	hodDetails, err := storage.GetHodDetails(empID)
	if err == nil {
		leaveApplications, err = storage.GetActiveHodRequests(hodDetails.DeptID, "hod")
		return leaveApplications, err
	}

	// ccFacultyDetails, err := storage.GetCCFacultyDetails(empID)
	// if err == nil {

	// }

	return leaveApplications, err

}

// ValidateComment validate comment and add to application
func ValidateComment(leaveCommentHistory storage.LeaveCommentHistory) error {
	leaveApplication, err := storage.GetLeaveApplication(leaveCommentHistory.LeaveID)
	if err != nil {
		return err
	}

	if leaveCommentHistory.Status == "send_back" {
		routeStatus, err := storage.GetRouteStatusFrom(leaveApplication.Applier, leaveApplication.RouteStatus)
		if err != nil {
			return err
		}

		err = storage.CommentAndChangeLeaveStatus(routeStatus.RouteFrom, "PENDING", leaveCommentHistory)
		if err != nil {
			return err
		}

	} else if leaveCommentHistory.Status == "approve" {
		routeStatus, err := storage.GetRouteStatusTo(leaveApplication.Applier, leaveApplication.RouteStatus)
		if err == sql.ErrNoRows {
			err = storage.CommentAndChangeLeaveStatus(routeStatus.RouteFrom, "APPROVED", leaveCommentHistory)
			if err != nil {
				return err
			}

			return nil
		} else if err != nil {
			return err
		}

		err = storage.CommentAndChangeLeaveStatus(routeStatus.RouteFrom, "PENDING", leaveCommentHistory)
		if err != nil {
			return err
		}
	} else if leaveCommentHistory.Status == "disapprove" {
		err = storage.CommentAndChangeLeaveStatus(leaveApplication.RouteStatus, "DISAPPROVED", leaveCommentHistory)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid Opertion")
	}

	return nil
}
