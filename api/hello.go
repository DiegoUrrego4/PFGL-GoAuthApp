package api

import (
	"encoding/json"
	"fmt"
	"github.com/goschool/crud/types"
	"net/http"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func HandleEchoUsers(w http.ResponseWriter, r *http.Request) {
	var user types.User

	// Decode the request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "this is the email: %s", user.Email)
}
