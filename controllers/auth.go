package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-server-android/helpers"
	"golang-server-android/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct{}

var User = new(models.User)

func (u AuthController) Login(c *gin.Context) {
	body := models.LoginRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := User.GetByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}
	token, err := helpers.GenerateJWT(user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}

func (u AuthController) Register(c *gin.Context) {
	body := models.RegisterRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	checkUser, _ := User.GetByEmail(body.Email)
	fmt.Println(*checkUser)
	if checkUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		c.Abort()
		return
	}
	err := User.Register(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	return
}
