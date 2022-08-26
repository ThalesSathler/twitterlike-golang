package main

import (
	"log"

	"github.com/thalessathler/twitterlike/internal/auth"
	"github.com/thalessathler/twitterlike/internal/server"
	"github.com/thalessathler/twitterlike/internal/tweet"
	"github.com/thalessathler/twitterlike/internal/user"
	"github.com/thalessathler/twitterlike/repository/inmemorydatabase"
)

func main() {
	tweetRepo := inmemorydatabase.NewTweetRepository()
	tweetsvc, _ := tweet.New(tweetRepo, nil)
	userRepo := inmemorydatabase.NewUserRepository()
	usersvc := user.New(userRepo)
	authsvc := auth.New(usersvc)

	cfg := &server.Config{
		Port: ":8080",
	}
	server := server.NewServer(cfg, tweetsvc, usersvc, authsvc)

	log.Println("Starting Server")
	log.Fatal(server.ListenAndServe())
}
