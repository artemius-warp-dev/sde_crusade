package tests

import (
	"testing"
	impl "calendar_service/frameworks/storage"
	header "calendar_service/interfaces/storage"
	e "calendar_service/domain/value_objects/event"
)

var _ header.EventStorage = (*impl.InMemoryEventStorage)(nil)

func TestEventStorage_AddEvent(t *testing.T) {
	
	storage := impl.InMemoryEventStorage{
		Events: []*e.Event{},
	}

	// Define test event
	event := &e.Event{
		ID:          "123",
		Title:       "Test Event",
		Description: "Test Description",
		StartTime:   "2024-04-20T09:00:00Z",
		EndTime:     "2024-04-20T10:00:00Z",
		Location:    "Test Location",
		Attendees:   []string{"test@example.com"},
	}

	// Add event to storage
	err := storage.AddEvent(event)
	if err != nil {
		t.Errorf("AddEvent() failed: %v", err)
	}

	if event != storage.Events[0] {
		t.Errorf("Events doesn't the same")
	}
}

// TestDeleteEvent tests the DeleteEvent method of InMemoryEventStorage.
func TestDeleteEvent(t *testing.T) {
    // Initialize the in-memory storage
    storage := impl.InMemoryEventStorage{
        Events: []*e.Event{
            // Initialize some events for testing
            {ID: "1", Title: "Event 1", Description: "Description 1"},
            {ID: "2", Title: "Event 2", Description: "Description 2"},
        },
    }

    // Call DeleteEvent method
    err := storage.DeleteEvent("1")
    if err != nil {
        t.Errorf("DeleteEvent() failed: %v", err)
    }

    // Check if the event was deleted correctly
    if len(storage.Events) != 1 || storage.Events[0].ID != "2" {
        t.Errorf("DeleteEvent() did not delete the event correctly")
    }
}

// TestListEvents tests the ListEvents method of InMemoryEventStorage.
func TestListEvents(t *testing.T) {
    // Initialize the in-memory storage
    storage := impl.InMemoryEventStorage{
        Events: []*e.Event{
            // Initialize some events for testing
            {ID: "1", Title: "Event 1", Description: "Description 1"},
            {ID: "2", Title: "Event 2", Description: "Description 2"},
        },
    }

    // Call ListEvents method
    events, err := storage.ListEvents()
    if err != nil {
        t.Errorf("ListEvents() failed: %v", err)
    }

    // Check if the events were listed correctly
    if len(events) != 2 {
        t.Errorf("ListEvents() did not return the correct number of events")
    }
}
