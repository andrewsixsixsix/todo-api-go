package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"todo-api/config"
	"todo-api/internal/logger"
	"todo-api/internal/router"
	"todo-api/internal/storage"
	"todo-api/openapi"
)

func main() {
	logger.Logger = logger.Init()
	logger.Logger.Info("Logger initialized. Level: INFO")

	if err := config.ReadAppConfig(); err != nil {
		logger.Logger.Error("failed to read app config", slog.String("err", err.Error()))
		os.Exit(1)
	}

	logger.Logger.Info("App config read successfully")

	if err := config.ReadStorageConfig(); err != nil {
		logger.Logger.Error("failed to read storage config", slog.String("err", err.Error()))
		os.Exit(1)
	}

	logger.Logger.Info("Storage config read successfully")

	err := storage.Init(config.GetStorageConfig().URL)
	if err != nil {
		logger.Logger.Error("failed to initialize storage", slog.String("err", err.Error()))
		os.Exit(1)
	}

	if err := storage.Storage().Ping(); err != nil {
		logger.Logger.Error("storage ping resulted with error", slog.String("err", err.Error()))
		os.Exit(1)
	}

	logger.Logger.Info("Storage initialized successfully")

	if err := openapi.GenereateOpenAPI(); err != nil {
		logger.Logger.Error("failed to generate openapi spec", slog.String("err", err.Error()))
		os.Exit(1)
	}

	logger.Logger.Info("OpenAPI spec generated successfully")

	r := router.Router()

	logger.Logger.Info("Router initialized")

	address := fmt.Sprintf("%s:%s", config.GetAppConfig().Host, config.GetAppConfig().Port)

	logger.Logger.Info("Starting server", slog.String("address", address))

	http.ListenAndServe(address, r)
}
