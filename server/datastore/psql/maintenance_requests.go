package psql

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

func (db *DB) AllMaintenanceRequests() ([]datastore.MaintenanceRequest, error) {
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
		//get any errors before the first iteration, checks if DB table is spelt correctly
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
