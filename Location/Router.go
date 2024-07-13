package Location

import (
	Utils "burgher/Utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/location/create", Utils.AuthHandler(createHandler)).Methods("POST")
	r.HandleFunc("/location/read", readHandler).Methods("POST", "GET")
	return r
}
