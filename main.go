package main

import (
	"log"
	config "twitter-clone-go/config"
	"twitter-clone-go/controllers"
	"twitter-clone-go/repository"
	router "twitter-clone-go/router"
	"twitter-clone-go/services"

	_ "github.com/lib/pq"
)

func main() {
	pool, err := config.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewMyAppRepository(pool)
	svc := services.NewMyAppService(repo)
	con := controllers.NewMyAppController(svc)

	router.Run(con)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
