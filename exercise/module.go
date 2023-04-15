package exercise

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewExersiceTemplateTable),
	fx.Provide(NewExerciseService),
	fx.Invoke(RegisterRoutes),
)
