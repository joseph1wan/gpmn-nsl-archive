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
	// Fetch all the rows from the DB
	//fmt.Println("BLAHHHH")

	rows, err := db.connection.Query("select * from public.nsl_maintenance LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	/*
		fmt.Println("-------------------")
		fmt.Println(rows)
		fmt.Println("-------------------")
	*/
	// Create a new empty array of MaintenanceRequests
	// INSERT CODE HERE
	//slice := make([]MaintenanceRequest)
	// For each row

	var D_Store []datastore.MaintenanceRequest
	defer rows.Close()
	for rows.Next() {

		// Create a new MaintenanceRequest
		var request2 datastore.MaintenanceRequest
		// INSERT CODE HERE
		// Scan the row values into MaintenanceRequest
		err = rows.Scan(&request2.ID, &request2.Request, &request2.UserSubmitted, &request2.DateSubmitted)

		if err != nil {
			// handle this error
			panic(err)
		}
		//fmt.Println(request2.ID, request2.Request, request2.UserSubmitted, request2.DateSubmitted)

		// Add the new MaintenanceRequest to the existing array of MaintenanceRequests
		D_Store = append(D_Store, request2)
		// INSERT CODE HERE

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// Return my array of MaintenanceRequests
	// INSERT CODE HERE
	return D_Store
}
