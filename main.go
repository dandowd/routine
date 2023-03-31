package main

import (
	"routine/common"
	"routine/exercise"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(NewHTTPServer),
		fx.Provide(common.NewLogger),
		fx.Provide(exercise.NewExerciseService),

		fx.Invoke(exercise.RegisterRoutes),
	)

	app.Run()
}
