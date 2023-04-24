package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": c.Param("id")})
	return
}

func (u UserController) Sync(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": c.Param("id")})
	return
}
