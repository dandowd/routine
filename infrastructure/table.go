package infrastructure

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type TableFactory[T interface{}] struct {
	client *dynamodb.DynamoDB
}

type Table[T interface{}] struct {
	client *dynamodb.DynamoDB
	config *TableConfig[T]
}

type TableConfig[T interface{}] struct {
	TableName   string
	Key         string
	EntityBlank func() T
}

func NewTableFactory[T interface{}](client *dynamodb.DynamoDB) *TableFactory[T] {
	return &TableFactory[T]{client: client}
}

func (f *TableFactory[T]) Create(config *TableConfig[T]) *Table[T] {
	return &Table[T]{client: f.client, config: config}
}

func (t *Table[T]) Get(id string) (*T, error) {
	partitionKey, err := dynamodbattribute.ConvertTo(fmt.Sprintf("%s#%s", t.config.Key, id))

	if err != nil {
		return nil, err
	}

	output, err := t.client.GetItem(&dynamodb.GetItemInput{
		TableName: &t.config.TableName,
		Key:       map[string]*dynamodb.AttributeValue{t.config.Key: partitionKey},
	})

	entity := t.config.EntityBlank()

	if err := dynamodbattribute.UnmarshalMap(output.Item, &entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (t *Table[T]) Put(entity T) (*T, error) {
	attributeMap, err := dynamodbattribute.ConvertToMap(entity)

	if err != nil {
		return nil, err
	}

	_, err = t.client.PutItem(&dynamodb.PutItemInput{
		TableName: &t.config.TableName,
		Item:      attributeMap,
	})

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *Table[T]) GetPage(limit int, startId *string) (*[]*T, error) {
	var key map[string]*dynamodb.AttributeValue

	if startId != nil {
		key = map[string]*dynamodb.AttributeValue{r.config.Key: {S: aws.String(fmt.Sprintf("%s#%s", r.config.Key, *startId))}}

	}

	output, err := r.client.Scan(&dynamodb.ScanInput{
		TableName:         &r.config.TableName,
		Limit:             aws.Int64(int64(limit)),
		ExclusiveStartKey: key,
	})

	if err != nil {
		return nil, err
	}

	collection := make([]*T, 0)

	if err := dynamodbattribute.UnmarshalListOfMaps(output.Items, collection); err != nil {
		return nil, err
	}

	return &collection, nil
}
