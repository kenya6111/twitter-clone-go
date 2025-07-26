package repository

import (
	"context"
	"fmt"
	"net/http"

	db "twitter-clone-go/tutorial"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)
var pool *pgxpool.Pool
func InitDB(p *pgxpool.Pool){
	pool=p
}
func SelectUsers(c *gin.Context) (*[]db.User, error){
	q := db.New(pool)
	fmt.Println(pool)
	fmt.Println("------------------------------")
	fmt.Println(q)
	fmt.Println("------------------------------")
	resultSet, err := q.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	fmt.Println("sqlc導入成功！")
	c.JSON(http.StatusOK, resultSet)

	return &resultSet,nil
}

func GetUserByEmail(c *gin.Context, email string) (*db.User, error) {
	q := db.New(pool)
	resultSet, err := q.GetUserByEmail(context.Background(),email)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return &resultSet,nil
}

func CreateUser (ctx context.Context, email string, hash []byte) bool {
	q := db.New(pool)
	userInfo := db.CreateUserParams{
		Name:email,
		Email:email,
		Password:string(hash),
	}
	resultSet, err := q.CreateUser(context.Background(), userInfo)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(resultSet)
	return true
}