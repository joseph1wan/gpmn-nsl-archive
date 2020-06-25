package psql

//DeleteMaintenanceRequest deletes the row of the maintenance request in the database with the specified id endpoint
func (db *DB) DeleteMaintenanceRequest(id int) (int, error) {
    err := db.connection.QueryRow("DELETE FROM maintenance_requests WHERE id=$1 RETURNING id", id).Scan(&id)
    if err != nil {
        return id, err
    }
    return id, err
}
