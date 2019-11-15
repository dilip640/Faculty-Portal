package storage

import (
	"errors"

	"github.com/dilip640/Faculty-Portal/util"
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

	return leaveApplication, err
}

type LeaveApplication struct {
	LeaveID     string
	EmpID       string
	NumOfDays   int
	Timestamp   string
	Applier     string
	RouteStatus string
	Status      string
	StartDate   string
}
