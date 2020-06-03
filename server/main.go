package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/a2fumn2022/gpmn-nsl/server/datastore/inmemory"
	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

type App struct {
	db datastore.DB
}

func main() {
	configFile := flag.String("config", "config.yaml", "Config file for nsl backend")
	config, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not configure server, %v", err)
		return
	}

	app := NewApp()

	r := gin.Default()

	r.POST("/login", app.Login)

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

func ReadConfig(file string) (*ServerConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	yaml.Unmarshal(data, &config)

	if config.Port == 0 {
		config.Port = 4000
	}
	return &config, nil
}
