package Post

import (
	Utils "burgher/Utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/post/create", Utils.AuthHandler(createHandler)).Methods("Post")
	r.HandleFunc("/post/read", readHandler).Methods("Post", "Get")
	r.HandleFunc("/post/read", readOneHandler).Methods("Post", "Get")
	return r
}
