package main

import (
	"log"
	"twitter-clone-go/application"
	"twitter-clone-go/infrastructure/email/mailcatcher"
	"twitter-clone-go/infrastructure/storage/postgres"
	"twitter-clone-go/interface/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// DB接続
	db, err := postgres.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var emailService = mailcatcher.NewMainCatcherEmailService("temp")
	// トランザクションの注入
	tx := postgres.NewTransaction(db)

	// リポジトリの注入
	repo := postgres.NewUserRepository(db)

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
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
