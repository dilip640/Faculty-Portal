package storage

// InsertDepartment add new dept
func InsertDepartment(deptName string) error {
	sqlStatement := `INSERT INTO department (dept_name) VALUES ($1);`
	_, err := db.Exec(sqlStatement, deptName)
	return err
}

// DeleteDepartment remove dept
func DeleteDepartment(deptID string) error {
	sqlStatement := `DELETE FROM department WHERE id = $1;`
	_, err := db.Exec(sqlStatement, deptID)
	return err
}

// GetAllDepartments returns all depts
func GetAllDepartments() ([]*Department, error) {
	depts := make([]*Department, 0)

	rows, err := db.Query(
		`SELECT id, dept_name FROM department`)
	if err != nil {
		return depts, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			deptID   string
			deptName string
		)

		if err := rows.Scan(&deptID, &deptName); err != nil {
			return depts, err
		}
		depts = append(depts, &Department{deptID, deptName})
	}

	return depts, nil
}

// Department struct
type Department struct {
	DeptID string
	Name   string
}
