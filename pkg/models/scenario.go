package models

import "time"

type Config struct {
	RequestCount int               `json:"request_count"`
	Concurrency  int               `json:"concurrency"`
	Timeout      int               `json:"timeout"` // in seconds
	Headers      map[string]string `json:"headers"`
	Query        map[string]string `json:"query"`
	Params       map[string]string `json:"params"`
	Payload      string            `json:"payload"` // could be JSON, XML, etc.
	Method       string            `json:"method"`  // GET, POST, PUT, etc.
}

type Scenario struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Name          int       `json:"name"`
	Description   string    `json:"description"`
	Config        Config    `json:"config"`
	CreatedBy     int       `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
}
