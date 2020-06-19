package psql

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/jackc/pgx"
)

/* Create a Struct that implements datastore.DB*/
type DB struct {
	connection *pgx.Conn
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
	db.connection = conn
	return nil
}

/*
func (db *DB) AllMaintenanceRequests() {
	var request2 datastore.MaintenanceRequest
	fmt.Println(" AllMaintenanceRequests")
	rows, err := db.connection.Query("select * from public.nsl_maintenance LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&request2.ID, &request2.Request, &request2.UserSubmitted, &request2.DateSubmitted)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(request2.ID, request2.Request, request2.UserSubmitted, request2.DateSubmitted)
		fmt.Println(reflect.TypeOf(request2.ID))
		fmt.Println(reflect.TypeOf(request2.Request))

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	//return nil
}
*/
