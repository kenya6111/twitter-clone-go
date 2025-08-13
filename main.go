package main

import (
	"log"
	"twitter-clone-go/api"
	"twitter-clone-go/config"

	_ "github.com/lib/pq"
)

func main() {
	pool, err := config.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	api.Run(pool)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
