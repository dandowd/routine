package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAWSConfig),
	fx.Provide(NewSession),
	fx.Provide(NewDynamoDbClient),
)
