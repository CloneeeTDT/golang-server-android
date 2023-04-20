package server

import (
	"github.com/gin-gonic/gin"
	"golang-server-android/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.GetInfo)
		}
		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/register", auth.Register)
		}
	}
	return router

}
