package Post

import (
	"burgher/Utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/post/create", Utils.AuthHandler(createHandler)).Methods("Post")
	r.HandleFunc("/post/read", readHandler).Methods("Post", "Get")
	r.HandleFunc("/post/readOne", readOneHandler).Methods("Post", "Get")
	r.HandleFunc("/post/comments/read", readCommentHandler).Methods("Post", "Get")

	return r
}
