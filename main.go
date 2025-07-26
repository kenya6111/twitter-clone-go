package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"twitter-clone-go/repository"
	router "twitter-clone-go/router"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var pool *pgxpool.Pool
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
	Password string `json:"password"`
}


func setupDB(dbDriver string, dsn string) (*pgxpool.Pool, error) {

	db, err := pgxpool.New(context.Background(), dsn)
	return db, err
}

func main(){

	dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
	dbDriver := "postgres"
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	pool, err = setupDB(dbDriver, dsn)

	repository.InitDB(pool)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()


	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	router.Run()
}