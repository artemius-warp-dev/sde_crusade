package storage

import (
	e "calendar_service/domain/value_objects/event"
	serr "calendar_service/interfaces/storage"
	"errors"
)

type InMemoryEventStorage struct {
	Events []*e.Event
}

func (s *InMemoryEventStorage) AddEvent(event *e.Event) error {
	s.Events = append(s.Events, event)
	return nil
}

func (s *InMemoryEventStorage) UpdateEvent(eventID string, updatedEvent *e.Event) error {
	for i, event := range s.Events {
		if event.ID == eventID {
			// Update the event
			s.Events[i] = updatedEvent
			return nil
		}
	}
	return serr.ErrEventNotFound
}

func (s *InMemoryEventStorage) DeleteEvent(eventID string) error {
	for i, event := range s.Events {
		if event.ID == eventID {
			// Delete the event
			s.Events = append(s.Events[:i], s.Events[i+1:]...)
			return nil
		}
	}
	return serr.ErrEventNotFound
}

func (s *InMemoryEventStorage) ListEvents() ([]*e.Event, error) {
	// Return a copy of events slice to prevent modifications to the original slice
	events := make([]*e.Event, len(s.Events))
	copy(events, s.Events)
	return events, nil
}

func (s *InMemoryEventStorage) FetchEventsForNotification() ([]*e.Event, error) {
	return nil, errors.New("Not implemented")
}

func (s *InMemoryEventStorage) MarkEventsAsNotified(eventID string) error {
	return errors.New("Not implemented")
}
