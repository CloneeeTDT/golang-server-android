package models

import (
	"fmt"
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

	fmt.Println(database)
	queryResult := database.Find(&u, "email = ?", email)
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
