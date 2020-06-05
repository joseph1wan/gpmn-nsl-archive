package psql

import (
	"fmt"
	"os"
	"github.com/jackc/pgx"
)

/* Create a Struct that can connect with the database.*/
type DataBase struct {
	host, password, dbname string
	port int32

	/* Implement the Init() function that connects to the postgres DB */

	/* Instructions */

	  /* Create a config variable to store your postgres credentials and connection
	  options. Should use variables from the config.yml file loaded in main.go
	  (so you need to pass that into Init) */

	  /* Connect to the postgres DB using pgx and return the connection so that the
	  app can call this interface to connect to the DB */
}
