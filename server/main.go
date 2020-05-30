package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port int `yaml:"port"`
}

func main() {
	configFile := flag.String("config", "config.yaml", "Config file for nsl backend")
	config, err := ReadConfig(*configFile)
	if err != nil {
		log.Fatalf("Could not configure server, %v", err)
		return
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":" + strconv.Itoa(config.Port))
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
