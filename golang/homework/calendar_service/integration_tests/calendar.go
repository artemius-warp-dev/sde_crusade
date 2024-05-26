package integration_tests

import (
	"flag"
	"os"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	e "calendar_service/domain/value_objects/event"
	"calendar_service/frameworks/storage"
	"errors"
)

var (
	calendar *storage.InMemoryEventStorage
)

var opts = godog.Options{
	Output:      colors.Colored(os.Stdout),
	Concurrency: 4,
	Format:      "pretty",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func anEmptyCalendar() error {
	calendar = &storage.InMemoryEventStorage{
		Events: []*e.Event{},
	}
	return nil
}

func iAddAnEventWithTitleAt(title, eventTime string) error {
	event := &e.Event{
		ID:        "1",
		Title:     title,
		StartTime: eventTime,
	}
	return calendar.AddEvent(event)
}

func theEventShouldBeAddedToTheCalendar(title string) error {
	events, _ := calendar.ListEvents()
	for _, event := range events {
		if event.Title == title {
			return nil
		}
	}
	return errors.New("event not found")
}

func aCalendarWithEvents() error {
	return iAddAnEventWithTitleAt("Meeting", "2024-05-26 10:00")
}

func iRequestEventsForTheDay(date string) error {
	return nil
}

func iShouldGetTheEvents(expected *godog.Table) error {
	events, _ := calendar.ListEvents()
	expectedEvents := map[string]string{}
	for _, row := range expected.Rows {
		expectedEvents[row.Cells[0].Value] = row.Cells[1].Value
	}

	for _, event := range events {
		expectedTime, ok := expectedEvents[event.Title]
		if !ok || event.StartTime != expectedTime {
			return errors.New("event mismatch")
		}
	}
	return nil
}

func aCalendarWithAnEvent() error {
	return iAddAnEventWithTitleAt("Meeting", "2024-05-26 10:00")
}

func theEventTimeIs(eventTitle, eventTime string) error {
	return nil
}

func aNotificationShouldBeSentForTheEvent(eventTitle string) error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^an empty calendar$`, anEmptyCalendar)
	ctx.Step(`^I add an event with title "([^"]*)" at "([^"]*)"$`, iAddAnEventWithTitleAt)
	ctx.Step(`^the event "([^"]*)" should be added to the calendar$`, theEventShouldBeAddedToTheCalendar)
	ctx.Step(`^a calendar with events$`, aCalendarWithEvents)
	ctx.Step(`^I request events for the day "([^"]*)"$`, iRequestEventsForTheDay)
	ctx.Step(`^I should get the events:$`, iShouldGetTheEvents)
	ctx.Step(`^a calendar with an event$`, aCalendarWithAnEvent)
	ctx.Step(`^the event "([^"]*)" time is "([^"]*)"$`, theEventTimeIs)
	ctx.Step(`^a notification should be sent for the event "([^"]*)"$`, aNotificationShouldBeSentForTheEvent)
}
