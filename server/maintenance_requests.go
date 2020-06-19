// This file provides NewApp() with functions for maintenance requests
package main

import(
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
)

type CreateRequest struct {
	Request string `json:"request"`
}

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *App) AllMaintenanceRequests(c *gin.Context) {
	/* Get data using psql.AllMaintenanceRequests and pass in app.db for the
	* connection */

	/* Takes the data returned in the psql func and formats it to a valid format
	* using the gin.H func. See if it you can just use a Struct */

	/* Use the gin.context.JSON func to return a status and the data. See auth.go
	* for an example */
}

func (app *App) CreateMaintenanceRequests(c *gin.Context) {
	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req CreateRequest
	err = json.Unmarshal(byteData, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := 101 // hard coded user ID for now. I don't know how or where to read in a user ID
	newID, err := app.db.CreateMaintenanceRequests(req.Request, userID, time.Now()) // either parse time here to mm/dd/y format or in CreateMaintenanceRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
	}
	c.JSON(http.StatusOK, gin.H{"id": newID})
}