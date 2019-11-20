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

// DeductLeave deduct of an employee
func DeductLeave(empID string, year, days int) error {
	currLeaves, err := GetRemainingLeaves(empID, year)
	if err != nil {
		return err
	}
	borrowedDays := 0
	if days > currLeaves {
		borrowedDays = days - currLeaves
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE leave SET no_of_leaves = no_of_leaves - $1 WHERE emp_id = $2 AND year = $3`,
		days-borrowedDays, empID, year)
	if err != nil {
		tx.Rollback()
		return err
	}

	if borrowedDays > 0 {
		_, err = tx.Exec(`UPDATE leave SET no_of_leaves = no_of_leaves - $1 WHERE emp_id = $2 AND year = $3`,
			borrowedDays, empID, year+1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

// CreateLeaveApplication create new leave
func CreateLeaveApplication(empID string, noOfDays, borrowedDays int, startDate, applier,
	routeStatus, status, comment string) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO leave_application(emp_id, no_of_days, start_date, 
				applier, route_status, status)
				VALUES($1, $2, $3, $4, $5, $6);`, empID, noOfDays, startDate, applier, routeStatus, status)
	if err != nil {
		tx.Rollback()
		return err
	}

	if borrowedDays > 0 {
		_, err = tx.Exec(`INSERT INTO borrowed_leave(leave_id, no_of_days)
				VALUES(currval('leave_application_id_seq'), $1);`, borrowedDays)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = tx.Exec(`INSERT INTO leave_comment_history(leave_id, signed_by, comment, status, position)
				VALUES(currval('leave_application_id_seq'), $1, $2, $3, $4);`, empID, comment, status, applier)
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
						WHERE emp_id = $1 AND (status='PENDING' OR status='INITIALIZED')`
	err := db.QueryRow(sqlStatement, empID).Scan(
		&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
		&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
		&leaveApplication.Status, &startDate)
	leaveApplication.StartDate = util.DateTimeToDate(startDate)
	leaveApplication.LeaveCommentHistories, err = GetLeaveCommentHistory(leaveApplication.LeaveID)
	leaveApplication.BorrowedDays, _ = GetBorrowedLeave(leaveApplication.LeaveID)

	return leaveApplication, err
}

// GetPastApplications returns the past application.
func GetPastApplications(empID string) ([]*LeaveApplication, error) {
	leaveApplications := make([]*LeaveApplication, 0)

	sqlStatement := `SELECT id, emp_id, no_of_days, time_stamp, applier, 
						route_status, status, start_date FROM leave_application
						WHERE emp_id = $1 AND NOT (status='PENDING' OR status='INITIALIZED')`
	rows, err := db.Query(sqlStatement, empID)

	if err != nil {
		return leaveApplications, err
	}

	for rows.Next() {
		leaveApplication := LeaveApplication{}
		var startDate string

		if err := rows.Scan(&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
			&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
			&leaveApplication.Status, &startDate); err == nil {
			leaveApplication.StartDate = util.DateTimeToDate(startDate)
			leaveApplication.LeaveCommentHistories, _ = GetLeaveCommentHistory(leaveApplication.LeaveID)
			leaveApplication.BorrowedDays, _ = GetBorrowedLeave(leaveApplication.LeaveID)

			leaveApplications = append(leaveApplications, &leaveApplication)
		} else {
			log.Error(err)
		}
	}

	return leaveApplications, nil
}

// GetLeaveApplication retuns LeaveApplication
func GetLeaveApplication(leavID int) (LeaveApplication, error) {
	var leaveApplication LeaveApplication

	sqlStatement := `SELECT id, emp_id, no_of_days, time_stamp, applier, 
						route_status, status, start_date FROM leave_application
						WHERE id = $1`
	err := db.QueryRow(sqlStatement, leavID).Scan(
		&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
		&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
		&leaveApplication.Status, &leaveApplication.StartDate)
	leaveApplication.StartDate = util.DateTimeToDate(leaveApplication.StartDate)
	leaveApplication.LeaveCommentHistories, _ = GetLeaveCommentHistory(leaveApplication.LeaveID)
	leaveApplication.BorrowedDays, _ = GetBorrowedLeave(leaveApplication.LeaveID)

	return leaveApplication, err
}

// GetActiveHodRequests will return all the Leave requests corresponding to that HOD
func GetActiveHodRequests(deptID int) ([]*LeaveApplication, error) {
	reqs := make([]*LeaveApplication, 0)
	rows, err := db.Query(
		`SELECT la.id, la.emp_id, la.no_of_days, la.time_stamp, la.applier, 
				la.route_status, la.status, la.start_date FROM leave_application AS la, faculty AS f 
				WHERE f.emp_id=la.emp_id AND f.dept_id=$1 AND la.route_status='hod'
				AND (la.status='PENDING' OR la.status='INITIALIZED') 
				AND la.applier <> 'hod'`, deptID)
	if err != nil {
		return reqs, err
	}
	defer rows.Close()

	for rows.Next() {
		leaveApplication := LeaveApplication{}
		var startDate string

		if err := rows.Scan(&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
			&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
			&leaveApplication.Status, &startDate); err == nil {
			leaveApplication.StartDate = util.DateTimeToDate(startDate)
			leaveApplication.LeaveCommentHistories, _ = GetLeaveCommentHistory(leaveApplication.LeaveID)
			leaveApplication.BorrowedDays, _ = GetBorrowedLeave(leaveApplication.LeaveID)

			reqs = append(reqs, &leaveApplication)
		} else {
			log.Error(err)
		}

	}

	return reqs, nil
}

// GetActiveCCFRequests will return all the Leave requests corresponding to that post
func GetActiveCCFRequests(postName string) ([]*LeaveApplication, error) {
	reqs := make([]*LeaveApplication, 0)
	rows, err := db.Query(`SELECT la.id, la.emp_id, la.no_of_days, la.time_stamp, la.applier, 
							la.route_status, la.status, la.start_date FROM leave_application AS la
							WHERE EXISTS(SELECT * FROM application_route AS ar 
												WHERE ar.applier = la.applier AND ar.ccf_post = $1)
							AND la.route_status='ccf' AND (la.status='PENDING' OR la.status='INITIALIZED') 
							AND la.applier <> 'ccf'`, postName)
	if err != nil {
		return reqs, err
	}
	defer rows.Close()

	for rows.Next() {
		leaveApplication := LeaveApplication{}
		var startDate string

		if err := rows.Scan(&leaveApplication.LeaveID, &leaveApplication.EmpID, &leaveApplication.NumOfDays,
			&leaveApplication.Timestamp, &leaveApplication.Applier, &leaveApplication.RouteStatus,
			&leaveApplication.Status, &startDate); err == nil {
			leaveApplication.StartDate = util.DateTimeToDate(startDate)
			leaveApplication.LeaveCommentHistories, _ = GetLeaveCommentHistory(leaveApplication.LeaveID)
			leaveApplication.BorrowedDays, _ = GetBorrowedLeave(leaveApplication.LeaveID)

			reqs = append(reqs, &leaveApplication)
		} else {
			log.Error(err)
		}

	}

	return reqs, nil
}

// GetLeaveCommentHistory returns all comment history of give application id
func GetLeaveCommentHistory(LeaveID int) ([]*LeaveCommentHistory, error) {
	leaveCommentHistories := make([]*LeaveCommentHistory, 0)

	rows, err := db.Query(
		`SELECT leave_id, signed_by, comment, status, time_stamp, position FROM leave_comment_history
					WHERE leave_id=$1`, LeaveID)
	if err != nil {
		return leaveCommentHistories, err
	}
	defer rows.Close()

	for rows.Next() {
		leaveCommentHistory := LeaveCommentHistory{}

		if err := rows.Scan(&leaveCommentHistory.LeaveID, &leaveCommentHistory.SignedBy,
			&leaveCommentHistory.Comment, &leaveCommentHistory.Status,
			&leaveCommentHistory.Timestamp, &leaveCommentHistory.Position); err == nil {

			leaveCommentHistories = append(leaveCommentHistories, &leaveCommentHistory)
		} else {
			log.Error(err)
		}
	}

	return leaveCommentHistories, nil
}

// CommentAndChangeLeaveStatus handle comment
func CommentAndChangeLeaveStatus(routeStatus, status string, leaveCommentHistory LeaveCommentHistory) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE leave_application SET route_status = $1, status = $2
						WHERE id = $3;`, routeStatus, status, leaveCommentHistory.LeaveID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`INSERT INTO leave_comment_history(leave_id, signed_by, comment, status, position)
				VALUES($1, $2, $3, $4, $5);`, leaveCommentHistory.LeaveID, leaveCommentHistory.SignedBy,
		leaveCommentHistory.Comment, leaveCommentHistory.Status, leaveCommentHistory.Position)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

// GetLatestRoute give us the route from which the application was sent back.
func GetLatestRoute(leaveID int) (string, error) {
	var latestRoute string
	err := db.QueryRow(`SELECT position FROM leave_comment_history 
							WHERE leave_id = $1 ORDER BY time_stamp DESC`, leaveID).Scan(&latestRoute)

	return latestRoute, err
}

// GetBorrowedLeave returns borrored leave
func GetBorrowedLeave(leaveID int) (int, error) {
	borrowedDays := 0
	err := db.QueryRow(`SELECT no_of_days FROM borrowed_leave 
							WHERE leave_id = $1`, leaveID).Scan(&borrowedDays)

	return borrowedDays, err
}

// LeaveApplication struct
type LeaveApplication struct {
	LeaveID               int
	EmpID                 string
	NumOfDays             int
	Timestamp             string
	Applier               string
	RouteStatus           string
	Status                string
	StartDate             string
	LeaveCommentHistories []*LeaveCommentHistory
	BorrowedDays          int
}

// LeaveCommentHistory struct
type LeaveCommentHistory struct {
	LeaveID   int
	SignedBy  string
	Comment   string
	Status    string
	Timestamp string
	Position  string
}
