package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": c.Param("id")})
	return
}
