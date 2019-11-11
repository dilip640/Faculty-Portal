package auth

import (
	"net/http"

	"github.com/dilip640/Faculty-Portal/storage"
	"github.com/dilip640/Faculty-Portal/templatemanager"
	log "github.com/sirupsen/logrus"
)

// HandleLogin handle the greeeting
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if GetUserName(r) != "" {
		http.Redirect(w, r, "/profile", 302)
	}

	data := struct {
		Error string
	}{}

	name := r.FormValue("name")
	pass := r.FormValue("password")
	if name != "" && pass != "" {
		passwd, err := storage.CheckPasswd(name)
		if err != nil {
			log.Error(err)
			data.Error = "Invalid Username or Password!"
		} else {
			exist := comparePasswords(passwd, pass)
			if exist {
				setSession(name, w)
				http.Redirect(w, r, "/profile", 302)
			} else {
				data.Error = "Invalid Username or Password!"
			}
		}

	}

	templatemanager.Render(w, GetUserName(r), data, "base", "templates/login.html", "templates/base.html")
}

// HandleRegister handle the greeeting
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if GetUserName(r) != "" {
		http.Redirect(w, r, "/profile", 302)
	}

	data := struct {
		Error string
	}{}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		pass := r.FormValue("password")
		fname := r.FormValue("first_name")
		lname := r.FormValue("last_name")
		email := r.FormValue("email")
		if name != "" && pass != "" && fname != "" && lname != "" && email != "" {
			err := storage.InsertEmployee(name, fname, lname, email, hashAndSalt(pass))
			if err == nil {
				setSession(name, w)
				http.Redirect(w, r, "/profile", 302)
			} else {
				log.Error(err)
				data.Error = "Username Taken!"
			}
		} else {
			data.Error = "Enter Correct Details!"
		}
	}
	templatemanager.Render(w, GetUserName(r), data, "base", "templates/register.html", "templates/base.html")
}

// HandleLogout logout the user
func HandleLogout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}
