package psql

import (
	"time"

	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

// Import the necessary packages
type rawTime []byte

func (t rawTime) Time() (time.Time, error) {
	return time.Parse("15:04:05", string(t))
}

func (db *DB) AllMaintenanceRequests() []datastore.MaintenanceRequest {

	rows, err := db.connection.Query("select * from public.nsl_maintenance LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		// handle this error better than this
		panic(err)
	}

	var D_Store []datastore.MaintenanceRequest
	defer rows.Close()
	for rows.Next() {

		// Create a new MaintenanceRequest
		var request2 datastore.MaintenanceRequest

		// Scan the row values into MaintenanceRequest
		err = rows.Scan(&request2.ID, &request2.Request, &request2.UserSubmitted, &request2.DateSubmitted)

		if err != nil {
			// handle this error
			panic(err)
		}

		// Add the new MaintenanceRequest to the existing array of MaintenanceRequests
		D_Store = append(D_Store, request2)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// Return my array of MaintenanceRequests
	return D_Store
}
