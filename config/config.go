// Package config manages application configuration
package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config stores all configuration variables
// The values are ready by viper from a config file or environment variable
type Config struct {
	AppHost string `mapstructure:"APP_HOST"`
	AppPort int    `mapstructure:"APP_PORT"`
}

// LoadConfig loads env variables to config object
func LoadConfig() (config Config, err error) {
	// config missing jugaad
	_, err = os.Open("config/config.yml")
	if err != nil {
		log.Print("copying config file template...")
		os.Link("config/sample.config.yml", "config/config.yml")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
