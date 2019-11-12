package storage

import (
	"github.com/dilip640/Faculty-Portal/util"
)

func GetRemainingLeaves(empID string) (int, error) {
	sqlStatement := `SELECT get_remaining_leave($1, $2)`

	var numLeaves int
	err := db.QueryRow(sqlStatement, empID, util.GetCurrentYear()).Scan(&numLeaves)

	return numLeaves, err
}
