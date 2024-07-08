package Token

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()
	// r.Use(logMW)
	r.HandleFunc("/token/refresh", tokenHandler).Methods("Post", "Get")
	return r
}
