package main

import (
	"log"
	"twitter-clone-go/application"
	"twitter-clone-go/infrastructure/email/mailcatcher"
	bcrypt "twitter-clone-go/infrastructure/password_hasher"
	"twitter-clone-go/infrastructure/session_store"
	"twitter-clone-go/infrastructure/storage/postgres"
	"twitter-clone-go/interface/http"

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
	sessionStore := session_store.NewSessionStore()

	// ユースケースの注入
	ser := application.NewUserUsecase(userRepo, emailVerifyRepo, transaction, userDomainService, emailService, passwordHasher, sessionStore)

	// ハンドラーの注入
	con := http.NewUserHandler(ser)

	router := gin.Default()
	router.GET("/", con.Home)
	router.POST("/signup", con.SignUp)
	router.POST("/activate", con.Activate)
	router.POST("/login", con.Login)
	router.POST("/logout", con.Logout)
	router.GET("/health_check", con.HealthCheck)

	loginCheckGroup := router.Group("/")
	loginCheckGroup.Use(http.CheckLogin(sessionStore))
	loginCheckGroup.GET("/users", con.GetUserList)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
