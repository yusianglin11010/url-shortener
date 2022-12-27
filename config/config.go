package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Rest  *RestConfig  `mapstructure:"rest"`
	Redis *RedisConfig `mapstructure:"redis"`
}

type RestConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type RedisConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

func NewConfig() AppConfig {
	viper.SetConfigFile("./config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Print("read config fail", err.Error())
	}

	var AppConfig AppConfig
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("fail to load config")
	}
	return AppConfig
}
