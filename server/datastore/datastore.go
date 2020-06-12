package datastore

import (
	"time"
	"github.com/jackc/pgx"
)

/* DB is an interface for the app database. Any struct used as the app database
* must implement these functions */
type DB interface {
	Init() error // Initialize pgx connection
	Connection() *pgx.Conn // Fetch established pgx connection
}


// MaintenanceRequest models the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Date
}
