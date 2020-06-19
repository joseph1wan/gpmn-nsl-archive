// This file provides NewApp() with functions for maintenance requests
package main

import (
	"github.com/gin-gonic/gin"
)

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *App) GetMaintenanceRequests(c *gin.Context) {
	// Get an array of MaintenanceRequests
	app.db.AllMaintenanceRequests()

	// Take array and format it into JSON format for gin.H
	// INSERT CODE HERE

	// Use the gin.context.JSON func to return a status and the JSON data
	//   Example: c.JSON(http.StatusOK, gin.H{"userID": user.ID, "authToken": token})
	// INSERT CODE HERE
}
