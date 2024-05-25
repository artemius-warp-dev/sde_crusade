-- Create a table to store events
CREATE TABLE events (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    location VARCHAR(255),
    attendees TEXT[]
);

-- Create an index on the id column
CREATE INDEX idx_events_id ON events (id);

-- Create an index on the start_time column
CREATE INDEX idx_events_start_time ON events (start_time);

-- Create an index on the end_time column
CREATE INDEX idx_events_end_time ON events (end_time);
