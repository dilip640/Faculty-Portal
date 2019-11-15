package leave

import (
	"net/http"
	"strconv"

	"github.com/dilip640/Faculty-Portal/auth"
	log "github.com/sirupsen/logrus"
)

// HandleLeave for leave application
func HandleLeave(w http.ResponseWriter, r *http.Request) {
	userName := auth.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}

	if r.Method == http.MethodPost {
		noOfDays := r.FormValue("no_of_days")
		comment := r.FormValue("comment")
		startDate := r.FormValue("start_date")
		i, err := strconv.Atoi(noOfDays)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = requestLeave(i, startDate, comment, userName)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
