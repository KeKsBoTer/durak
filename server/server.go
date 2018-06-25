package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It works")
}

func makeRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/ws", websocket)
	return router
}

// Start starts web server
func Start() {
	router := makeRouter()
	http.ListenAndServe(":8080", router)
}
