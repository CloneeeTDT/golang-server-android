package db

import (
	"database/sql"
	"fmt"
	"golang-server-android/config"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()
	var (
		host     = c.GetString("DB_HOST")
		port     = c.GetInt("DB_PORT")
		user     = c.GetString("DB_USER")
		password = c.GetString("DB_PASSWORD")
		dbname   = c.GetString("DB_NAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ = sql.Open("postgres", psqlInfo)
}

func GetDb() *sql.DB {
	return db
}
