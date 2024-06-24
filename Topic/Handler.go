package Topic

import (
	Utils "burgher/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var user TopicRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	userresp, notcreated := create(
		Topic{Id: Utils.GenerateId(), Name: user.Name})
	if notcreated != nil {
		w.WriteHeader(503)
		fmt.Println(notcreated)
		w.Write([]byte(notcreated.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(TopicResponse{Id: userresp.Id, Name: userresp.Name})
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		Name string `json:"name"`
		// Tag      int    `json:"tag"`
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
	userresp, notfound := read(readRequest.Name)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resplist := []TopicResponse{}
	for _, p := range userresp {
		resplist = append(resplist, TopicResponse{Id: p.Id, Name: p.Name})
	}
	json.NewEncoder(w).Encode(&resplist)
}

func readAllHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
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
	userresp, notfound := readAll()

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resplist := []TopicResponse{}
	for _, p := range userresp {
		resplist = append(resplist, TopicResponse{Id: p.Id, Name: p.Name})
	}
	json.NewEncoder(w).Encode(&resplist)
}
