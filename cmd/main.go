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

	if err := openapi.GenereateOpenAPI(); err != nil {
		logger.Logger.Error("failed to generate openapi spec", slog.String("err", err.Error()))
		os.Exit(1)
	}

	r := router.Router()

	logger.Logger.Info("Router initialized")

	address := fmt.Sprintf("%s:%s", config.GetAppConfig().Host, config.GetAppConfig().Port)
	http.ListenAndServe(address, r)
}
