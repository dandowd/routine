package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewSession),
	fx.Provide(NewAWSConfig),
	fx.Provide(NewDynamoDbClient),
)
