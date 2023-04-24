package models

import (
	"golang-server-android/db"
	"time"
)

type NotedWord struct {
	WordID    uint   `gorm:"primary_key;not_null, autoIncrement:false"`
	UserID    uint   `gorm:"primary_key;not_null, autoIncrement:false"`
	Note      string `gorm:"not_null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n NotedWord) GetByUserID(userID uint) (*[]NotedWord, error) {
	database := db.GetDb()
	var result []NotedWord
	queryResult := database.Find(&result, "user_id = ?", userID).Order("word_id asc")
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return &result, nil
}
