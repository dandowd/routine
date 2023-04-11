package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	logger Logger
}

func NewValidator(logger Logger) *Validator {
	return &Validator{logger}
}

func (v *Validator) ValidateJSONBody(body interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if validateErr := validator.New().Struct(&body); validateErr != nil {
			v.logger.Error(fmt.Sprintf("Request body failed to validate due to: %e", validateErr))
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return

		}

		c.Set("body", body)
		c.Next()
	}
}
