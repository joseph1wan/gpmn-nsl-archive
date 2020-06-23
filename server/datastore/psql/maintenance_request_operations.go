package psql

import (
    "github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

func (db *DB) AllMaintenanceRequests() ([]datastore.MaintenanceRequest, error) {
	var maintenanceRequestTable []datastore.MaintenanceRequest
	rows, err := db.connection.Query("SELECT * FROM maintenance_requests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		// Create a new MaintenanceRequest
		var maintenanceRequest datastore.MaintenanceRequest
		// Scan the row values into MaintenanceRequest
		err = rows.Scan(&maintenanceRequest.ID, &maintenanceRequest.Request, &maintenanceRequest.UserSubmitted, &maintenanceRequest.DateSubmitted)
		//get any errors before the first iteration, checks if DB table is spelt correctly
		if err != nil {
			return nil, err
		}
		// Add the new MaintenanceRequest to the existing array of MaintenanceRequests
		maintenanceRequestTable = append(maintenanceRequestTable, maintenanceRequest)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// Return my array of MaintenanceRequests
	return maintenanceRequestTable, nil
}

func (db *DB) CreateMaintenanceRequest(request datastore.MaintenanceRequest) (*datastore.MaintenanceRequest, error) {
    var lastID int
    err := db.connection.QueryRow("INSERT INTO maintenance_requests (request, user_submitted, date_submitted) VALUES ($1, $2, $3) RETURNING id",
    request.Request, request.UserSubmitted, request.DateSubmitted).Scan(&lastID)
    if err != nil {
        return nil, err
    }
    request.ID = lastID
    return &request, nil
}