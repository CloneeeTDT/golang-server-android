package server

import "golang-server-android/config"

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	err := r.Run(c.GetString("server.port")) // Run in 8080
	if err != nil {
		return
	}
}
