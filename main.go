package main

import (
	"log"
	"twitter-clone-go/application"
	"twitter-clone-go/infrastructure/email/mailcatcher"
	bcrypt "twitter-clone-go/infrastructure/password_hasher"
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
	transaction := postgres.NewTransaction(db)

	// リポジトリの注入
	userRepo := postgres.NewUserRepository(db)
	emailVerifyRepo := postgres.NewEmailVerifyRepository(db)

	// サービスの注入
	userDomainService := application.NewUserDomainService(userRepo)
	passwordHasher := bcrypt.NewBcryptHasher()

	// ユースケースの注入
	ser := application.NewUserUsecase(userRepo, emailVerifyRepo, transaction, userDomainService, emailService, passwordHasher)

	// ハンドラーの注入
	con := http.NewUserHandler(ser)

	store := memstore.NewStore([]byte("secret"))

	router := gin.Default()
	router.Use(sessions.Sessions("mySession", store))
	router.GET("/", con.Home)
	router.GET("/users", con.GetUserList)
	router.POST("/signup", con.SignUp)
	router.POST("/activate", con.Activate)
	router.GET("/health_check", con.HealthCheck)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
