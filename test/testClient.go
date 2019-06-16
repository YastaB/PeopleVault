package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("client1signingKey")

const baseURL = "http://localhost:8001/api/v1/person"
const CLIENT_ID_LABEL = "Clientid"
const TOKEN_LABEL = "Token"

func addPerson(clientID string, newPerson *model.Person) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}

	data, err := json.Marshal(newPerson)

	req, _ := http.NewRequest("POST", baseURL, bytes.NewBuffer(data))
	req.Header.Set(TOKEN_LABEL, validToken)
	req.Header.Set(CLIENT_ID_LABEL, clientID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func getPerson(clientID string, personID string) string {
	return performGetRequest(clientID, baseURL+"/"+personID)
}

func getPeopleWithName(clientID string, firstName string, lastName string) string {
	return performGetRequest(clientID, baseURL+"/queryname"+"?firstName="+firstName+"&lastName="+lastName)
}
func getPeopleWithAgeRange(clientID string, minAge int, maxAge int) string {
	return performGetRequest(clientID, baseURL+"/queryage"+"?minAge="+strconv.Itoa(minAge)+"&maxAge="+strconv.Itoa(maxAge))
}

func performGetRequest(clientID string, url string) string {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
		return ""
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set(TOKEN_LABEL, validToken)
	req.Header.Set(CLIENT_ID_LABEL, clientID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	newPerson := &model.Person{
		PersonID:  "1",
		FirstName: "Bob",
		LastName:  "Fisher",
		Address:   "USA",
		Age:       24}

	newPerson2 := &model.Person{
		PersonID:  "2",
		FirstName: "Berkay",
		LastName:  "Tacyildiz",
		Address:   "Turkey",
		Age:       26}
	addPerson("client1", newPerson)
	addPerson("client1", newPerson2)
	getPerson("client1", "1")
	getPerson("client1", "2")
	getPeopleWithName("client1", "Bob", "Fisher")
	getPeopleWithAgeRange("client1", 24, 26)
}
