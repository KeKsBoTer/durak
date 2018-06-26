package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if len(username) == 0 {
		http.Error(w, "missing username", http.StatusBadRequest)
		return
	}
	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "player")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		session.Values["name"] = username
	}
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)
}

func username(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "player")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.IsNew {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	username := (session.Values["name"]).(string)
	fmt.Fprint(w, username)
}
