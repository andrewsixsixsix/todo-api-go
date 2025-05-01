package main

import (
	"net/http"

	"todo-api/internal/router"
)

func main() {
	r := router.Router()

	http.ListenAndServe(":6969", r)
}
