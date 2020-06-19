// This file provides NewApp() with functions for maintenance requests
package main

import (
	"github.com/gin-gonic/gin"
)

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *App) GetMaintenanceRequests(c *gin.Context) {
	//id, _ := c.Param.Get("id")
	app.db.AllMaintenanceRequests()
	//fmt.Println(j)

	//app.db.AllMaintenanceRequests()
	/* Get data using psql.AllMaintenanceRequests and pass in app.db for the
	* connection */

	/* Takes the data returned in the psql func and formats it to a valid format
	* using the gin.H func. See if it you can just use a Struct */

	/* Use the gin.context.JSON func to return a status and the data. See auth.go
	* for an example */
}
