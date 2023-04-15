package exercise

import (
	"routine/common"
	"routine/infrastructure"
)

var tableName = "Exercises"

type ExerciseTemplateRepo struct {
	logger common.Logger
}

type ExerciseTemplateEntity struct {
	Id          string
	Name        string
	Description string
}

func NewExersiceTemplateTable(factory *infrastructure.TableFactory[ExerciseTemplateEntity]) common.DbTable[ExerciseTemplateEntity] {
	return factory.Create(&infrastructure.TableConfig[ExerciseTemplateEntity]{
		TableName: "Exercises",
		Key:       "Id",
		EntityBlank: func() ExerciseTemplateEntity {
			return ExerciseTemplateEntity{}
		},
	})
}
