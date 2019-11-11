package storage

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
