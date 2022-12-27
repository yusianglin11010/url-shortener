package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string
	Host string
}

func NewConfig() AppConfig {
	viper.SetConfigFile("./config/config.yaml")
	
	if err := viper.ReadInConfig(); err != nil {
		log.Print("read config fail", err.Error())
	}
	port := viper.GetString("rest.port")
	host := viper.GetString("rest.host")
	return AppConfig{
		Port: port,
		Host: host,
	}
}
