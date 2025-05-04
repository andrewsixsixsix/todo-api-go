package openapi

import (
	"os"
	"todo-api/openapi/api"

	"github.com/swaggest/openapi-go/openapi3"
)

const openapiJsonOutput = "/openapi/openapi.json"

func GenereateOpenAPI() error {
	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.1.0"}
	reflector.Spec.Info.
		WithTitle("ToDo API").
		WithVersion("0.0.1").
		WithDescription("ToDo API OpenAPI spec")

	if err := api.TodosOpenAPI(&reflector); err != nil {
		return err
	}

	schema, err := reflector.Spec.MarshalJSON()
	if err != nil {
		return err
	}

	root, err := os.Getwd()
	if err != nil {
		return err
	}

	openapiJson := root + openapiJsonOutput
	os.WriteFile(openapiJson, schema, 644)

	return nil
}
