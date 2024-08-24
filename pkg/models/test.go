package models

import "time"

// Test represents a test execution.
type Test struct {
	ID          int       `json:"id"`
	ScenarioID  int       `json:"scenario_id"`
	UserID      int       `json:"user_id"`
	Status      string    `json:"status"` // e.g., 'pending', 'running', 'completed', 'aborted'
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// TestResult represents a result of a test.
type TestResult struct {
	ID          int       `json:"id"`
	TestID      int       `json:"test_id"`
	MetricType  string    `json:"metric_type"` // e.g., 'response_time', 'error_rate', 'throughput'
	MetricValue float64   `json:"metric_value"`
	Timestamp   time.Time `json:"timestamp"`
}
