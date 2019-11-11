package util

import "time"

// DateTimeToDate convert datetime to date
func DateTimeToDate(dtstr string) string {
	dt, _ := time.Parse("2006-01-02T15:04:05Z", dtstr)

	dtstr2 := dt.Format("02-01-2006")
	return dtstr2
}
