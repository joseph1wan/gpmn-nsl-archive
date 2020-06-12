package datastore

import (
	"time"
)

/* DB is an interface for the app database. Any struct used as the app database
* must implement these functions */
type DB interface {
	Init(config *ServerConfig) error // Initialize pgx connection
}

type ServerConfig struct {
	Port int `yaml:"port"`
	Host string `yaml:"host"`
	User string `yaml: "user"`
	Password string `yaml: "password"`
	Dbname string `yaml: "dbname"`
}
// MaintenanceRequest models the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
