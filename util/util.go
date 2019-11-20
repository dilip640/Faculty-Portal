package util

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// DateTimeToDate convert datetime to date
func DateTimeToDate(dtstr string) string {
	dt, err := time.Parse("2006-01-02T15:04:05Z", dtstr)
	if err != nil {
		log.Error(err)
	}

	dtstr2 := dt.Format("2006-01-02")
	return dtstr2
}

// GetCurrentYear returns current year
func GetCurrentYear() int16 {
	return int16(time.Now().Year())
}

// DateTToYear convert datet to year
func DateTToYear(dtstr string) int {
	dt, err := time.Parse("2006-01-02", dtstr)
	if err != nil {
		log.Error(err)
	}

	return dt.Year()
}
