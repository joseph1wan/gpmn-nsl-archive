package main

import (
	"flag"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore/psql"
	"github.com/gin-gonic/gin"
	"strconv"
	yaml "gopkg.in/yaml.v2"
)



type App struct {
	db datastore.DB
}

// Run the application
func main() {
	configFile := flag.String("config", "config.yaml", "Config file for nsl backend")
	config, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not configure server, %v", err)
		return
	}
	fmt.Println("Read configfile")
	app := NewApp(config)
	/* Example connection */

	// Create struct to hold data
	var request datastore.MaintenanceRequest

	// Connect to the database and run query
	fmt.Println("Trying to connect to database")
	app.db.Connection().QueryRow(`select * from public.nsl_maintenance where id = 1`).Scan(&request.ID, &request.Request, &request.UserSubmitted, &request.DateSubmitted)
	fmt.Println("Connected")
	// Print to test
	fmt.Println(request.ID)
	fmt.Println(request.Request)
	fmt.Println(request.UserSubmitted)
	fmt.Println(request.DateSubmitted)

	r := gin.Default()

	/* POST endpoint that calls app's Login function defined in auth.go */
	// TODO:
	//r.POST("/login", app.Login)

	/* GET endpoint that calls app's ___ function defined in maintenance.go */
	  // r.GET("/maintenance_requests", app.AllMaintenanceRequests())

	// NOTE: To add a group of endpoints with an authorized user, see the following commented out code
	//authorized := r.Group("/maintenance", gin.BasicAuth(AuthorizedUsers))
	//authorized.GET("/", func(c *gin.Context) {
	//	// To get the user's email..
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//	// do stuff with user...
	//	id := SimpleUserDB[user].ID
	//	...
	//})

	r.Run(":" + strconv.Itoa(config.Port))
}

// NewApp creates the app that holds all the functions to interact with the DB
func NewApp(config *datastore.ServerConfig) App {
	app := App{
    db: &psql.DB{},
	}
	err := app.db.Init(config)
	fmt.Println(err)
	if err != nil {
		log.Fatal("could not initialize database")
	}

	return app
}

// Read the config.yml file
func ReadConfig(file string) (*datastore.ServerConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Create a ServerConfig struct and read file info into the struct
	var config datastore.ServerConfig
	yaml.Unmarshal(data, &config)

	if config.Port == 0 {
		config.Port = 4000
	}
	return &config, nil
}
