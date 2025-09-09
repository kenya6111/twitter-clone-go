package api

import (
	usecase "twitter-clone-go/application/user"
	"twitter-clone-go/infrasctructure/postgres"
	presentation "twitter-clone-go/presentation/user"

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

	repo := postgres.NewUserRepository(pool)
	tx := postgres.NewTransaction(pool)
	dSer := postgres.NewUserDomainService(repo)
	ser := usecase.NewUserService(repo, tx, dSer)
	con := presentation.NewUserController(ser)

	store := memstore.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mySession", store))
	router.GET("/", con.Home)
	router.GET("/users", con.GetUserListHandler)
	router.POST("/signup", con.SignUpHandler)
	router.GET("/health_check", con.HealthCheck)
	return router
}
