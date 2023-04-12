package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	logger Logger
}

type ValidationError struct {
	Key   string
	Error string
}

func NewValidator(logger Logger) *Validator {
	return &Validator{logger}
}

func (v *Validator) ValidateJSONBody(body interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		v.logger.Info("Request failed due to validation error")

		if validationErr := c.ShouldBindBodyWith(&body, binding.JSON); validationErr != nil {

			errorList := validationErr.(validator.ValidationErrors)
			errors := make([]*ValidationError, 0)

			for _, error := range errorList {
				currentError := &ValidationError{Key: error.Field(), Error: error.ActualTag()}
				errors = append(errors, currentError)
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}

		c.Set("body", body)
		c.Next()
	}
}
