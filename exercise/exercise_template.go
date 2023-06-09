package exercise

import (
	"net/http"
	"routine/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createExerciseDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ExerciseService struct {
	logger common.Logger
	repo   common.DbTable[ExerciseTemplateEntity]
}

func (r *ExerciseService) createExerciseHandler(c *gin.Context) {
	r.logger.Info("Creating exercise")

	exercise := c.MustGet("body").(*createExerciseDto)
	item, err := r.repo.Put(ExerciseTemplateEntity{Id: uuid.New().String(), Name: exercise.Name, Description: exercise.Description})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
	return
}

func RegisterRoutes(router *gin.Engine, s *ExerciseService, v *common.Validator) {
	router.POST("/exercise", v.ValidateJSONBody(&createExerciseDto{}), s.createExerciseHandler)
}

func NewExerciseService(logger common.Logger, repo common.DbTable[ExerciseTemplateEntity]) *ExerciseService {
	return &ExerciseService{logger: logger, repo: repo}
}
