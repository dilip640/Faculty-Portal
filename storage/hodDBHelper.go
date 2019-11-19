package storage

import "github.com/dilip640/Faculty-Portal/util"

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

// HOD struct
type HOD struct {
	EmpID     string `json:"empID"`
	DeptID    int    `json:"deptID"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
