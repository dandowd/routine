package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidateJSONBody(body interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if validateErr := validate.Struct(&body); validateErr != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return

		}

		c.Set("body", body)
		c.Next()
	}
}
