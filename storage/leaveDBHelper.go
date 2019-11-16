package storage

import (
	"errors"

	"github.com/dilip640/Faculty-Portal/util"
	log "github.com/sirupsen/logrus"
)

// GetRemainingLeaves returns remaining leave
func GetRemainingLeaves(empID string, y ...int) (int, error) {
	var year int

	if len(y) == 1 {
		year = y[0]
	} else if len(y) == 0 {
		year = int(util.GetCurrentYear())
	} else {
		return year, errors.New("Invalid Arguments")
	}

	sqlStatement := `SELECT get_remaining_leave($1, $2)`

	var numLeaves int
	err := db.QueryRow(sqlStatement, empID, year).Scan(&numLeaves)

	return numLeaves, err
}

// CreateLeaveApplication create new leave
func CreateLeaveApplication(empID string, noOfDays int, startDate, applier, status, comment string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO leave_application(emp_id, no_of_days, start_date, 
				applier, route_status, status)
				VALUES($1, $2, $3, $4, $4, $5);`, empID, noOfDays, startDate, applier, status)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO leave_comment_history(leave_id, signed_by, comment, status)
				VALUES(currval('leave_application_id_seq'), $1, $2, $3);`, empID, comment, status)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

// GetActiveApplication returns the current active application.
func GetActiveApplication(empID string) (LeaveApplication, error) {
	var startDate string

	leaveApplication := LeaveApplication{}
	sqlStatement := `SELECT id, emp_id, no_of_days, time_stamp, applier, 
						route_status, status, start_date FROM leave_application
						WHERE emp_id = $1`
	err := db.QueryRow(sqlStatement, empID).Scan(
		&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
		&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
		&leaveApplication.Status, &startDate)
	leaveApplication.StartDate = util.DateTimeToDate(startDate)
	leaveApplication.LeaveCommentHistories, err = GetLeaveCommentHistory(leaveApplication.LeaveID)

	return leaveApplication, err
}

// GetActiveHodRequests will return all the Leave requests corresponding to that HOD
func GetActiveHodRequests(deptID string, routes []*Route) ([]*LeaveApplication, error) {
	reqs := make([]*LeaveApplication, 0)
	for _, route := range routes {
		rows, err := db.Query(
			`SELECT la.id, la.emp_id, la.no_of_days, la.time_stamp, la.applier, 
				la.route_status, la.status, la.start_date FROM leave_application AS la, faculty AS f 
				WHERE f.emp_id=la.emp_id AND f.dept_id=$1 AND la.route_status=$2 
				AND (la.status='PENDING' OR la.status='INITIALIZED')`, deptID, route.RouteFrom)
		if err != nil {
			log.Error(err)
		}
		defer rows.Close()
		if err == nil {
			for rows.Next() {
				leaveApplication := LeaveApplication{}
				var startDate string

				if err := rows.Scan(&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
					&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
					&leaveApplication.Status, &startDate); err == nil {
					leaveApplication.StartDate = util.DateTimeToDate(startDate)
					leaveApplication.LeaveCommentHistories, err = GetLeaveCommentHistory(leaveApplication.LeaveID)
					reqs = append(reqs, &leaveApplication)
				} else {
					log.Error(err)
				}
			}
		}

	}

	return reqs, nil
}

// GetLeaveCommentHistory returns all comment history of give application id
func GetLeaveCommentHistory(LeaveID string) ([]*LeaveCommentHistory, error) {
	leaveCommentHistories := make([]*LeaveCommentHistory, 0)

	rows, err := db.Query(
		`SELECT leave_id, signed_by, comment, status, time_stamp FROM leave_comment_history
					WHERE leave_id=$1`, LeaveID)
	if err != nil {
		return leaveCommentHistories, err
	}
	defer rows.Close()

	for rows.Next() {
		leaveCommentHistory := LeaveCommentHistory{}

		if err := rows.Scan(&leaveCommentHistory.LeaveID, &leaveCommentHistory.SignedBy,
			&leaveCommentHistory.Comment, &leaveCommentHistory.Status,
			&leaveCommentHistory.Timestamp); err == nil {

			leaveCommentHistories = append(leaveCommentHistories, &leaveCommentHistory)
		} else {
			log.Error(err)
		}
	}

	return leaveCommentHistories, nil
}

// LeaveApplication struct
type LeaveApplication struct {
	LeaveID               string
	EmpID                 string
	NumOfDays             int
	Timestamp             string
	Applier               string
	RouteStatus           string
	Status                string
	StartDate             string
	LeaveCommentHistories []*LeaveCommentHistory
}

// LeaveCommentHistory struct
type LeaveCommentHistory struct {
	LeaveID   string
	SignedBy  string
	Comment   string
	Status    string
	Timestamp string
}
