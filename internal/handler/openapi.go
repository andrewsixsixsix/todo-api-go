package handler

import (
	"fmt"
	"net/http"
	"os"
)

const (
	swaggerHtml = "/openapi/swagger.html"
	openapiJson = "/openapi/openapi.json"
)

func UI(w http.ResponseWriter, r *http.Request) {
	projectRoot, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to get project's root directory")
		return
	}

	swaggerHtmlPath := projectRoot + swaggerHtml
	swaggerHtml, err := os.ReadFile(swaggerHtmlPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to read swagger.html")
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(swaggerHtml)
}

func OpenAPI(w http.ResponseWriter, r *http.Request) {
	projectRoot, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to get project's root directory")
		return
	}

	openapiJsonPath := projectRoot + openapiJson
	openapiJson, err := os.ReadFile(openapiJsonPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("failed to read openapi.json")
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(openapiJson)
}
