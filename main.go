package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/users",GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}",GetUser).Methods("GET")
	r.HandleFunc("/users",CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}",UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081",r))
}

func main() {
	 InitialMigration()
	 initializeRouter()
}