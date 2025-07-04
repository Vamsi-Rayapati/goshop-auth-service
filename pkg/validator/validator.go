package validator

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors  []FieldError `json:"errors,omitempty"`
	Message string       `json:"message"`
}

func GetMessageByTag(tag string) string {
	switch tag {
	case "required":
		return "required field"
	case "email":
		return "invalid email format"
	default:
		return tag
	}
}
func ValidateBody(c *gin.Context, obj any) *ValidationError {
	if err := c.ShouldBindJSON(obj); err != nil {
		return &ValidationError{
			Message: "Invalid request payload",
		}
	}

	var validate = validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	if err := validate.Struct(obj); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			validationErrors := []FieldError{}

			for _, e := range ve {
				validationErrors = append(validationErrors, FieldError{Field: e.Field(), Message: GetMessageByTag(e.Tag())})
			}

			return &ValidationError{
				Errors:  validationErrors,
				Message: "Invalid request payload",
			}
		}
		return &ValidationError{
			Message: "Invalid request payload",
		}
	}

	return nil

}
