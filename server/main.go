package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"fmt"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore/inmemory"
	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port int `yaml:"port"`
	Host string `yaml:"host"`
	User string `yaml: "user"`
	Password string `yaml: "password"`
	Dbname string `yaml: "dbname"`
}

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

	app := NewApp()

	fmt.Println(app)
	r := gin.Default()

	/* POST endpoint that calls app's Login function defined in auth.go */
	// r.POST("/login", app.Login)

	/* GET endpoint that calls app's ___ function defined in maintenance.go */
	  // <-- Implementation here. -->
	
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

// Creates the app that holds all the functions to interact with the DB
func NewApp() App {
	app := App{
		db: &inmemory.InMemory{},
	}
	err := app.db.Init()
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
