package main

import (
	"context"
	"fmt"
	"log"
	config "twitter-clone-go/config"
	"twitter-clone-go/repository"
	router "twitter-clone-go/router"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	pool, err := config.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	repository.InitDB(pool)
	defer pool.Close()

	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	router.Run()
}
