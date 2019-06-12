package data

import (
	"errors"
	"fmt"

	"../model"
)

// people storage is accessible by the ids.
var people_storage = map[string]model.Person{}

func CreatePerson(person model.Person) error {
	if _, ok := people_storage[person.PersonID]; ok {
		errorMsg := "Person id: " + person.PersonID + " already exists"
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}
	people_storage[person.PersonID] = person
	return nil
}

func DeletePerson(personID string) error {
	if _, ok := people_storage[personID]; ok == false {
		errorMsg := "Person id: " + personID + " does not exists"
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}
	delete(people_storage, personID)
	return nil
}

func RetrievePerson(personID string) (model.Person, error) {
	if _, ok := people_storage[personID]; ok == false {
		errorMsg := "Person id: " + personID + " does not exists"
		fmt.Println(errorMsg)
		return model.Person{}, errors.New(errorMsg)
	}
	return people_storage[personID], nil
}
