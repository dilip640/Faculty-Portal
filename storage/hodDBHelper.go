package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
	log "github.com/sirupsen/logrus"
)

// InsertHOD will add HOD in table
func InsertHOD(hod HOD) error {
	sqlStatement := `INSERT INTO hod (emp_id, dept_id, start_date) 
						VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, hod.EmpID, hod.DeptID, hod.StartDate)
	return err
}

// GetHodDetails returns hod details
func GetHodDetails(empID string) (HOD, error) {

	hod := HOD{}
	sqlStatement := `SELECT emp_id, dept_id, start_date, end_date FROM hod
						WHERE emp_id = $1`
	err := db.QueryRow(sqlStatement, empID).Scan(
		&hod.EmpID, &hod.DeptID, &hod.StartDate, &hod.EndDate)
	hod.StartDate = util.DateTimeToDate(hod.StartDate)
	hod.EndDate = util.DateTimeToDate(hod.EndDate)
	return hod, err
}

// GetAllHOD returns all hods
func GetAllHOD() ([]*HODDetail, error) {
	hods := make([]*HODDetail, 0)

	rows, err := db.Query(
		`SELECT emp_id, dept_id, start_date, end_date FROM hod`)
	if err != nil {
		return hods, err
	}
	defer rows.Close()

	for rows.Next() {
		var hod HODDetail
		var deptID int

		if err := rows.Scan(&hod.EmpID, &deptID, &hod.StartDate, &hod.EndDate); err == nil {
			hod.StartDate = util.DateTimeToDate(hod.StartDate)
			hod.EndDate = util.DateTimeToDate(hod.EndDate)
			hod.Dept, _ = GetDepartment(deptID)
			hod.EmployeeDetail, _ = GetEmployeeDetails(hod.EmpID)

			hods = append(hods, &hod)
		} else {
			log.Error(err)
		}
	}

	return hods, nil
}

// HOD struct
type HOD struct {
	EmpID     string `json:"empID"`
	DeptID    int    `json:"deptID"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// HODDetail struct for admin page
type HODDetail struct {
	EmpID          string     `json:"empID"`
	Dept           Department `json:"dept"`
	EmployeeDetail Employee
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}
