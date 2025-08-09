package router

import (
	"twitter-clone-go/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("userInfo", store))
	router.GET("/", controllers.Home)
	router.GET("/users", controllers.GetUserListHandler)
	router.POST("/signup", controllers.SignUpHandler)
	router.GET("/health_check", controllers.HealthCheck)
	return router
}
