package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

var appConfig AppConfig

type AppConfig struct {
	Host string `env:"APP_HOST" env-required:"true"`
	Port string `env:"APP_PORT" env-required:"true"`
}

func ReadAppConfig() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	configPath := root + "/.env"
	if err := cleanenv.ReadConfig(configPath, &appConfig); err != nil {
		return err
	}

	return nil
}

func GetAppConfig() *AppConfig {
	return &appConfig
}
