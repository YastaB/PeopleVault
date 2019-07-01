
# PEOPLE VAULT REST SERVICE

This rest service provides serveral APIs to manage people information. 

## Getting Started

After the project has been cloned, in the project directory,<br/>
Please call the following two commands in order to install the dependencies. 
go get github.com/gorilla/mux <br/>
go get github.com/dgrijalva/jwt-go <br/>

<br/>
After every dependencies are installed please call <br/>
go run main.go <br/>
In order to start the server.

### Prerequisites

Golang version: go1.10.7 darwin/amd64<br/>
gorilla/mux: Router capability for the rest service<br/>
dgrijalva/jwt-go: authentication with json web tokens<br/>

This project uses port "8001" as default <br/>
please make sure that this port is not already in use.

## Running the tests

In order to interact with the APIs, under the test directory, testClient.go is implemented. <br/>
Currently, there are two clients which can authorized to interact with the API that are predefined under /data/clientStorage.go. <br/>
Sining keys defined in the data folder is only for development and testing purposes. In production, please use other database solutions that can be used to store client signing keys (which can utilized salted hashes). 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
