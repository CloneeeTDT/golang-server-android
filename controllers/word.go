package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-server-android/api"
	"golang-server-android/models"
	"net/http"
	"strconv"
)

type WordController struct{}

var SaveWord = new(models.SavedWord)

func (u WordController) GetExamples(c *gin.Context) {
	result := api.GetExamples(c.Param("word"))
	c.JSON(http.StatusOK, gin.H{"examples": result})
	return
}

func (u WordController) SearchWord(c *gin.Context) {
	word := api.SearchWord(c.Param("word"))
	c.JSON(http.StatusOK, word)
	return
}

func (u WordController) SaveWord(c *gin.Context) {
	var payload models.SaveWordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := SaveWord.SaveWord(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Word saved successfully"})
	return
}

func (u WordController) UnSaveWord(c *gin.Context) {
	var payload models.SaveWordRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := SaveWord.UnSaveWord(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Word unsaved successfully"})
	return
}

func (u WordController) GetSavedWords(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	words, _ := SaveWord.GetByUserID(uint(id))
	c.JSON(http.StatusOK, *words)
	return
}
