package builder

import (
	"routine/common"
	"routine/exercise"
	"routine/infrastructure"

	"go.uber.org/fx"
)

func defaultOptions() fx.Option {
	return fx.Options(
		fx.Provide(NewHTTPServer),
		common.Module,
		infrastructure.Module,
		exercise.Module,
	)
}

func defaultInvokers() fx.Option {
	return fx.Options(
		fx.Invoke(exercise.RegisterRoutes),
	)
}

func DefaultAppBuilder() *fx.App {
	app := fx.New(
		defaultOptions(),
	)

	return app
}

func AppBuilderWithOptions(options fx.Option) *fx.App {
	app := fx.New(
		defaultOptions(),
		options,
	)

	return app
}
