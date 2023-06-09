package main

import (
	"gormlogin/model"
	"gormlogin/server"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	model.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/", server.HandlePeople).Methods("GET")
	router.HandleFunc("/create", server.HandleCreatePerson).Methods("POST")
	router.HandleFunc("/delete/{id}", server.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/find/{id}", server.PersonFinder).Methods("GET")
	router.HandleFunc("/find/{name}/{password}", server.FinderByNameAndPassword).Methods("GET")
	router.HandleFunc("/pagination/{page}/{limit}", server.PaginationHandler).Methods("GET")
	// Enable CORS for all routes
	corsRouter := handlers.CORS()(router)
	http.ListenAndServe(":8080", corsRouter)

}
