package main

import (
	"log"
	"twitter-clone-go/application"
	"twitter-clone-go/infrastructure/email/mailcatcher"
	bcrypt "twitter-clone-go/infrastructure/password_hasher"
	"twitter-clone-go/infrastructure/session_store"
	"twitter-clone-go/infrastructure/storage/file"
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

	tweetRepo := postgres.NewTweetRepository(db)
	tweetImageRepo := postgres.NewTweetImageRepository(db)

	// サービスの注入
	userDomainService := application.NewUserDomainService(userRepo)
	passwordHasher := bcrypt.NewBcryptHasher()
	sessionStore := session_store.NewSessionStore()
	fileUploadService := file.NewLocalFileUploader()

	// ユースケースの注入
	userUscase := application.NewUserUsecase(userRepo, emailVerifyRepo, transaction, userDomainService, emailService, passwordHasher, sessionStore)

	tweetUsecase := application.NewTweetUsecase(tweetRepo, tweetImageRepo, transaction, fileUploadService)

	// ハンドラーの注入
	userCon := http.NewUserHandler(userUscase)
	tweetCon := http.NewTweetHandler(tweetUsecase)

	router := gin.Default()
	router.GET("/", userCon.Home)
	router.POST("/signup", userCon.SignUp)
	router.POST("/activate", userCon.Activate)
	router.POST("/login", userCon.Login)
	router.POST("/logout", userCon.Logout)
	router.GET("/health_check", userCon.HealthCheck)

	loginCheckGroup := router.Group("/")
	loginCheckGroup.Use(http.CheckLogin(sessionStore))
	loginCheckGroup.GET("/users", userCon.GetUserList)
	router.POST("/tweets", tweetCon.CreateTweet)
	// loginCheckGroup.GET("/tweets", tweetCon.CreateTweet) 一覧取得
	// loginCheckGroup.GET("/tweets/:id", tweetCon.CreateTweet) 詳細取得
	// loginCheckGroup.DELETE("/tweets/:id", tweetCon.CreateTweet) ツイート削除

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
