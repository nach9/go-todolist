package validation

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

func getErrorMsg(fe validator.FieldError, jsonFieldName string) string {
	tag, param := fe.Tag(), fe.Param()
	switch tag {
	case "lte":
		return "Should be less than " + param
	case "gte":
		return "Should be greater than " + param
	case "required":
		return fmt.Sprintf("%s cannot be null", jsonFieldName)
	case "max":
		return fmt.Sprintf("Cannot be longer than %s", param)
	case "min":
		return fmt.Sprintf("Must be longer than %s", param)
	case "email":
		return "Invalid email format"
	case "boolean":
		return "Invalid boolean value"
	case "len":
		return fmt.Sprintf("Must be %s characters long", param)
	}

	return "Unknown error"
}

func ErrorValidation(err error, c *gin.Context, sType any) {
	var ve validator.ValidationErrors
	structType := reflect.TypeOf(sType)
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			fieldName := fe.StructField()
			jsonFieldName := getJSONFieldName(structType, fieldName)
			out[i] = ErrorMsg{
				"Bad Request",
				getErrorMsg(fe, jsonFieldName),
				fe.Field(),
			}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, out[0])
	}
}

func getJSONFieldName(structType reflect.Type, fieldName string) string {
	field, _ := structType.FieldByName(fieldName)
	jsonTag := field.Tag.Get("json")
	jsonFieldName := strings.Split(jsonTag, ",")[0]
	return jsonFieldName
}
