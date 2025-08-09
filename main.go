package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!!"})
}

func main(){
	router := gin.Default()
	router.GET("/",home)
  	router.GET("/health_check", func(c *gin.Context) {
    c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	router.Run()
}