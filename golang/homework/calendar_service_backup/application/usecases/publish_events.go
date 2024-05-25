package usecases

import (
	e "calendar_service/domain/value_objects/event"
	"encoding/json"
)

func PublishEvent(event *e.Event) error {
	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return EventQueue.Publish(EventQueue.GetQueue(), body)

}
func MarkEventAsNotified(eventID string) error {
	err := globalStorage.MarkEventsAsNotified(eventID)
	if err != nil {
		return err
	}

	return nil
}

func MarkEventsAsNotified(events []*e.Event) error {
	for _, event := range events {
		err := globalStorage.MarkEventsAsNotified(event.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
