package models

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func GetValidationMessage(fe validator.FieldError) string {
	switch fe.Field() {
	case "Name":
		switch fe.Tag() {
		case "required":
			return "Name is required."
		case "min":
			return fmt.Sprintf("Name must be at least %s characters long.", fe.Param())
		case "max":
			return fmt.Sprintf("Name must be at most %s characters long.", fe.Param())
		case "alphanum":
			return "Name must contain only letters and numbers."
		}
	case "EvolvesAt":
		switch fe.Tag() {
		case "required":
			return "Evolution Level is required."
		case "min":
			return fmt.Sprintf("Evolution Level must be at least %s.", fe.Param())
		case "max":
			return fmt.Sprintf("Evolution Level must be at most %s.", fe.Param())
		}
	}
	switch fe.Tag() {
	case "required_at_least_one":
		return "At least one field must be provided."
	}
	return fe.Error()
}

func requiredAtLeastOne(sl validator.StructLevel) {
	dto := sl.Current().Interface().(PokemonUpdateDTO)
	if dto.Name == "" && dto.EvolvesAt == 0 {
		sl.ReportError(
			reflect.ValueOf(""),
			"AtLeastOneField",
			"",
			"required_at_least_one",
			"",
		)
	}
}

func init() {
	validate.RegisterStructValidation(requiredAtLeastOne, PokemonUpdateDTO{})
}
