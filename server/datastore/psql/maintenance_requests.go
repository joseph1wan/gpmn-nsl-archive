package psql

import (
	"fmt"
	"reflect"
)

// Import the necessary packages

func (db *DB) AllMaintenanceRequests() []MaintenanceRequest {
	// Fetch all the rows from the DB
	rows, err := db.connection.Query("select * from public.nsl_maintenance LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		// handle this error better than this
		panic(err)
	}

	// Create a new empty array of MaintenanceRequests
	// INSERT CODE HERE
	// For each row
	defer rows.Close()
	for rows.Next() {
		// Create a new MaintenanceRequest
		// INSERT CODE HERE

		// Scan the row values into MaintenanceRequest
		err = rows.Scan(&request2.ID, &request2.Request, &request2.UserSubmitted, &request2.DateSubmitted)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(request2.ID, request2.Request, request2.UserSubmitted, request2.DateSubmitted)
		fmt.Println(reflect.TypeOf(request2.ID))
		fmt.Println(reflect.TypeOf(request2.Request))

		// Add the new MaintenanceRequest to the existing array of MaintenanceRequests
		// INSERT CODE HERE

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// Return my array of MaintenanceRequests
	// INSERT CODE HERE
}
