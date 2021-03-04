package config

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var configFilePath *string

var log *zerolog.Logger

// LoadConfig - Load the config parameters
func LoadConfig() {

	// file path is loaded
	configFilePath = flag.String("config-file-path", "config/", "config/")
	flag.Parse()

	// file name is loaded
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	// file path is loaded
	viper.AddConfigPath(*configFilePath)

	// file is read
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panic().Msgf("No config file found at %s\n", *configFilePath)
		} else {
			log.Panic().Msgf("Error reading config file: %s\n", readErr)
		}
	}
}
