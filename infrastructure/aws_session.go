package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession(config *aws.Config) *session.Session {
	session, err := session.NewSession(config)

	if err != nil {
		panic(err)
	}

	return session
}
