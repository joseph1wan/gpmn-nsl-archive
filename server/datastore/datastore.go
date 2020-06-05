package datastore

import (
	"time"
)

// Placeholder while the DB struct gets implemented
type DB interface {
	Init() error
}

// Struct to represent the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
