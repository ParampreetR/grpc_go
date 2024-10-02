package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPass string `mapstructure:"POSTGRES_PASS"`
	PostgresDB   string `mapstructure:"POSTGRES_DB"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort int    `mapstructure:"POSTGRES_PORT"`
	ClientPort   int    `mapstructure:"CLIENT_PORT"`
	ServerPort   int    `mapstructure:"SERVER_PORT"`
}

var AppConfig Config

func LoadEnvVariables() {
	viper.SetConfigName(".env")

	// Set the config type to "env"
	viper.SetConfigType("env")

	// Add the path to look for the config file
	viper.AddConfigPath("config")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error to decode ENV data into struct: %s", err)
	}
}

func GetEnvConfig() Config {
	return AppConfig
}
