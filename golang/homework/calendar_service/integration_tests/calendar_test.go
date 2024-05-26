package integration_tests

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestCalendar(m *testing.T) {

	status := godog.TestSuite{
		Name:                 "calendar",
		TestSuiteInitializer: nil,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if status == 2 {
		m.SkipNow()
	}

	if status != 0 {
		m.Fatalf("zero status code expected, %d received", status)
	}

}
