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

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
