package Feed

import (
	Utils "burgher/Utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	// r.HandleFunc("/location/create", createHandler).Methods("POST")
	r.HandleFunc("/feed/read", Utils.AuthHandler(readHandler)).Methods("POST", "GET")
	return r
}
