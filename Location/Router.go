package Location

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/location/create", createHandler).Methods("POST")
	r.HandleFunc("/location/read", readHandler).Methods("POST", "GET")
	return r
}
