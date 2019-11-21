package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
	log "github.com/sirupsen/logrus"
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

// GetAllFacultyDetails returns all  the faculties.
func GetAllFacultyDetails() ([]*FacultyDetails, error) {
	faculties := make([]*FacultyDetails, 0)

	rows, err := db.Query(`SELECT emp_id, dept_id, start_date FROM faculty`)
	if err != nil {
		return faculties, err
	}

	defer rows.Close()

	for rows.Next() {
		var startDate string
		faculty := FacultyDetails{}

		if err := rows.Scan(&faculty.EmpID, &faculty.Dept.DeptID, &startDate); err == nil {
			faculty.StartDate = util.DateTimeToDate(startDate)
			faculty.Dept, err = GetDepartment(faculty.Dept.DeptID)
			faculty.EmployeeDetail, _ = GetEmployeeDetails(faculty.EmpID)
			faculties = append(faculties, &faculty)
		} else {
			log.Error(err)
		}
	}

	return faculties, nil
}

// Faculty struct
type Faculty struct {
	EmpID     string
	StartDate string
	Dept      Department
}

// Faculty struct
type FacultyDetails struct {
	EmpID          string
	EmployeeDetail Employee
	StartDate      string
	Dept           Department
}
