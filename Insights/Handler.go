package Insights

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func likeHandler(w http.ResponseWriter, r *http.Request) {
	type LikeRequest struct {
		Count  int    `json:"count"`
		PostId string `json:"postId"`
	}
	var likeReq LikeRequest
	err := json.NewDecoder(r.Body).Decode(&likeReq)
	// ctx := r.Context()
	// userId := ctx.Value(Utils.ContextUserKey)
	// locationReq.UserId = userId.(string)
	// fmt.Println(locationReq)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err2 := Like(likeReq.Count, likeReq.PostId)
	if err2 != nil {
		w.WriteHeader(503)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}
