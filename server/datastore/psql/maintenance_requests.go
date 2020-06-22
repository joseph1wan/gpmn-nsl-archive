package psql

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

// Import the necessary packages

func (db *DB) Requests() ([]datastore.MaintenanceRequest, error) {
	var maintenanceRequestTable []datastore.MaintenanceRequest
	rows, err := db.connection.Query("SELECT * FROM maintenance_requests LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		return maintenanceRequestTable, err
	}

	defer rows.Close()
	for rows.Next() {
		// Create a new MaintenanceRequest
		var maintenance_request datastore.MaintenanceRequest
		// Scan the row values into MaintenanceRequest
		err = rows.Scan(&maintenance_request.ID, &maintenance_request.Request, &maintenance_request.UserSubmitted, &maintenance_request.DateSubmitted)
		if err != nil {
			return maintenanceRequestTable, err
		}
		// Add the new MaintenanceRequest to the existing array of MaintenanceRequests
		maintenanceRequestTable = append(maintenanceRequestTable, maintenance_request)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return maintenanceRequestTable, err
	}
	// Return my array of MaintenanceRequests
	return maintenanceRequestTable, nil
}
