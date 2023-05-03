package server

import (
	"github.com/gin-gonic/gin"
	"golang-server-android/controllers"
	"golang-server-android/helpers"
	"net/http"
	"strings"
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
		v1.Use(JWTMiddleware())
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
			authGroup.PUT("/info", auth.ChangeInfo)
			authGroup.PUT("/password", auth.ChangePassword)
		}
		wordGroup := v1.Group("word")
		{
			word := new(controllers.WordController)
			wordGroup.Group("/:word").
				GET("", word.SearchWord).
				GET("/example", word.GetExamples)
			wordGroup.Group("/saved").
				GET("/:id", word.GetSavedWords).
				POST("", word.SaveWord).
				DELETE("", word.UnSaveWord)
		}
		translateGroup := v1.Group("translate")
		{
			translate := new(controllers.TranslateController)
			translateGroup.GET("/", translate.GetTranslate)
			translateGroup.GET("/audio", translate.GetAudio)
			translateGroup.POST("/ocr", translate.GetOCR)
			translateGroup.POST("/speech2text", translate.Speech2Text)
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

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/v1/auth") {
			c.Next()
			return
		}
		if strings.HasPrefix(c.Request.URL.Path, "/v1/translate") {
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if len(strings.Split(authHeader, " ")) < 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		err := helpers.ValidateToken(c, tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
