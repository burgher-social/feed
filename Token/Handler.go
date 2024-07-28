package Token

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	type TokenRequest struct {
		RefreshToken string `json:"refreshToken"`
	}
	var readRequest TokenRequest
	err := json.NewDecoder(r.Body).Decode(&readRequest)
	fmt.Println(readRequest)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	tokenResponse, err2 := getNewToken(readRequest.RefreshToken)
	if err2 != nil {
		// return HTTP 400 bad request
		w.WriteHeader(401)
		w.Write([]byte(err2.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.
		NewEncoder(w).
		Encode(tokenResponse)

}
