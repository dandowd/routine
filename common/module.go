package common

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewValidator),
)
