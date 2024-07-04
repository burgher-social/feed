package User

import (
	"net/http"

	"github.com/gorilla/mux"
)

func authHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	}
}

// func logMW(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

// 		// compare the return-value to the authMW
// 		next.ServeHTTP(w, r)
// 	})
// }

func RegisterRouters() http.Handler {
	r := mux.NewRouter()
	// r.Use(logMW)
	r.HandleFunc("/user/create", createHandler).Methods("Post")
	r.HandleFunc("/user/read", authHandler(readHandler)).Methods("Post", "Get")
	// r.HandleFunc("/user/read", authHandler(readHandler)).Methods("Post", "Get").Subrouter().Use(logMW) // Specific route middleware handler
	return r
}
