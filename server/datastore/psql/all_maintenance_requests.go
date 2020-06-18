package psql

// Import the necessary packages

// AllMaintenanceRequests performs a SQL query to fetch all rows in the
// maintenance_requests table.
func (db *DB) AllMaintenanceRequests() {
	/* AllMaintenanceRequests takes in a pgx connection as a parameter */

	/* Using the pgx connection, run a SQL query to fetch all rows in the
	*  maintenance_requests table. For example:
	* app.db.Connection().QueryRow(`select * from nsl_maintenance where id = 1`).Scan(&request.ID, &request.Request, &request.UserSubmitted, &request.DateSubmitted)
	 */

	/* Return the data as-is */
		var dbRequest datastore.MaintenanceRequest

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
        fmt.Println(dbRequest.ID, 
        dbRequest.Request, 
        dbRequest.UserSubmitted, 
        dbRequest.DateSubmitted)
    }
    row.Close()
    return nil
}
