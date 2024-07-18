package validator

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"reflect"
	"strings"
)

func getStructValidationMessage(field string, t string) string {
	switch t {
	case "required":
		return fmt.Sprintf("the %s attribute field is required.", field)
	case "len":
		return fmt.Sprintf("the %s attribute length is required.", field)
	default:
		return fmt.Sprintf("%s", t)
	}
}

func getTypeValidationMessage(field string, t reflect.Type) string {
	switch t.Name() {
	case "string":
		return fmt.Sprintf("the %s must be a string.", field)
	case "int", "float32", "float64":
		return fmt.Sprintf("the %s must be a number.", field)
	default:
		return fmt.Sprintf("the %s has the wrong type.", field)
	}
}

func validateJsonUnmarshal(body io.Reader, v interface{}) map[string]string {
	if err := json.NewDecoder(body).Decode(&v); err != nil {
		errors := make(map[string]string)
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			errors[e.Field] = getTypeValidationMessage(e.Field, e.Type)
		case *json.InvalidUnmarshalError:
			errors[e.Type.String()] = fmt.Sprintf("%v", e.Type)
		default:
			log.Println(err)
		}
		return errors
	}
	return nil
}

func validateRequest(v interface{}) map[string]string {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strings.TrimPrefix(err.StructNamespace(), "QuoteRequest.")
			fmt.Printf("%v", err)
			errors[fieldName] = getStructValidationMessage(fieldName, err.Tag())
		}
		return errors
	}
	return nil
}

func Validate(body io.Reader, v interface{}) map[string]string {
	if errors := validateJsonUnmarshal(body, v); errors != nil {
		return errors
	}
	if errors := validateRequest(v); errors != nil {
		return errors
	}
	return nil
}
