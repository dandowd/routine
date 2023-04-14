package infrastructure

import "github.com/aws/aws-sdk-go/aws"

func NewAWSConfig() *aws.Config {
	return &aws.Config{
		Region: aws.String("us-east-1"),
	}
}
