// The datastore package defines the database interface and all the Structs used
// by the database.
package datastore

import (
	"time"
)

// DB is an interface for the app database.
// Any Struct used as the app database must implement these functions.
type DB interface {
	Init(config DatabaseConfig) error // Initialize pgx connection
	CreateMaintenanceRequest(request MaintenanceRequest) (*MaintenanceRequest, error)
}

// DatbaseConfig models the fields needed to connect to the database.
type DatabaseConfig struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// MaintenanceRequest models the maintenance_requests table.
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
