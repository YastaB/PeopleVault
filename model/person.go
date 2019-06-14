package model

// Person struct holds the basic info in order to indentify stored people
type Person struct {
	PersonID  string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       uint   `json:"age"`
	Address   string `json:"adress"`
}
