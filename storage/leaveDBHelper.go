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
func CreateLeaveApplication(empID string, noOfDays int, applier, status, comment string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO leave_application(emp_id, no_of_days, applier, route_status, status)
				VALUES($1, $2, $3, $3, $4);`, empID, noOfDays, applier, status)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`INSERT INTO leave_application_history(leave_id, emp_id, signed_by, comment, status)
				VALUES(currval('leave_application_id_seq'), $1, $1, $2, $3);`, empID, comment, status)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}
