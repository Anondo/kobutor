package config

import (
	"log"

	"github.com/spf13/viper"
)

// Init initializes the config
func Init() {
	viper.SetConfigName("kobutor_config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read config file")
	}

	loadAuth()
}
