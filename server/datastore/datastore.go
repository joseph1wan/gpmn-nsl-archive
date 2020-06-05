package datastore

import (
	"time"
)

type DB interface {
	Init() error
}

type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
