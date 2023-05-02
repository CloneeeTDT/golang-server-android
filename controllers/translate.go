package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ocr "github.com/ranghetto/go_ocr_space"
	"golang-server-android/api"
	"golang-server-android/config"
	"golang-server-android/models"
	"net/http"
	"strings"
)

type TranslateController struct{}

func (u TranslateController) GetTranslate(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	text := c.Query("text")
	if len(text) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Text is empty"})
		return
	}
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
	result, err := ocrConfig.ParseFromBase64(fmt.Sprintf("data:image/png;base64,%s", body.Image))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result.JustText()})
	return
}

func (u TranslateController) Speech2Text(c *gin.Context) {
	body := models.Speech2TextRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body.Audio = strings.ReplaceAll(body.Audio, "\n", "")
	googleResponse, err := api.GetTextFromAudio(body.Audio, body.From)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	text := googleResponse.Results[0].Alternatives[0].Transcript
	translateResponse := api.GetTranslate(body.From, body.To, text)
	if translateResponse == nil || len(translateResponse.Sentences) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Translate error"})
		return
	}
	translateText := translateResponse.Sentences[0].Trans
	audio := api.GetAudio(body.To, translateText)
	result := models.Speech2TextResponse{
		Content:   googleResponse.Results[0].Alternatives[0].Transcript,
		Translate: translateResponse.Sentences[0].Trans,
		From:      body.From,
		To:        body.To,
		Audio:     audio,
	}
	c.JSON(http.StatusOK, result)
}
