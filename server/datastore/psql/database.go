package psql

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/jackc/pgx"
)

/* Create a Struct that implements datastore.DB*/
type DB struct {
	conn *pgx.Conn
}

/* Init establishes a reusable connection */
func (db *DB) Init(config datastore.DatabaseConfig) error {
	connConfig := pgx.ConnConfig{
		User:              config.User,
		Password:          config.Password,
		Host:              config.Host,
		Port:              uint16(config.Port),
		Database:          config.Dbname,
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
func (db *DB) Connection() *pgx.Conn {
	return db.conn
}
