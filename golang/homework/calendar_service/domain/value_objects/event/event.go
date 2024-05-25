package event

import "fmt"

type Event struct {
	Attendees   []string `json:"attendees"`
	Description string   `json:"description"`
	EndTime     string   `json:"end_time"`
	ID          string   `json:"id"`
	Location    string   `json:"location"`
	StartTime   string   `json:"start_time"`
	Title       string   `json:"title"`
}

func (e Event) String() string {
	return fmt.Sprintf("ID: %s, Title: %s, Description: %s, StartTime: %s, EndTime: %s, Location: %s, Attendees: %v",
		e.ID, e.Title, e.Description, e.StartTime, e.EndTime, e.Location, e.Attendees)
}
