package leave

import (
	"errors"
	"github.com/dilip640/Faculty-Portal/storage"
)

func requestLeave(noOfDays int, startDate, comment, empID string) error {
	currLeave, err := storage.GetRemainingLeaves(empID)
	if err != nil {
		return err
	}
	var applier string

	_, err = storage.GetCCFacultyDetails(empID)
	if err == nil {
		applier = "cc_faculty"
	} else if _, err = storage.GetHodDetails(empID); err == nil {
		applier = "hod"
	} else if _, err = storage.GetFacultyDetails(empID); err == nil {
		applier = "faculty"
	} else {
		return errors.New("Nhi milegi")
	}

	if noOfDays <= currLeave {
		err = storage.CreateLeaveApplication(empID, noOfDays, startDate, applier, "INITIALIZED", comment)
	}
	return err
}
