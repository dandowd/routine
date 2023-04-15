package integration_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"routine/builder"
	"strings"
	"testing"

	"go.uber.org/fx"
)

func buildIntegrationOptions() fx.Option {
	return fx.Options(
		fx.Replace(NewAWSIntegrationTestConfig()),
		fx.Invoke(CreateTables),
	)

}

func TestMain(m *testing.M) {
	cxt := context.Background()

	dbContainer := NewDbContainer()

	app := builder.AppBuilderWithOptions(buildIntegrationOptions())

	startErr := app.Start(cxt)

	if startErr != nil {
		fmt.Printf("m: %v\n", startErr)
	}

	m.Run()

	stopErr := app.Stop(cxt)

	dbContainer.Cleanup()

	if stopErr != nil {
		fmt.Printf("m: %v\n", stopErr)
	}

	os.Exit(0)
}

func TestPostExerciseShouldFailWithNoBody(t *testing.T) {
	// Given
	client := http.Client{}
	// When
	res, err := client.Post("http://localhost:8080/exercise", "application/json", strings.NewReader(`{}`))
	// Then
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, res.StatusCode)
	}

	body, _ := ioutil.ReadAll(res.Body)

	t.Logf("Response Body: %s", body)
}

func TestPostExerciseShouldPassWithBody(t *testing.T) {
	// Given
	client := http.Client{}
	// When
	res, err := client.Post("http://localhost:8080/exercise", "application/json", strings.NewReader(`{"name": "bench press", "description": "flat"}`))
	// The
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, res.StatusCode)
	}
}
