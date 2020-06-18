package psql

import (
    "time"
    "github.com/a2fumn2022/gpmn-nsl/server/datastore"
)

func (db *DB) CreateMaintenanceRequests(request string, userID int, time time.Time) error {
    var dbRequest datastore.MaintenanceRequest
    lastID := 0
    row, err := db.connection.Query("SELECT * FROM maintenance_requests")
    if err != nil {
         return err
    }
    for row.Next() {
        err := row.Scan(&dbRequest.ID,
        &dbRequest.Request,
        &dbRequest.UserSubmitted,
        &dbRequest.DateSubmitted)
        if err != nil {
            return err
        }
    lastID = dbRequest.ID
    }
    row.Close()
    if lastID == 0 {
        insert, err := db.connection.Query("INSERT INTO maintenance_requests VALUES (1, $1, $2, $3)", request, userID, time)
        if err != nil {
            return err
        }
        insert.Close()
    } else {
        insert, err := db.connection.Query("INSERT INTO maintenance_requests VALUES ($1, $2, $3, $4)", lastID + 1, request, userID, time)
        if err != nil {
            return err
        }
        insert.Close()
    }

    return nil
}