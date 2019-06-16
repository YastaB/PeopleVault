package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	// check for negative age value
	if aPerson.Age < 0 {
		toolkit.Respond(w, toolkit.Message(false, "Age value cannot be negative"))
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

var RetrievePeopleWithName = func(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	firstName := queryParams.Get("firstName")
	lastName := queryParams.Get("lastName")

	if firstName == "" && lastName == "" {
		toolkit.Respond(w, toolkit.Message(false, "invalid parameters for querying with name"))
		return
	}

	fmt.Println("Querying people with FirstName: " + firstName + " LastName: " + lastName)

	peopleList, err := data.RetrievePeopleWithName(firstName, lastName)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Cannot find the person"))
		return
	}

	data, err := json.Marshal(&peopleList)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Marshalling error"))
		return
	}

	resp := toolkit.Message(true, "success")
	resp["data"] = string(data)
	toolkit.Respond(w, resp)
}

var RetrievePeopleWithAgeRange = func(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	minAge := queryParams.Get("minAge")
	maxAge := queryParams.Get("maxAge")

	if minAge == "" || maxAge == "" {
		toolkit.Respond(w, toolkit.Message(false, "invalid parameters for querying with age range"))
		return
	}

	// conver age values to integer
	minAgeInt, err1 := strconv.Atoi(minAge)
	maxAgeInt, err2 := strconv.Atoi(maxAge)
	if err1 != nil || err2 != nil {
		toolkit.Respond(w, toolkit.Message(false, "invalid parameters for querying with age range"))
		return
	}

	if minAgeInt < 0 || maxAgeInt < 0 {
		toolkit.Respond(w, toolkit.Message(false, "age value cannot be negative"))
		return
	}

	fmt.Println("Querying people with MinAge: " + minAge + " MaxAge: " + maxAge)
	peopleList, err := data.RetrievePeopleWithAgeRange(minAgeInt, maxAgeInt)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Cannot find the person"))
		return
	}

	data, err := json.Marshal(&peopleList)
	if err != nil {
		toolkit.Respond(w, toolkit.Message(false, "Marshalling error"))
		return
	}

	resp := toolkit.Message(true, "success")
	resp["data"] = string(data)
	toolkit.Respond(w, resp)
}
