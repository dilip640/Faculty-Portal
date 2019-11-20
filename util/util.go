package util

import (
	"time"
)

// DateTimeToDate convert datetime to date
func DateTimeToDate(dtstr string) string {
	dt, _ := time.Parse("2006-01-02T15:04:05Z", dtstr)

	dtstr2 := dt.Format("2006-01-02")
	return dtstr2
}

// GetCurrentYear returns current year
func GetCurrentYear() int16 {
	return int16(time.Now().Year())
}

// DateTToYear convert datet to year
func DateTToYear(dtstr string) int {
	dt, _ := time.Parse("2006-01-02", dtstr)

	return dt.Year()
}
