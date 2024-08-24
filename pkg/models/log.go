package models

import "time"

// Log represents a log entry for a test.
type Log struct {
	ID         int       `json:"id"`
	TestID     int       `json:"test_id"`
	LogMessage string    `json:"log_message"`
	LogLevel   string    `json:"log_level"` // e.g., 'INFO', 'WARNING', 'ERROR'
	Timestamp  time.Time `json:"timestamp"`
}
