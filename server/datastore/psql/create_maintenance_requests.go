package psql

import (
    //"time"
    "github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

func (db *DB) CreateMaintenanceRequests(request datastore.MaintenanceRequest) (datastore.MaintenanceRequest, error) {
    insert, err := db.connection.Query("INSERT INTO maintenance_requests (request, user_submitted, date_submitted) VALUES ($1, $2, $3)", request.Request, request.UserSubmitted, request.DateSubmitted)
    if err != nil {
        return request, err
    }
    insert.Close() // Inserts the new request into the database

    row, err := db.connection.Query("SELECT * FROM maintenance_requests") // Return the struct of the newly created request
    if err != nil {
         return request, err
    }
    for row.Next() {
        err := row.Scan(&request.ID,
        &request.Request,
        &request.UserSubmitted,
        &request.DateSubmitted)
        if err != nil {
            return request, err
        }
    }
    row.Close()

    return request, nil
}