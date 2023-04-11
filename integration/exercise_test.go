package integration_test

import (
	"context"
	"net/http"
	"os"
	"routine/builder"
	"strings"
	"testing"

	"go.uber.org/fx"
)

func buildIntegrationOptions() fx.Option {
	return fx.Options()

}

func TestMain(m *testing.M) {
	cxt := context.Background()

	app := builder.AppBuilderWithOptions(buildIntegrationOptions())

	startErr := app.Start(cxt)

	if startErr != nil {
		panic(startErr)
	}

	code := m.Run()

	stopErr := app.Stop(cxt)

	if stopErr != nil {
		panic(stopErr)
	}

	os.Exit(code)
}

func TestPostExerciseShouldFailWithNoBody(t *testing.T) {
	// Given
	client := http.Client{}
	// When
	res, err := client.Post("http://localhost:8080/exercise", "application/json", nil)
	// Then
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestPostExerciseShouldPassWithBody(t *testing.T) {
	// Given
	client := http.Client{}
	// When
	res, err := client.Post("http://localhost:8080/exercise", "application/json", strings.NewReader(`{"name": "bench press", "reps": 5}`))
	// Then
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res.StatusCode)
	}
}
