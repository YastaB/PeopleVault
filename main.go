package main

import (
	"fmt"
	"net/http"

	"./controller"

	"github.com/gorilla/mux"
)

// this is main entry of the server
func main() {
	const PORT = "8000"
	const personAPILabel = "/api/v1/person"

	router := mux.NewRouter()
	router.HandleFunc(personAPILabel, controller.CreatePerson).Methods("POST")
	router.HandleFunc(personAPILabel+"/{personID}", controller.DeletePerson).Methods("DELETE")
	router.HandleFunc(personAPILabel+"/{personID}", controller.RetrievePerson).Methods("GET")

	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		fmt.Print(err)
	}
}
