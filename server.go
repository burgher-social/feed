package main

import (
	_ "burgher/Init"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	Feed "burgher/Feed"
	Insights "burgher/Insights"
	Location "burgher/Location"
	Post "burgher/Post"
	Token "burgher/Token"
	Topic "burgher/Topic"
	User "burgher/User"
)

func main() {

	router := mux.NewRouter()
	router.PathPrefix("/user").Handler(User.RegisterRouters())
	router.PathPrefix("/post").Handler(Post.RegisterRouters())
	router.PathPrefix("/location").Handler(Location.RegisterRouters())
	router.PathPrefix("/topic").Handler(Topic.RegisterRouters())
	router.PathPrefix("/feed").Handler(Feed.RegisterRouters())
	router.PathPrefix("/token").Handler(Token.RegisterRouters())
	router.PathPrefix("/insights").Handler(Insights.RegisterRouters())
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
	handler := c.Handler(router)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error while starting server")
	} else {
		fmt.Println("server started")
	}
}
