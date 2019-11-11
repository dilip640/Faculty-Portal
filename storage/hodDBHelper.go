package storage

// InsertHOD will add HOD in table
func InsertHOD(hod HOD) error {
	sqlStatement := `INSERT INTO hod (emp_id, dept_id, start_date) 
						VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, hod.EmpID, hod.DeptID, hod.StartDate)
	return err
}

// HOD struct
type HOD struct {
	EmpID     string `json:"empID"`
	DeptID    string `json:"deptID"`
	StartDate string `json:"startDate"`
}
