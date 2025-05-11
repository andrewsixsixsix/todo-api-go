package service

import (
	"reflect"
	"strings"
	"todo-api/internal/model"

	"github.com/go-playground/validator/v10"
)

var rbvService *RequestBodyValidationService

type RequestBodyValidationService struct {
	Validate *validator.Validate
}

func InitRbvService(validate *validator.Validate) {
	validate.RegisterTagNameFunc(jsonTagNameExtractor)
	rbvService = &RequestBodyValidationService{Validate: validate}
}

func GetRbvService() *RequestBodyValidationService {
	return rbvService
}

// TODO: implement Validate method

func (rbvs *RequestBodyValidationService) CollectValidationFails(ve validator.ValidationErrors) []model.ValidationFail {
	// collect all messages for a particular field
	// field: []message
	fm := make(map[string][]string)

	for _, e := range ve {
		field := e.Field()
		msg := strings.Split(e.Error(), "Error:")[1]

		if _, ok := fm[field]; ok {
			fm[field] = append(fm[field], msg)
		} else {
			fm[field] = []string{msg}
		}
	}

	// collect all fields with their messages into slice
	var res []model.ValidationFail

	for key, val := range fm {
		vf := model.ValidationFail{
			Field:    key,
			Messages: val,
		}
		res = append(res, vf)
	}

	return res
}

func jsonTagNameExtractor(f reflect.StructField) string {
	name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}
