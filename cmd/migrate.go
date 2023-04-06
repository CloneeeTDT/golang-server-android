package main

import (
	"golang-server-android/config"
	"golang-server-android/db"
	"golang-server-android/models"
)

func main() {
	config.Init()
	db.Init()
	database := db.GetDb()
	err := database.AutoMigrate(&models.User{}, &models.Word{})
	if err != nil {
		return
	}
}
