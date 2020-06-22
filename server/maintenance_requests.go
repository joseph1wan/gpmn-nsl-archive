// This file provides NewApp() with functions for maintenance requests
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMaintenanceRequests returns all the maintenance_requests from the database
// GetMainte
func (app *App) getMaintenanceRequests(c *gin.Context) {

	req, _ := app.db.Requests()

	c.JSON(http.StatusOK, gin.H{"maintenance_requests": req})

}
