package api

import (
	application "twitter-clone-go/application/user"
	"twitter-clone-go/infrasctructure/email/mailcatcher"
	"twitter-clone-go/infrasctructure/storage/postgres"
	"twitter-clone-go/interface/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(pool *pgxpool.Pool) {
	router := NewRouter(pool)
	router.Run()
}

func NewRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	var emailService = mailcatcher.NewMainCatcherEmailService("test")
	repo := postgres.NewUserRepository(pool)
	tx := postgres.NewTransaction(pool)
	dSer := application.NewUserDomainService(repo)
	ser := application.NewUserUsecase(repo, tx, dSer, emailService)
	con := http.NewUserHandler(ser)

	store := memstore.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mySession", store))
	router.GET("/", con.Home)
	router.GET("/users", con.GetUserListHandler)
	router.POST("/signup", con.SignUpHandler)
	router.GET("/health_check", con.HealthCheck)
	return router
}
