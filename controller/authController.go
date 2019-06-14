package controller

import (
	"fmt"
	"net/http"

	data "../data"
	u "../toolkit"
	jwt "github.com/dgrijalva/jwt-go"
)

const CLIENT_ID_LABEL = "Clientid"
const TOKEN_LABEL = "Token"

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header[TOKEN_LABEL] == nil {
			u.ReturnHttpError("Missing auth token", http.StatusForbidden, w)
			return
		}

		if r.Header[CLIENT_ID_LABEL] == nil {
			u.ReturnHttpError("Missing client id", http.StatusForbidden, w)
			return
		}

		clientID := r.Header[CLIENT_ID_LABEL][0]

		token, err := jwt.Parse(r.Header[TOKEN_LABEL][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			if _, ok := data.ClientSigningKeys[clientID]; !ok {
				return nil, fmt.Errorf("Client ID is not founded")
			}

			return []byte(data.ClientSigningKeys[clientID]), nil
		})

		if err != nil {
			u.ReturnHttpError("Error parsing auth token", http.StatusForbidden, w)
			return
		}

		if token.Valid == false {
			u.ReturnHttpError("Invalid auth token", http.StatusForbidden, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
