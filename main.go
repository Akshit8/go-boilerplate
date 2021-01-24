package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// config struct all configuration variables
// The values are ready by viper from a config file or environment variable
type config struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppHost string `mapstructure:"APP_HOST"`
	AppPort int    `mapstructure:"APP_PORT"`
}

// Config holds all configuration variable
var Config config

func init() {
	// set application logger configuration
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	// config missing jugaad
	_, err := os.Open("config/config.yml")
	if err != nil {
		log.Print("copying config file template...")
		os.Link("config/sample.config.yml", "config/config.yml")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		log.Error("cannot read config: ", err)
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Error("cannot unmarshal config: ", err)
	}
}

func main() {
	log.Info(Config)
}
