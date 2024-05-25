package validators

import (
	e "calendar_service/domain/value_objects/event"
	"errors"
	"regexp"
)

func ValidateEvent(event *e.Event) error {
	if event.ID == "" {
		return errors.New("ID is required")
	}

	if event.Title == "" {
		return errors.New("Title is required")
	}

	if event.StartTime == "" {
		return errors.New("Start time is required")
	}

	if event.EndTime == "" {
		return errors.New("End time is required")
	}

	if event.Location == "" {
		return errors.New("Location is required")
	}

	if len(event.Attendees) == 0 {
		return errors.New("At least one attendee is required")
	}

	for _, attendee := range event.Attendees {
		if !isValidEmail(attendee) {
			return errors.New("Invalid email format for attendee: " + attendee)
		}
	}

	return nil
}

func isValidEmail(email string) bool {
	// Simple email format validation using regular expression
	// This is a basic validation, you may need to adjust it based on your requirements
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}
