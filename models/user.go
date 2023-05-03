package models

import (
	"golang-server-android/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key;auto_increment;not_null"`
	Name      string
	Email     string
	Password  string
	Birthday  datatypes.Date
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) GetByEmail(email string) (*User, error) {
	database := db.GetDb()

	queryResult := database.Find(&u, "email = ?", email)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return &u, nil
}

func (u User) GetByID(id uint) (*User, error) {
	database := db.GetDb()

	queryResult := database.Find(&u, "id = ?", id)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return &u, nil
}

func (u User) Register(payload RegisterRequest) error {
	database := db.GetDb()

	u.Name = payload.Name
	u.Email = payload.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	birthday, err := time.Parse("2006-01-02", payload.Birthday)
	if err != nil {
		return err
	}
	u.Birthday = datatypes.Date(birthday)
	queryResult := database.Create(&u)
	if queryResult.Error != nil {
		return queryResult.Error
	}
	return nil
}

func (u User) UpdateInfo(name string) error {
	database := db.GetDb()
	u.Name = name
	queryResult := database.Save(&u)
	if queryResult.Error != nil {
		return queryResult.Error
	}
	return nil
}

func (u User) UpdatePassword(password string) error {
	database := db.GetDb()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	queryResult := database.Save(&u)
	if queryResult.Error != nil {
		return queryResult.Error
	}
	return nil
}
