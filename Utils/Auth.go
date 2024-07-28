package Utils

import (
	"burgher/Token"
	"context"
	"fmt"
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = "userId"

func UserFromContext(ctx context.Context) string {
	return ctx.Value(ContextUserKey).(string)
}

func AuthHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(401)
			w.Write([]byte("Unautorized"))
			return
		}
		claims, err := Token.GetTokenClaims(authHeader)
		fmt.Println("TOKEN claims")
		fmt.Println(claims)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("Unautorized"))
			return
		}
		ctx := r.Context()
		req := r.WithContext(context.WithValue(ctx, ContextUserKey, claims.UserId))
		*r = *req
		next.ServeHTTP(w, r)
	}
}
