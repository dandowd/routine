package integration_test

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type DbContainer struct {
	containerId string
	client      *client.Client
}

func NewDbContainer() *DbContainer {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.41", nil, nil)
	if err != nil {
		panic(err)
	}

	containerConfig := &container.Config{
		Image: "amazon/dynamodb-local",
		Cmd:   []string{"-jar", "DynamoDBLocal.jar", "-inMemory"},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"8000/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8000",
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, nil, "")

	cleanup := func() {
		if err := cli.ContainerRemove(context.Background(), resp.ID, types.ContainerRemoveOptions{}); err != nil {
			panic(err)
		}
	}

	if err != nil {
		cleanup()
		panic(err)
	}

	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	return &DbContainer{containerId: resp.ID, client: cli}
}

func (d *DbContainer) Cleanup() {
	d.client.ContainerStop(context.Background(), d.containerId, container.StopOptions{})
	d.client.ContainerRemove(context.Background(), d.containerId, types.ContainerRemoveOptions{RemoveVolumes: true})
}

func CreateTables(client *dynamodb.DynamoDB) {
	client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String("Exercises"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("partition"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("partition"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
}
