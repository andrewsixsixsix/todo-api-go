package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type StorageConfig struct {
	URL string `env:"DB_URL" env-required:"true"`
}

var storageConfig StorageConfig

func ReadStorageConfig() error {
	root, err := os.Getwd()
	if err != nil {
		return nil
	}

	configPath := root + "/.env"
	if err := cleanenv.ReadConfig(configPath, &storageConfig); err != nil {
		return err
	}

	return nil
}

func GetStorageConfig() *StorageConfig {
	return &storageConfig
}
