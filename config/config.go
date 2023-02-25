package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigFile("config/config.yaml")
	err = config.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
