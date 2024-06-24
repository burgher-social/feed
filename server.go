package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	Feed "burgher/Feed"
	Location "burgher/Location"
	Post "burgher/Post"
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
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error while starting server")
	} else {
		fmt.Println("server started")
	}
}
