package util

import "time"

// DateTimeToDate convert datetime to date
func DateTimeToDate(dtstr string) string {
	dt, _ := time.Parse("2006-01-02 15:04:05", dtstr)

	dtstr2 := dt.Format("2006-01-02")
	return dtstr2
}
