// This file provides NewApp() with functions for maintenance requests
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Request string `json:"request"`
}

/* AllMaintenanceRequests returns all the maintenance_requests from the database
*  AllMaintenanceRequests is a function of app */
func (app *app) GetMaintenanceRequests(c *gin.Context) {
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
	app.db.AllMaintenanceRequests()
	fmt.Println(app.db.AllMaintenanceRequests)

	//fmt.Println(j)

	//app.db.AllMaintenanceRequests()
	/* Get data using psql.AllMaintenanceRequests and pass in app.db for the
	* connection */

	/* Takes the data returned in the psql func and formats it to a valid format
	* using the gin.H func. See if it you can just use a Struct */

	/* Use the gin.context.JSON func to return a status and the data. See auth.go
	* for an example */
}
