package storage

import (
	"database/sql"
	"fmt"

	e "calendar_service/domain/value_objects/event"

	"github.com/lib/pq"
)

// PostgreSQLEventStorage represents a PostgreSQL implementation of the EventStorage interface
type PostgreSQLEventStorage struct {
	db *sql.DB
}

//TODO replace fmt to logger

// NewPostgreSQLEventStorage creates a new PostgreSQLEventStorage instance with the provided database connection string
func NewPostgreSQLEventStorage(connStr string) (*PostgreSQLEventStorage, error) {
	// Open a connection to the database using the connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Return a new instance of PostgreSQLEventStorage with the connected database
	return &PostgreSQLEventStorage{db: db}, nil
}

// AddEvent adds a new event to the PostgreSQL database
func (ps *PostgreSQLEventStorage) AddEvent(event *e.Event) error {
	query := "INSERT INTO events (id, title, description, start_time, end_time, location, attendees) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := ps.db.Exec(query, event.ID, event.Title, event.Description, event.StartTime, event.EndTime, event.Location, pq.Array(event.Attendees))
	if err != nil {
		return fmt.Errorf("failed to add event: %v", err)
	}
	return nil
}

// UpdateEvent updates an existing event in the PostgreSQL database
func (ps *PostgreSQLEventStorage) UpdateEvent(eventID string, updatedEvent *e.Event) error {
	query := "UPDATE events SET title=$1, description=$2, start_time=$3, end_time=$4, location=$5, attendees=$6 WHERE id=$7"
	_, err := ps.db.Exec(query, updatedEvent.Title, updatedEvent.Description, updatedEvent.StartTime, updatedEvent.EndTime, updatedEvent.Location, pq.Array(updatedEvent.Attendees), eventID)
	if err != nil {
		return fmt.Errorf("failed to update event: %v", err)
	}
	return nil
}

// DeleteEvent deletes an event from the PostgreSQL database
func (ps *PostgreSQLEventStorage) DeleteEvent(eventID string) error {
	query := "DELETE FROM events WHERE id=$1"
	_, err := ps.db.Exec(query, eventID)
	if err != nil {
		return fmt.Errorf("failed to delete event: %v", err)
	}
	return nil
}

// ListEvents retrieves all events from the PostgreSQL database
func (ps *PostgreSQLEventStorage) ListEvents() ([]*e.Event, error) {
	var events []*e.Event

	rows, err := ps.db.Query("SELECT id, title, description, start_time, end_time, location, attendees FROM events")
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event e.Event
		var attendees pq.StringArray

		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.StartTime, &event.EndTime, &event.Location, &attendees)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event row: %v", err)
		}

		event.Attendees = attendees
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over events: %v", err)
	}

	return events, nil
}

func (ps *PostgreSQLEventStorage) FetchEventsForNotification() ([]*e.Event, error) {
	rows, err := ps.db.Query("SELECT id, title, description, start_time, end_time, location, attendees FROM events WHERE start_time <= NOW() AND notified = FALSE")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []*e.Event
	for rows.Next() {
		var event e.Event
		var attendees pq.StringArray
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.StartTime, &event.EndTime, &event.Location, &attendees); err != nil {
			return nil, err
		}
		event.Attendees = attendees
		events = append(events, &event)
	}
	return events, nil
}

func (ps *PostgreSQLEventStorage) MarkEventsAsNotified(eventID string) error {
	_, err := ps.db.Exec("UPDATE events SET notified = TRUE WHERE id = $1", eventID)
	return err
}
