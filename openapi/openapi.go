package openapi

import (
	"fmt"
	"os"
	"todo-api/openapi/api"

	"github.com/swaggest/openapi-go/openapi3"
)

const openapiJsonOutput = "/openapi/openapi.json"

func GenereateOpenAPI() {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}
	reflector.Spec.Info.
		WithTitle("ToDo API").
		WithVersion("0.0.1").
		WithDescription("ToDo API OpenAPI spec")

	if err := api.TodosOpenAPI(&reflector); err != nil {
		fmt.Println("failed to create openapi spec for todos")
		os.Exit(1)
	}

	schema, err := reflector.Spec.MarshalJSON()
	if err != nil {
		fmt.Println("failed to marshal schema json")
		os.Exit(1)
	}

	root, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get project's root directory")
		os.Exit(1)
	}

	openapiJson := root + openapiJsonOutput
	os.WriteFile(openapiJson, schema, 644)
}
