package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore/psql"
	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port     int                      `yaml:"port"`
	DBConfig datastore.DatabaseConfig `yaml:"database"`
}

type App struct {
	conf   ServerConfig
	db     datastore.DB
	server *gin.Engine
}

// Run the application
func main() {
	configFile := flag.String("config", "config.yaml", "Config file for nsl backend")
	config, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not configure server, %v", err)
		return
	}
	app := NewApp(config)

	app.setupRoutes()
	app.Start()
}

func (app *App) setupRoutes() {
	app.server = gin.Default()

	/* POST endpoint that calls app's Login function defined in auth.go */
	app.server.POST("/login", app.Login)

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
}

func (app *App) Start() {
	app.server.Run(":" + strconv.Itoa(app.conf.Port))
}

// NewApp creates the app that holds all the functions to interact with the DB
func NewApp(config *ServerConfig) App {
	app := App{
		conf: *config,
		db:   &psql.DB{},
	}
	err := app.db.Init(config.DBConfig)
	if err != nil {
		log.Fatal("could not initialize database")
	}
	return app
}

// Read the config.yml file
func ReadConfig(file string) (*ServerConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Create a ServerConfig struct and read file info into the struct
	var config ServerConfig
	yaml.Unmarshal(data, &config)

	if config.Port == 0 {
		config.Port = 4000
	}
	return &config, nil
}
