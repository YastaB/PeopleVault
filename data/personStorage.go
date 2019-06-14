package data

import (
	"errors"
	"fmt"

	"../model"
)

// people storage is accessible by the ids.
var peopleStorage = map[string]model.Person{}

func CreatePerson(person model.Person) error {
	if _, ok := peopleStorage[person.PersonID]; ok {
		errorMsg := "Person id: " + person.PersonID + " already exists"
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}
	peopleStorage[person.PersonID] = person
	return nil
}

func DeletePerson(personID string) error {
	if _, ok := peopleStorage[personID]; ok == false {
		errorMsg := "Person id: " + personID + " does not exists"
		fmt.Println(errorMsg)
		return errors.New(errorMsg)
	}
	delete(peopleStorage, personID)
	return nil
}

func RetrievePerson(personID string) (model.Person, error) {
	if _, ok := peopleStorage[personID]; ok == false {
		errorMsg := "Person id: " + personID + " does not exists"
		fmt.Println(errorMsg)
		return model.Person{}, errors.New(errorMsg)
	}
	return peopleStorage[personID], nil
}
