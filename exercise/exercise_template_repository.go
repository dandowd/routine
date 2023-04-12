package exercise

import (
	"fmt"
	"routine/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var tableName = "Exercises"

type ExerciseTemplateRepo struct {
	client *dynamodb.DynamoDB
	logger common.Logger
}

func NewExersiceTemplateRepo(client *dynamodb.DynamoDB, logger common.Logger) common.CollectionRepository[ExerciseTemplateEntity] {
	return &ExerciseTemplateRepo{client, logger}
}

func (r *ExerciseTemplateRepo) GetPage(page int, limit int) (*[]*ExerciseTemplateEntity, *common.RepositoryError) {
	output, err := r.client.Scan(&dynamodb.ScanInput{
		TableName: &tableName,
		Limit:     aws.Int64(int64(limit)),
	})

	if err != nil {
		return nil, common.NewRepositoryError(common.DatabaseError, err.Error())
	}

	collection := make([]*ExerciseTemplateEntity, 0)

	if err := dynamodbattribute.UnmarshalListOfMaps(output.Items, collection); err != nil {
		return nil, common.NewRepositoryError(common.DatabaseError, err.Error())
	}

	return &collection, nil
}

func (r *ExerciseTemplateRepo) Get(id string) (*ExerciseTemplateEntity, *common.RepositoryError) {
	partitionKey, err := dynamodbattribute.ConvertTo(fmt.Sprintf("exercise#%s", id))

	if err != nil {
		r.logger.Error("Could not create exercise partitionKey")
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not create exercise partitionKey"))
	}

	output, err := r.client.GetItem(&dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       map[string]*dynamodb.AttributeValue{"partition": partitionKey},
	})

	if output.Item == nil {
		return nil, common.NewRepositoryError(common.NotFound, fmt.Sprintf("Item not found"))
	}

	entity := ExerciseTemplateEntity{}

	if err := dynamodbattribute.UnmarshalMap(output.Item, &entity); err != nil {
		r.logger.Error(fmt.Sprintf("Fatal error while unmarshalling exercise from dynamodb: %s", err))
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not insert Exercise Entity"))
	}

	return &entity, nil
}

func (r *ExerciseTemplateRepo) Insert(entity ExerciseTemplateEntity) (*ExerciseTemplateEntity, *common.RepositoryError) {
	av, err := dynamodbattribute.MarshalMap(entity)

	if err != nil {
		r.logger.Error(fmt.Sprintf("Error marshalling exercise: %s", err))
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not insert Exercise Entity"))
	}

	partitionKey, err := dynamodbattribute.ConvertTo(fmt.Sprintf("exercise#%s", entity.Id))

	if err != nil {
		r.logger.Error(fmt.Sprintf("Error while creating partitionKey for exercise: %s", err))
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not insert Exercise Entity"))
	}

	av["partition"] = partitionKey

	output, err := r.client.PutItem(&dynamodb.PutItemInput{
		TableName:    &tableName,
		Item:         av,
		ReturnValues: aws.String("ALL_NEW"),
	})

	if err != nil {
		r.logger.Error(fmt.Sprintf("Error while inserting exercise: %s", err))
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not insert Exercise Entity"))
	}

	updateEntity := ExerciseTemplateEntity{}

	if err := dynamodbattribute.UnmarshalMap(output.Attributes, &updateEntity); err != nil {
		// if the entity is inserted but can't be unmarshalled we will get an error everytime this record is fetched
		r.logger.Fatal(fmt.Sprintf("Fatal error while unmarshalling exercise from dynamodb: %s", err))
		return nil, common.NewRepositoryError(common.DatabaseError, fmt.Sprintf("Could not insert Exercise Entity"))
	}

	return &updateEntity, nil
}
