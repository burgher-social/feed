package Post

import (
	Utils "burgher/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var post PostRequest
	err := json.NewDecoder(r.Body).Decode(&post)
	fmt.Println(post)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	posts, notcreated := create(
		Post{Id: Utils.GenerateId(), UserId: post.UserId, Content: post.Content, ParentId: post.ParentId}, post.Topics)
	if notcreated != nil {
		w.WriteHeader(503)
		fmt.Println(notcreated)
		w.Write([]byte(notcreated.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(PostResponse{Id: posts.Id, ParentId: posts.ParentId, Content: posts.Content, UserId: posts.UserId})
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		UserId string `json:"userId"`
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
	posts, notfound := Read(readRequest.UserId)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	resplist := []PostResponse{}
	for _, p := range posts {
		resplist = append(resplist, PostResponse{Id: p.Id, Content: p.Content, UserId: p.UserId, ParentId: p.ParentId})
	}
	json.NewEncoder(w).Encode(&resplist)
}

func readOneHandler(w http.ResponseWriter, r *http.Request) {
	type ReadRequest struct {
		PostId string `json:"postId"`
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
	post, notfound := ReadOne(readRequest.PostId)

	if notfound != nil {
		w.WriteHeader(503)
		fmt.Println(notfound)
		w.Write([]byte(notfound.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	// resplist := []PostInfo{}
	// for _, p := range post {
	// 	resplist = append(resplist, PostResponse{Id: p.Id, Content: p.Content, UserId: p.UserId, ParentId: p.ParentId})
	// }
	json.NewEncoder(w).Encode(&post)
}
