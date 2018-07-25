package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["FName"] = "Gabie"
	session.Values["LName"] = "Senar"
	session.Values["Address"] = "3219 Hannover St Corona CA 92882"
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}

func Check(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get session values.
	fmt.Fprintln(w, session.Name())
	fmt.Fprintln(w, session.Values["FName"])
	fmt.Fprintln(w, session.Values["LName"])
	fmt.Fprintln(w, session.Values["Address"])

}

func main() {
	http.HandleFunc("/", MyHandler)
	http.HandleFunc("/check", Check)
	http.ListenAndServe(":8080", nil)
}
