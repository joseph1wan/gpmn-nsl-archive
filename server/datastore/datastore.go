package datastore

import (
	"time"
)

/* DB provides functions to interact with the postgres database. Interactions
are defined in the psql package */
type DB interface {
	/* Implement the Init() function that connects to the postgres DB */
	Init() error // Modify this function, see below

	/* Init should take the config.yml file loaded in main.go as a parameter.
	*  Use a structe (like in auth.go) to validate the fields defined in
	*  config.yml. */

	/* Using the fields in config.yml, create a pgx config variable to store your
	*  postgres credentials and connection options. */

	/* Connect to the postgres DB using the pgx config variable, and store the
	*  connection in a variable that can be used by other functions (defined in
	*  the psql package) to interact with the postgres database. */

}

// MaintenanceRequest models the maintenance_request table
type MaintenanceRequest struct {
	ID            int
	Request       string
	UserSubmitted int
	DateSubmitted time.Time
}
