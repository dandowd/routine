package exercise

import (
	"routine/common"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ExerciseRepo struct {
	client *dynamodb.DynamoDB
}

func (*ExerciseRepo) Get(id any) ExerciseEntity {
	panic("not implemented")
}

func (*ExerciseRepo) Insert(entity ExerciseEntity) {
	panic("not implemented")
}

func NewExersiceRepo(client *dynamodb.DynamoDB) common.Repository[ExerciseEntity] {
	return &ExerciseRepo{client}
}
