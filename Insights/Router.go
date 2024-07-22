package Insights

import (
	Utils "burgher/Utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/insights/like/add", Utils.AuthHandler(likeHandler)).Methods("POST")
	return r
}
