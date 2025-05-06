package handler

import (
	"log/slog"
	"net/http"
	"os"
	"todo-api/internal/logger"
)

const (
	swaggerHtml = "/openapi/swagger.html"
	openapiJson = "/openapi/openapi.json"
)

func UI(w http.ResponseWriter, r *http.Request) {
	projectRoot, err := os.Getwd()
	if err != nil {
		logger.Logger.Error("failed to get project's root directory", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	swaggerHtmlPath := projectRoot + swaggerHtml
	swaggerHtml, err := os.ReadFile(swaggerHtmlPath)
	if err != nil {
		logger.Logger.Error("failed to read swagger.html", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(swaggerHtml)
}

func OpenAPI(w http.ResponseWriter, r *http.Request) {
	projectRoot, err := os.Getwd()
	if err != nil {
		logger.Logger.Error("failed to get project's root directory", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	openapiJsonPath := projectRoot + openapiJson
	openapiJson, err := os.ReadFile(openapiJsonPath)
	if err != nil {
		logger.Logger.Error("failed to read openapi.json", slog.String("err", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(openapiJson)
}
