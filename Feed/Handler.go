package Feed

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		UserId string `json:"userId"`
		Offset int    `json:"offset"`
		Limit  int    `json:"limit"`
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
	userfeed, notfound := read(readRequest.UserId, readRequest.Offset, readRequest.Limit)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resplist := []UserFeedResponse{}
	for _, p := range userfeed {
		resplist = append(resplist, UserFeedResponse{PostId: p.PostId, Score: p.Score, Timestamp: p.Timestamp})
	}
	json.NewEncoder(w).Encode(&resplist)
}
