package router

import (
	"twitter-clone-go/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Run(con *controllers.MyAppController) {
	router := SetupRouter(con)
	router.Run()
}

func SetupRouter(con *controllers.MyAppController) *gin.Engine {
	router := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("userInfo", store))
	router.GET("/", con.Home)
	router.GET("/users", con.GetUserListHandler)
	router.POST("/signup", con.SignUpHandler)
	router.GET("/health_check", con.HealthCheck)
	return router
}
