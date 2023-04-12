package exercise

import (
	"net/http"
	"routine/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createExerciseDto struct {
	Name      string    `json:"name" binding:"required"`
	RepScheme RepScheme `json:"repScheme" binding:"required,oneof=flat rpe range"`
}

type ExerciseEntity struct {
	Id        string
	Name      string
	RepScheme RepScheme
}

type RepScheme string

const (
	Flat  RepScheme = "flat"
	Rpe   RepScheme = "rpe"
	Range RepScheme = "range"
)

type ExerciseService struct {
	logger common.Logger
	repo   common.Repository[ExerciseEntity]
}

func (r *ExerciseService) createExerciseHandler(c *gin.Context) {
	r.logger.Info("Creating exercise")

	exercise := c.MustGet("body").(*createExerciseDto)
	item, err := r.repo.Insert(ExerciseEntity{Id: uuid.New().String(), Name: exercise.Name, RepScheme: exercise.RepScheme})

	if err != nil && err.Type() == common.DatabaseError {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
	return
}

func RegisterRoutes(router *gin.Engine, s *ExerciseService, v *common.Validator) {
	router.POST("/exercise", v.ValidateJSONBody(&createExerciseDto{}), s.createExerciseHandler)
}

func NewExerciseService(logger common.Logger, repo common.CollectionRepository[ExerciseEntity]) *ExerciseService {
	return &ExerciseService{logger: logger, repo: repo}
}
