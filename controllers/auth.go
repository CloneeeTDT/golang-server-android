package controllers

import (
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
	token, err := helpers.GenerateJWT(*user)
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

func (u AuthController) ChangeInfo(c *gin.Context) {
	body := models.UpdateUserRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	id := c.GetUint("id")
	user, err := User.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	err = user.UpdateInfo(body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update success"})
	return
}

func (u AuthController) ChangePassword(c *gin.Context) {
	body := models.UpdatePasswordRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	id := c.GetUint("id")
	user, err := User.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		c.Abort()
		return
	}
	err = user.UpdatePassword(body.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update success"})
	return
}
