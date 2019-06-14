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
	if _, ok := peopleStorage[personID]; !ok {
		errorMsg := "Person id: " + personID + " does not exists"
		fmt.Println(errorMsg)
		return model.Person{}, errors.New(errorMsg)
	}
	return peopleStorage[personID], nil
}

//RetrievePeopleWithName if any of the parameters are empty only query with non empty fields
func RetrievePeopleWithName(firstName string, lastName string) ([]model.Person, error) {
	foundedPeople := []model.Person{}
	for _, v := range peopleStorage {
		if (firstName != "" && lastName != "" && firstName == v.FirstName && lastName == v.LastName) ||
			(firstName == "" && lastName != "" && lastName == v.LastName) ||
			(firstName != "" && lastName == "" && firstName == v.FirstName) {
			foundedPeople = append(foundedPeople, v)
		}
	}
	if len(foundedPeople) == 0 {
		return foundedPeople, errors.New("specified person does not exists")
	}
	return foundedPeople, nil
}
