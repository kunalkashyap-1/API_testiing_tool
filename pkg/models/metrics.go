package models

import "time"

// Metrics represents aggregated metrics for a test.
type Metrics struct {
	ID              int       `json:"id"`
	TestID          int       `json:"test_id"`
	RequestCount    int       `json:"request_count"`
	AvgResponseTime float64   `json:"avg_response_time"`
	ErrorRate       float64   `json:"error_rate"`
	Throughput      float64   `json:"throughput"`
	Timestamp       time.Time `json:"timestamp"`
}
