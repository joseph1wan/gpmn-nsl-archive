package psql

import (
	"github.com/jackc/pgx"
)

/* Create a Struct that implements datastore.DB*/
type DB struct {
	conn *pgx.Conn
}

/* Init establishes a reusable connection */
func (db *DataBase) Init() (error) {
	// The hard-coded params should be replaced by the config.yml file
	// You'll need to update the interface to accept a param also
	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "postgres",
		Host:              "localhost",
		Port:              5432,
		Database:          "nsl_dev",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

/* Connection fetches an establish pgx connection */
func (db *DataBase) Connection() (*pgx.Conn) {
	return db.Connection
}
