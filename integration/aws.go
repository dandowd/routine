package integration_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func NewAWSIntegrationTestConfig() *aws.Config {
	return &aws.Config{
		Credentials: credentials.NewStaticCredentials("test", "test", "test"),
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:8000"),
	}
}
