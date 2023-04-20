package db

import (
	"fmt"
	"golang-server-android/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	c := config.GetConfig()
	var (
		host     = c.GetString("db.host")
		port     = c.GetInt("db.port")
		user     = c.GetString("db.user")
		password = c.GetString("db.password")
		dbname   = c.GetString("db.name")
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}
}

func GetDb() *gorm.DB {
	return db
}
