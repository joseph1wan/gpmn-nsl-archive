// This file provides NewApp() with functions for maintenance requests
package main

import(
	"github.com/a2fumn2022/gpmn-nsl/server/datastore/psql"
)

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *app) AllMaintenanceRequests() {
	/* Get data using psql.AllMaintenanceRequests and pass in app.db for the
	* connection */

	/* Takes the data returned in the psql func and formats it to a valid format
	* using the gin.H func. See if it you can just use a Struct */

	/* Use the gin.context.JSON func to return a status and the data. See auth.go
	* for an example */
}
