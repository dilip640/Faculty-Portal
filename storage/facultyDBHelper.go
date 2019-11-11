package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
)

// InsertFaculty add new faculty
func InsertFaculty(empId, startDate, deptID string) error {
	sqlStatement := `INSERT INTO faculty (emp_id, dept_id, start_date)
									VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, empId, deptID, startDate)
	return err
}

// UpdateFaculty update faculty
func UpdateFaculty(empId, deptId string) error {
	sqlStatement := `UPDATE faculty SET dept_id = $2
						WHERE emp_id = $1;`
	_, err := db.Exec(sqlStatement, empId, deptId)
	return err
}

// GetFacultyDetails returns faculty details
func GetFacultyDetails(empId string) (Faculty, error) {
	var startDate string

	faculty := Faculty{}
	sqlStatement := `SELECT emp_id, dept_id, start_date FROM faculty
						WHERE emp_id = $1`
	err := db.QueryRow(sqlStatement, empId).Scan(
		&faculty.EmpID, &faculty.DeptID, &startDate)
	faculty.StartDate = util.DateTimeToDate(startDate)
	return faculty, err
}

// Faculty struct
type Faculty struct {
	EmpID     string
	StartDate string
	DeptID    string
}
