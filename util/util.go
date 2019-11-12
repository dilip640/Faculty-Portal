package util

import "time"

// DateTimeToDate convert datetime to date
func DateTimeToDate(dtstr string) string {
	dt, _ := time.Parse("2006-01-02T15:04:05Z", dtstr)

	dtstr2 := dt.Format("02-01-2006")
	return dtstr2
}

// GetCurrentYear returns current year
func GetCurrentYear() int16 {
	return int16(time.Now().Year())
}
