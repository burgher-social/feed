package User

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/create", createHandler).Methods("Post")
	r.HandleFunc("/user/read", readHandler).Methods("Post", "Get")
	return r
}
