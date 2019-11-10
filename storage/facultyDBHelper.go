package storage

import (
	"database/sql"
	"time"

	"github.com/dilip640/Faculty-Portal/util"
)

// InsertFaculty add new faculty
func InsertFaculty(uname, startDate, dept string) error {
	sqlStatement := `INSERT INTO faculty (emp_id, start_date, dept)
									VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, uname, startDate, dept)
	return err
}

// UpdateFaculty update faculty
func UpdateFaculty(uname, startDate, endDate, dept string) error {
	sqlStatement := `UPDATE faculty SET dept = $3 end_date = $4
						WHERE emp_id = $1 AND start_date = $2;`
	_, err := db.Exec(sqlStatement, uname, startDate, dept, endDate)
	return err
}

// SetFacultyCV set faculty cv id
func SetFacultyCV(uname, cvID string) error {
	sqlStatement := `UPDATE faculty SET cv_id = $1 WHERE emp_id = $2`
	_, err := db.Exec(sqlStatement, cvID, uname)
	return err
}

// GetFacultyCV get faculty cv id
func GetFacultyCV(uname string) (string, error) {
	var cvID string
	dt := time.Now()
	sqlStatement := `Select cv_id FROM faculty WHERE emp_id = $1 AND end_date >= $2`
	err := db.QueryRow(sqlStatement, uname, dt.Format("2006-01-02")).Scan(&cvID)
	return cvID, err
}

// GetFacultyDetails returns faculty details
func GetFacultyDetails(uname string) (Faculty, error) {
	var (
		startDate string
		endDate   string
	)
	faculty := Faculty{}
	sqlStatement := `SELECT emp_id, start_date, end_date, dept, cv_id FROM faculty
						WHERE emp_id = $1 ORDER BY start_date DESC`
	err := db.QueryRow(sqlStatement, uname).Scan(
		&faculty.Uname, &startDate, &endDate, &faculty.Dept, &faculty.CVID)
	faculty.StartDate = util.DateTimeToDate(startDate)
	faculty.EndDate = util.DateTimeToDate(endDate)
	return faculty, err
}

// Faculty struct
type Faculty struct {
	Uname     string
	StartDate string
	EndDate   string
	Dept      string
	CVID      sql.NullString
}
