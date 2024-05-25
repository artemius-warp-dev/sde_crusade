package storage

import (
	e "calendar_service/domain/value_objects/event"
)

type EventStorage interface {
	AddEvent(event *e.Event) error

	UpdateEvent(eventID string, updatedEvent *e.Event) error

	DeleteEvent(eventID string) error

	ListEvents() ([]*e.Event, error)

	FetchEventsForNotification() ([]*e.Event, error)

	MarkEventsAsNotified(eventID string) error
}
