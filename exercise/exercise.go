package exercise

import (
	"fmt"
	"routine/common"

	"github.com/gin-gonic/gin"
)

type ExerciseDto struct {
	Name      string `json:"name" binding:"required"`
	RepScheme int    `json:"reps" binding:"required"`
}

type ExerciseService struct {
	logger common.Logger
}

func (r *ExerciseService) createExerciseHandler(c *gin.Context) {
	exercise := c.MustGet("body").(*ExerciseDto)
	r.logger.Info(fmt.Sprintf("Exercise: %v", exercise))
}

func RegisterRoutes(router *gin.Engine, s *ExerciseService) {
	router.POST("/exercise", common.ValidateJSONBody(&ExerciseDto{}), s.createExerciseHandler)
}

func NewExerciseService(logger common.Logger) *ExerciseService {
	return &ExerciseService{logger: logger}
}
