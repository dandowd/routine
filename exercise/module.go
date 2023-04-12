package exercise

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewExersiceTemplateRepo),
	fx.Provide(NewExerciseService),
	fx.Invoke(RegisterRoutes),
)
