package main

import (
	"fmt"
	"net/http"
	"os"

	"todo-api/config"
	"todo-api/internal/router"
	"todo-api/openapi"
)

func main() {
	if err := config.ReadAppConfig(); err != nil {
		fmt.Println("failed to read app config")
		os.Exit(1)
	}

	if err := openapi.GenereateOpenAPI(); err != nil {
		fmt.Println("failed to generate openapi spec")
		os.Exit(1)
	}

	r := router.Router()

	address := fmt.Sprintf("%s:%s", config.GetAppConfig().Host, config.GetAppConfig().Port)
	http.ListenAndServe(address, r)
}
