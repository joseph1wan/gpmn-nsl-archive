package psql

import (
    "github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

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