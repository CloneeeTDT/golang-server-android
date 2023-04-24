package models

import (
	"golang-server-android/db"
	"time"
)

type SavedWord struct {
	WordID    uint `gorm:"primary_key;not_null, autoIncrement:false"`
	UserID    uint `gorm:"primary_key;not_null, autoIncrement:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s SavedWord) GetByUserID(userID uint) (*[]SavedWord, error) {
	database := db.GetDb()
	var result []SavedWord
	queryResult := database.Find(&result, "user_id = ?", userID).Order("word_id asc")
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return &result, nil
}

func (s SavedWord) SaveWord(payload SaveWordRequest) error {
	database := db.GetDb()
	s.WordID = payload.WordID
	s.UserID = payload.UserID
	queryResult := database.Save(&s)
	if queryResult.Error != nil {
		return queryResult.Error
	}
	return nil
}

func (s SavedWord) UnSaveWord(payload SaveWordRequest) error {
	database := db.GetDb()
	queryResult := database.Delete(&s, "word_id = ? AND user_id = ?", payload.WordID, payload.UserID)
	if queryResult.Error != nil {
		return queryResult.Error
	}
	return nil
}
