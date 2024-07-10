package User

import (
	"net/http"

	Utils "burgher/Utils"

	"github.com/gorilla/mux"
)

// type ContextKey string

// const ContextUserKey ContextKey = "userId"

// func UserFromContext(ctx context.Context) string {
// 	return ctx.Value(ContextUserKey).(string)
// }
// func authHandler(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Our middleware logic goes here...
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			w.WriteHeader(401)
// 			w.Write([]byte("Unautorized"))
// 			return
// 		}
// 		claims, err := Token.GetTokenClaims(authHeader)
// 		if err != nil {
// 			w.WriteHeader(401)
// 			w.Write([]byte("Unautorized"))
// 			return
// 		}
// 		ctx := r.Context()
// 		req := r.WithContext(context.WithValue(ctx, ContextUserKey, claims.UserId))
// 		*r = *req
// 		next.ServeHTTP(w, r)
// 	}
// }

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
	r.HandleFunc("/user/read", Utils.AuthHandler(readHandler)).Methods("Post", "Get")
	r.HandleFunc("/user/profile/image/upload", Utils.AuthHandler(imageHandler)).Methods("Post", "Get")
	r.HandleFunc("/user/read/email", readWithEmailHandler).Methods("Post", "Get")
	// r.HandleFunc("/user/read", authHandler(readHandler)).Methods("Post", "Get").Subrouter().Use(logMW) // Specific route middleware handler
	return r
}
