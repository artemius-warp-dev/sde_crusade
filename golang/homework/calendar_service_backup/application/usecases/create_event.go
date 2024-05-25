package usecases

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/dates_handlers"
	"calendar_service/domain/value_objects/event"
	"calendar_service/frameworks/logger"
	"calendar_service/frameworks/storage"
	"calendar_service/interfaces/queue"
	storage_interface "calendar_service/interfaces/storage"
	"fmt"
)

// TODO move to one place
var globalStorage storage_interface.EventStorage
var EventQueue queue.EventQueue

type DateParams struct {
	Day         string
	StartOfWeek string
	EndOfWeek   string
	Month       int
	Year        int
}

// TODO move to init func
func SetStorage(options *configs.Storage) {
	switch options.Type {
	case "memory":
		globalStorage = &storage.InMemoryEventStorage{}
	case "database":
		var err error
		globalStorage, err = storage.NewPostgreSQLEventStorage(*options.ConnStr)
		if err != nil {
			logger.Fatal("Connection error ", err)
		}
	}

	logger.Info("Storage is set")
}

func AddEventToStorage(event *event.Event) error {

	// Call the AddEvent method
	err := globalStorage.AddEvent(event)
	if err != nil {
		// Handle error
		return err
	}
	return nil
}

func UpdateEventFromStorage(id string, event *event.Event) error {

	// Call the AddEvent method
	err := globalStorage.UpdateEvent(id, event)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func DeleteEventFromStorage(id string) error {

	// Call the AddEvent method
	err := globalStorage.DeleteEvent(id)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func ListEventsFromStorage(_type string, dateParam DateParams) ([]*event.Event, error) {
	fmt.Println(dateParam)
	// Call the AddEvent method
	events, err := globalStorage.ListEvents()
	switch _type {
	case "day":
		events = dates_handlers.GetEventsForDay(events, dateParam.Day)
	case "week":
		events = dates_handlers.GetEventsForWeek(events, dateParam.StartOfWeek, dateParam.EndOfWeek)
	case "month":
		events = dates_handlers.GetEventsForMonth(events, dateParam.Month, dateParam.Year)
	}

	if err != nil {
		// Handle error
		return nil, err
	}

	return events, nil
}
