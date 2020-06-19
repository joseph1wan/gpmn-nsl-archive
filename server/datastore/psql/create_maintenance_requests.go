package psql

import (
    "time"
    "github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

func (db *DB) CreateMaintenanceRequests(request string, userID int, time time.Time) (int, error) {
    var dbRequest datastore.MaintenanceRequest
    lastID := 0
    row, err := db.connection.Query("SELECT * FROM maintenance_requests")
    if err != nil {
         return lastID, err
    }
    for row.Next() {
        err := row.Scan(&dbRequest.ID,
        &dbRequest.Request,
        &dbRequest.UserSubmitted,
        &dbRequest.DateSubmitted)
        if err != nil {
            return lastID, err
        }
    lastID = dbRequest.ID
    }
    row.Close()
    lastID++
    insert, err := db.connection.Query("INSERT INTO maintenance_requests VALUES ($1, $2, $3, $4)", lastID, request, userID, time)
    if err != nil {
        return lastID, err
    }
    insert.Close()
    return lastID, nil
}