package validation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	tag, param := fe.Tag(), fe.Param()
	switch tag {
	case "lte":
		return "Should be less than " + param
	case "gte":
		return "Should be greater than " + param
	case "required":
		return "Field is required"
	case "max":
		return fmt.Sprintf("Cannot be longer than %s", param)
	case "min":
		return fmt.Sprintf("Must be longer than %s", param)
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf("Must be %s characters long", param)
	}

	return "Unknown error"
}

func ErrorValidation(err error, c *gin.Context) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, out)
	}
}
