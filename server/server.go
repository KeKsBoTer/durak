package server

import (
	"net/http"

	"github.com/KeKsBoTer/durak/entity"

	"github.com/gorilla/mux"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}

func makeRouter() *mux.Router {
	hub := newHub()
	go hub.run()
	router := mux.NewRouter()
	router.Use(cors)
	router.HandleFunc("/login", login)
	router.HandleFunc("/username", username)
	router.HandleFunc("/ws", ws)
	router.HandleFunc("/queue", func(w http.ResponseWriter, r *http.Request) {
		queue(hub, w, r)
	})
	router.StrictSlash(true)
	return router
}

// Start starts web server
func Start() {
	router := makeRouter()
	http.ListenAndServe(":8080", router)
}

func startGame(users []entity.User) {

}
