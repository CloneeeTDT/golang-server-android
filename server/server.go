package server

import (
	"fmt"
	"golang-server-android/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	err := r.Run(fmt.Sprintf(":%s", c.GetString("server.port"))) // Run in 8080
	if err != nil {
		return
	}
}
