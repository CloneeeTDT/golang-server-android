package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var config *viper.Viper

func Init() {
	var err error
	config = viper.New()
	env, _ := os.LookupEnv("RUN_ENV")
	if env == "PROD" {
		config.SetConfigFile("config/config.prod.yaml")
	} else {
		config.SetConfigFile("config/config.yaml")
	}
	err = config.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
