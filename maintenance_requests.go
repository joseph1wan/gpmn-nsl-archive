// This file provides NewApp() with functions for maintenance requests
package main

import(
	// Import the necessary packages
)

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *app) AllMaintenanceRequests() {
	// AllMaintenanceRequests takes a gin context as a parameter */

	/* Use the psql package to fetch the requests. You will need to provide the
	*  initialized DB (see NewApp() in main.go) as a parameter to the psql 
	*  function */

	// Using the gin.JSON() function, return the requests in a valid JSON format
}
