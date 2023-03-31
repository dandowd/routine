package builder

import (
	"routine/common"
	"routine/exercise"

	"go.uber.org/fx"
)

func defaultProviders() fx.Option {
	return fx.Options(
		fx.Provide(NewHTTPServer),
		fx.Provide(common.NewLogger),
		fx.Provide(exercise.NewExerciseService),
	)
}

func defaultInvokers() fx.Option {
	return fx.Options(
		fx.Invoke(exercise.RegisterRoutes),
	)
}

func DefaultAppBuilder() *fx.App {
	app := fx.New(
		defaultProviders(),
		defaultInvokers(),
	)

	return app
}

func AppBuilder(addOptions fx.Option) *fx.App {
	app := fx.New(
		defaultProviders(),
		addOptions,
		defaultInvokers(),
	)

	return app
}
