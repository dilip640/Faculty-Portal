package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
	log "github.com/sirupsen/logrus"
)

// InsertCCFaculty add new cc_faculty
func InsertCCFaculty(empID, startDate, postID string) error {
	sqlStatement := `INSERT INTO cc_faculty (emp_id, post_id, start_date)
									VALUES ($1, $2, $3);`
	_, err := db.Exec(sqlStatement, empID, postID, startDate)
	return err
}

// UpdateCCFaculty update cc_faculty
func UpdateCCFaculty(empID, postID string) error {
	sqlStatement := `UPDATE cc_faculty SET post_id = $2
						WHERE emp_id = $1;`
	_, err := db.Exec(sqlStatement, empID, postID)
	return err
}

// GetCCFacultyDetails returns cc_faculty details
func GetCCFacultyDetails(empID string) (CCFaculty, error) {
	var startDate string

	ccFaculty := CCFaculty{}
	sqlStatement := `SELECT emp_id, post_id, start_date FROM cc_faculty
						WHERE emp_id = $1`
	err := db.QueryRow(sqlStatement, empID).Scan(
		&ccFaculty.EmpID, &ccFaculty.Post.PostID, &startDate)
	ccFaculty.StartDate = util.DateTimeToDate(startDate)
	ccFaculty.Post, err = GetPost(ccFaculty.Post.PostID)
	return ccFaculty, err
}

// GetAllccFacultyDetails returns all  the faculties.
func GetAllccFacultyDetails() ([]*CCFacultyDetails, error) {
	ccFaculties := make([]*CCFacultyDetails, 0)

	rows, err := db.Query(`SELECT emp_id, post_id, start_date FROM cc_faculty`)
	if err != nil {
		return ccFaculties, err
	}

	defer rows.Close()

	for rows.Next() {
		var startDate string
		ccFaculty := CCFacultyDetails{}

		if err := rows.Scan(&ccFaculty.EmpID, &ccFaculty.Post.PostID, startDate); err == nil {
			ccFaculty.StartDate = util.DateTimeToDate(startDate)
			ccFaculty.Post, err = GetPost(ccFaculty.Post.PostID)
			ccFaculty.EmployeeDetail, _ = GetEmployeeDetails(ccFaculty.EmpID)

			ccFaculties = append(ccFaculties, &ccFaculty)
		} else {
			log.Error(err)
		}
	}

	return ccFaculties, nil
}

// CCFaculty struct
type CCFaculty struct {
	EmpID     string
	StartDate string
	Post      Post
}

// CCFacultyDetails struct
type CCFacultyDetails struct {
	EmpID          string
	EmployeeDetail Employee
	StartDate      string
	Post           Post
}
