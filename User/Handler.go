package User

import (
	Utils "burgher/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var user UserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, notcreated := create(
		User{Id: Utils.GenerateId(), UserName: user.UserName, Name: user.Name, Tag: user.Tag})
	if notcreated != nil {
		w.WriteHeader(503)
		fmt.Println(notcreated)
		w.Write([]byte(notcreated.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(UserResponse{Id: userresp.Id, Name: userresp.Name, UserName: userresp.UserName, Tag: userresp.Tag, IsVerified: userresp.IsVerified})
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		Username string `json:"username"`
		Tag      int    `json:"tag"`
	}
	var readRequest ReadRequest
	err := json.NewDecoder(r.Body).Decode(&readRequest)
	fmt.Println(readRequest)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, notfound := read(readRequest.Username, readRequest.Tag)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(UserResponse{Id: userresp.Id, Name: userresp.Name, UserName: userresp.UserName, Tag: userresp.Tag, IsVerified: userresp.IsVerified})
}
