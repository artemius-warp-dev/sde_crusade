package usecases

import (
	e "calendar_service/domain/value_objects/event"
)

func FetchEventsForNotification() ([]*e.Event, error) {
	return globalStorage.FetchEventsForNotification()
}
