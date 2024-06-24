package Topic

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/topic/create", createHandler).Methods("Post")
	r.HandleFunc("/topic/read", readHandler).Methods("Post", "Get")
	r.HandleFunc("/topic/read/all", readAllHandler).Methods("Post", "Get")
	return r
}
