package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
// User モデル
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!!"})
}
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}


func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM users")
	
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}
	defer rows.Close()
}

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
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
	fmt.Println(dsn)
	var err error
	db, err = setupDB(dbDriver, dsn)
	fmt.Println(db)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/",home)
	router.GET("/users",getUsers)
	router.GET("/health_check",healthCheck)
	router.Run()
}