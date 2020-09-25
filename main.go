package main

import (
	"email-service/controller"
	"email-service/middleware"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	router.Use(middleware.Cors())
	router.Any("/ping", controller.Ping)
	router.GET("/health.json", controller.Health)
	router.POST("/email", controller.EmailHandler)
	router.GET("/user/:instance/:id", controller.UserHandler)
	router.Run("127.0.0.1:9000")
}
