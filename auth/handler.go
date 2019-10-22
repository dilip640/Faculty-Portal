package auth

import (
	"fmt"
	"net/http"
)

// HandleGreeting handle the greeeting
func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
