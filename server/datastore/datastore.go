package datastore

import (
	"time"
)

/* DB provides functions to interact with the postgres database. Interactions
are defined in the psql package */
type DB interface {
	/* Implement the Init() function that connects to the postgres DB */
	Init() error // Modify this function

	/* Instructions */

	  /* Create a config variable to store your postgres credentials and connection
	  options. Should use variables from the config.yml file loaded in main.go
	  (so you need to pass that into Init) */

	  /* Connect to the postgres DB using pgx and return the connection so that the
	  app can call this interface to connect to the DB */
}

// Struct to represent the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
