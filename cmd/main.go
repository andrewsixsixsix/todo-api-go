package main

import (
	"net/http"

	"todo-api/internal/router"
	"todo-api/openapi"
)

func main() {
	openapi.GenereateOpenAPI()

	r := router.Router()

	http.ListenAndServe(":6969", r)
}
