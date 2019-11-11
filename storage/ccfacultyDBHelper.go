package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
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

// CCFaculty struct
type CCFaculty struct {
	EmpID     string
	StartDate string
	Post      Post
}
