package storage

import log "github.com/sirupsen/logrus"

// InsertDepartment add new dept
func InsertDepartment(deptName string) error {
	sqlStatement := `INSERT INTO department (dept_name) VALUES ($1);`
	_, err := db.Exec(sqlStatement, deptName)
	return err
}

// DeleteDepartment remove dept
func DeleteDepartment(deptID int) error {
	sqlStatement := `DELETE FROM department WHERE id = $1;`
	_, err := db.Exec(sqlStatement, deptID)
	return err
}

// GetDepartment returns one dept
func GetDepartment(deptID int) (Department, error) {
	dept := Department{}
	sqlStatement := `SELECT id, dept_name FROM department
						WHERE id = $1`
	err := db.QueryRow(sqlStatement, deptID).Scan(
		&dept.DeptID, &dept.Name)

	return dept, err
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
			deptID   int
			deptName string
		)

		if err := rows.Scan(&deptID, &deptName); err == nil {
			depts = append(depts, &Department{deptID, deptName})
		} else {
			log.Error(err)
		}
	}

	return depts, nil
}

// Department struct
type Department struct {
	DeptID int
	Name   string
}
