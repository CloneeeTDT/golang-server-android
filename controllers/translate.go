package controllers

import (
	"github.com/gin-gonic/gin"
	ocr "github.com/ranghetto/go_ocr_space"
	"golang-server-android/api"
	"golang-server-android/config"
	"golang-server-android/models"
	"net/http"
)

type TranslateController struct{}

func (u TranslateController) GetTranslate(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	text := c.Query("text")
	result := api.GetTranslate(from, to, text)
	c.JSON(http.StatusOK, gin.H{"from": from, "to": to, "origin": text, "result": result.Sentences[0].Trans})
	return
}

func (u TranslateController) GetAudio(c *gin.Context) {
	tl := c.Query("tl")
	text := c.Query("text")
	data := api.GetAudio(tl, text)
	if len(data) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Empty data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tl": tl, "text": text, "data": data})
	return
}

func (u TranslateController) GetOCR(c *gin.Context) {
	mainConfig := config.GetConfig()
	apiKey := mainConfig.GetString("ocr.key")
	if len(apiKey) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API KEY is not set"})
		return
	}
	ocrConfig := ocr.InitConfig(apiKey, "eng")
	body := models.OCRRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := ocrConfig.ParseFromBase64(body.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result.JustText()})
	return
}
