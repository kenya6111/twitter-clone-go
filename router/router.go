package router

import (
	"twitter-clone-go/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Home)
	router.GET("/users", controllers.GetUserListHandler)
	router.POST("/signup", controllers.SignUpHandler)
	router.POST("/activate", controllers.ActivateHandler)
	router.GET("/health_check", controllers.HealthCheck)
	return router
}
