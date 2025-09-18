package main

import (
	"log"
	application "twitter-clone-go/application/user"
	"twitter-clone-go/infrasctructure/email/mailcatcher"
	"twitter-clone-go/infrasctructure/storage/postgres"
	"twitter-clone-go/interface/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// DB接続
	pool, err := postgres.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	var emailService = mailcatcher.NewMainCatcherEmailService("temp")
	// トランザクションの注入
	tx := postgres.NewTransaction(pool)

	// リポジトリの注入
	repo := postgres.NewUserRepository(pool)

	// サービスの注入
	dSer := application.NewUserDomainService(repo)

	// ユースケースの注入
	ser := application.NewUserUsecase(repo, tx, dSer, emailService)

	// ハンドラーの注入
	con := http.NewUserHandler(ser)

	store := memstore.NewStore([]byte("secret"))

	router := gin.Default()
	router.Use(sessions.Sessions("mySession", store))
	router.GET("/", con.Home)
	router.GET("/users", con.GetUserListHandler)
	router.POST("/signup", con.SignUpHandler)
	router.GET("/health_check", con.HealthCheck)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
