package exercise

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewExersiceRepo),
	fx.Provide(NewExerciseService),
	fx.Invoke(RegisterRoutes),
)
