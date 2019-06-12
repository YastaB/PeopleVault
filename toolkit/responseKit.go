package toolkit

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, responseMsg string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": responseMsg}
}

func Respond(w http.ResponseWriter, responseData map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
