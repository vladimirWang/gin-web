package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
		// Host     string
		// Port     string
		// User     string
		// Password string
		// Name     string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v\n", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v\n", err)
	}

	initDB()
}
