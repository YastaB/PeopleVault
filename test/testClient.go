package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../model"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("client1signingKey")

const baseURL = "http://localhost:8001/api/v1/person"
const CLIENT_ID_LABEL = "Clientid"
const TOKEN_LABEL = "Token"

func addPerson(clientID string) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	newPerson := &model.Person{
		PersonID:  "1",
		FirstName: "Bob",
		LastName:  "Fisher",
		Address:   "USA",
		Age:       24}

	data, err := json.Marshal(&newPerson)

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

func getPerson(personID string, clientID string) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", baseURL+"/"+personID, nil)
	req.Header.Set(TOKEN_LABEL, validToken)
	req.Header.Set(CLIENT_ID_LABEL, clientID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	addPerson("client1")
	//getPerson("1", "client1")
}
