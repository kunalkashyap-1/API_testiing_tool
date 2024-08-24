package models

import "time"

type ReportData struct {
	RequestCount    int            `json:"request_count"`
	SuccessCount    int            `json:"success_count"`
	FailureCount    int            `json:"failure_count"`
	AvgResponseTime float64        `json:"avg_response_time"`
	MinResponseTime float64        `json:"min_response_time"`
	MaxResponseTime float64        `json:"max_response_time"`
	ErrorDetails    []ErrorDetail  `json:"error_details"`
	ResponseTimes   map[string]int `json:"response_times"`
}

// ErrorDetail represents the details of an error encountered during the test.
type ErrorDetail struct {
	Timestamp  time.Time `json:"timestamp"`   // When the error occurred
	Message    string    `json:"message"`     // Error message
	Endpoint   string    `json:"endpoint"`    // The endpoint that caused the error
	StatusCode int       `json:"status_code"` // The HTTP status code returned
}

// Report represents a generated report for a test.
type Report struct {
	ID          int        `json:"id"`
	TestID      int        `json:"test_id"`
	ReportData  ReportData `json:"report_data"` // Nested struct for detailed report data
	GeneratedAt time.Time  `json:"generated_at"`
}
