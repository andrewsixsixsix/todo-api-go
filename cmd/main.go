package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"todo-api/config"
	"todo-api/internal/logger"
	"todo-api/internal/router"
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
