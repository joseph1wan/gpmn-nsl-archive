package psql

import (
	"fmt"

	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/jackc/pgx"
	// Import the necessary packages
)

/* AllMaintenanceRequests performs a SQL query to fetch all rows in the
*  maintenance_requests table */
func AllMaintenanceRequests(conn *pgx.Conn) {
	var request2 datastore.MaintenanceRequest
	fmt.Println(" AllMaintenanceRequests")
	rows, err := conn.Query("select * from public.nsl_maintenance LIMIT $1", 777) //Number limits # of requests, can be anything
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&request2.ID, &request2.Request, &request2.UserSubmitted, &request2.DateSubmitted)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(request2.ID, request2.Request, request2.UserSubmitted, request2.DateSubmitted)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	/* Using the pgx connection, run a SQL query to fetch all rows in the
	*  maintenance_requests table. For example:
	* app.db.Connection().QueryRow(`select * from nsl_maintenance where id = 1`).Scan(&request.ID, &request.Request, &request.UserSubmitted, &request.DateSubmitted)
	 */

	/* Return the data as-is */
}
