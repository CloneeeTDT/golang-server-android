package server

import (
	"github.com/gin-gonic/gin"
	"golang-server-android/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.GetInfo)
			userGroup.GET("/:id/sync", user.Sync)
		}
		authGroup := v1.Group("auth")
		{
			auth := new(controllers.AuthController)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/register", auth.Register)
		}
		wordGroup := v1.Group("word")
		{
			word := new(controllers.WordController)
			wordGroup.GET("/:word/example", word.GetExamples)
			wordGroup.GET("/saved/:id", word.GetSavedWords)
			wordGroup.POST("/save", word.SaveWord)
			//wordGroup.PUT("/save", word.AddNote)
			wordGroup.DELETE("/save", word.UnSaveWord)
			wordGroup.GET("/:word", word.SearchWord)
		}
		translateGroup := v1.Group("translate")
		{
			translate := new(controllers.TranslateController)
			translateGroup.GET("/", translate.GetTranslate)
			translateGroup.GET("/audio", translate.GetAudio)
			translateGroup.POST("/ocr", translate.GetOCR)
		}
	}
	return router

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
