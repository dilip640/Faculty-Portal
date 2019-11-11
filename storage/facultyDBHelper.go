package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
)

// InsertFaculty add new faculty
func InsertFaculty(empID, startDate, deptID string) error {
	sqlStatement := `INSERT INTO faculty (emp_id, dept_id, start_date)
									VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, empID, deptID, startDate)
	return err
}

// UpdateFaculty update faculty
func UpdateFaculty(empID, deptID string) error {
	sqlStatement := `UPDATE faculty SET dept_id = $2
						WHERE emp_id = $1;`
	_, err := db.Exec(sqlStatement, empID, deptID)
	return err
}

// GetFacultyDetails returns faculty details
func GetFacultyDetails(empID string) (Faculty, error) {
	var startDate string

	faculty := Faculty{}
	sqlStatement := `SELECT emp_id, dept_id, start_date FROM faculty
						WHERE emp_id = $1`
	err := db.QueryRow(sqlStatement, empID).Scan(
		&faculty.EmpID, &faculty.Dept.DeptID, &startDate)
	faculty.StartDate = util.DateTimeToDate(startDate)
	faculty.Dept, err = GetDepartment(faculty.Dept.DeptID)
	return faculty, err
}

// Faculty struct
type Faculty struct {
	EmpID     string
	StartDate string
	Dept      Department
}
