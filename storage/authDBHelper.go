package storage

import "strings"

// InsertEmployee add new employee
func InsertEmployee(uname, fname, lname, email, passwd string) error {
	sqlStatement := `INSERT INTO employee (id, first_name, last_name, email, passwd)
							VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, strings.ToLower(uname), fname, lname, email, passwd)
	return err
}

// CheckPasswd checks the user password
func CheckPasswd(uname string) (string, error) {
	var passwd string

	sqlStatement := `SELECT passwd FROM employee WHERE id = $1`
	row := db.QueryRow(sqlStatement, strings.ToLower(uname))
	err := row.Scan(&passwd)

	return passwd, err
}
