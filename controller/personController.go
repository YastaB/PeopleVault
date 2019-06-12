package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
	"../model"
	"../toolkit"
	"github.com/gorilla/mux"
)

var CreatePerson = func(w http.ResponseWriter, r *http.Request) {
	var aPerson model.Person
	err := json.NewDecoder(r.Body).Decode(&aPerson)
	if err != nil {
		fmt.Println(err)
		toolkit.Respond(w, toolkit.Message(false, "Error while parsing the request body"))
		return
	}
	err = data.CreatePerson(aPerson)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Unable to create person"))
		return
	}

	resp := toolkit.Message(true, "success")
	toolkit.Respond(w, resp)
}

var DeletePerson = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["personID"]

	err := data.DeletePerson(personID)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Cannot find the person"))
		return
	}

	resp := toolkit.Message(true, "success")
	toolkit.Respond(w, resp)
}

var RetrievePerson = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["personID"]
	person, err := data.RetrievePerson(personID)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Cannot find the person"))
		return
	}

	data, err := json.Marshal(&person)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Marshalling error"))
		return
	}

	resp := toolkit.Message(true, "success")
	resp["data"] = string(data)
	toolkit.Respond(w, resp)
}
