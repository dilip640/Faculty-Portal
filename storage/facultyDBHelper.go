package storage

import (
	"database/sql"
	"time"
)

// InsertUpdateFaculty add new faculty
func InsertUpdateFaculty(uname, startDate, endDate, dept string) error {
	sqlStatement := `SELECT insert_update_faculty($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, uname, startDate, endDate, dept)
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
	dt := time.Now()
	faculty := Faculty{}
	sqlStatement := `SELECT emp_id, start_date, end_date, dept, cv_id FROM faculty WHERE emp_id = $1 AND end_date >= $2`
	err := db.QueryRow(sqlStatement, uname, dt.Format("2006-01-02")).Scan(
		&faculty.Uname, &faculty.StartDate, &faculty.EndDate, &faculty.Dept, &faculty.CVID)
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

// FacultyConvert extract datetime from date
func (f *Faculty) FacultyConvert() {
	f.StartDate = f.StartDate[:10]
	f.EndDate = f.EndDate[:10]
}
