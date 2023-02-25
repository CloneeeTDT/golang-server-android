package main

import (
	"golang-server-android/config"
	"golang-server-android/db"
	"golang-server-android/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
