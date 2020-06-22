// This file provides NewApp() with functions for maintenance requests
package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"time"

	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/gin-gonic/gin"
)

// getMaintenanceRequests returns all the maintenance_requests from the database
func (app *App) getMaintenanceRequests(c *gin.Context) {
	requests, err := app.db.AllMaintenanceRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error!, Check to make sure your DB table is spelt correctly": err.Error()})
		return
	}
	if requests == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Possible error!": "Your Table is Empty"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"maintenance_requests": requests})
}

// CreateMaintenanceRequests creates a maintenance_request in the database
func (app *App) CreateMaintenanceRequest(c *gin.Context) {
	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var newRequest datastore.MaintenanceRequest
	err = json.Unmarshal(byteData, &newRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newRequest.Request == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check that the request field is not blank or missing."})
		return
	}
	newRequest.DateSubmitted = time.Now() //hard coded time submitted for now
	newRequest.UserSubmitted = 101        // hard coded user ID for now. I don't know how or where to read in a user ID
	req, err := app.db.CreateMaintenanceRequest(newRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": req.ID})
}
