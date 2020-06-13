package psql

import (
	"github.com/jackc/pgx"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

/* Create a Struct that implements datastore.DB*/
type DB struct {
	connection *pgx.Conn
}
/* Init establishes a reusable connection */
func (db *DB) Init(config *datastore.DatabaseConfig) (error) {
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
	db.connection = conn
	return nil
}
