package psql

import (
	"github.com/jackc/pgx"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

/* Create a Struct that implements datastore.DB*/
type DB struct {
	Connection *pgx.Conn
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
	connection, err := pgx.Connect(connConfig)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	db.Conn = connection
=======
	db.Connection = conn
>>>>>>> 1d42b59b9e05b48406542a5aa2136618bf3ab152
	return nil
}
