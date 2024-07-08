package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	Feed "burgher/Feed"
	Location "burgher/Location"
	Post "burgher/Post"
	Token "burgher/Token"
	Topic "burgher/Topic"
	User "burgher/User"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	router := mux.NewRouter()
	router.PathPrefix("/user").Handler(User.RegisterRouters())
	router.PathPrefix("/post").Handler(Post.RegisterRouters())
	router.PathPrefix("/location").Handler(Location.RegisterRouters())
	router.PathPrefix("/topic").Handler(Topic.RegisterRouters())
	router.PathPrefix("/feed").Handler(Feed.RegisterRouters())
	router.PathPrefix("/token").Handler(Token.RegisterRouters())
	handler := cors.Default().Handler(router)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error while starting server")
	} else {
		fmt.Println("server started")
	}
}
