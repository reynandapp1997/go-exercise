package main

import (
	"crud-golang-mongodb/controllers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user", controllers.HandleUsersList).Methods("GET")
	r.HandleFunc("/api/user/{id}", controllers.HandleUserSingle).Methods("GET")
	r.HandleFunc("/api/user", controllers.HandleInsertUser).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
