package Insights

import (
	Utils "burgher/Utils"
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
	userId := r.Context().Value(Utils.ContextUserKey).(string)
	fmt.Printf("%+v\n", likeReq)
	// fmt.Println(locationReq)
	if err != nil {
		// return HTTP 400 bad request
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	var err2 error
	if likeReq.Count > 0 {
		err2 = Like(likeReq.Count, likeReq.PostId, userId)
	} else {
		err2 = UnLike(likeReq.Count, likeReq.PostId, userId)
	}

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
