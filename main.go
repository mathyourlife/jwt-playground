package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)

	// Initialize the router
	r := mux.NewRouter()

	// Main page accessible without a login
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Login request. accessible without a login but will look for basic auth
	r.Handle("/login", PostLoginHandler).Methods("POST")
	// Resource available to a logged in user with a valid JWT
	r.Handle("/my-page", loggedInHandler(UserPageHandler)).Methods("GET")

	// start our server
	http.ListenAndServe(":5000", handlers.CombinedLoggingHandler(os.Stdout, r))
}

var UserPageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This page shows resources for a logged in user"))
})
