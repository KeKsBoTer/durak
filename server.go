package durak

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func makeRouter() *mux.Router {
	router := mux.NewRouter()	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "It works")
	})
	return router
}

func Start() {
	router := makeRouter()
	http.ListenAndServe(":8080", router)
}
