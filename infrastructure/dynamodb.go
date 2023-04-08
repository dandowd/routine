package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDbClient(session *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(session)
}
