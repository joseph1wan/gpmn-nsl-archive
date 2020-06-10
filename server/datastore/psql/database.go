package psql

import (
	"github.com/jackc/pgx"
	"io/ioutil"
	"flag"
	"log"
	yaml "gopkg.in/yaml.v2"
)
type InitConfig struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
	User string `yaml: "user"`
	Password string `yaml: "password"`
	Dbname string `yaml: "dbname"`
}
/* Create a Struct that implements datastore.DB*/
type DB struct {
	conn *pgx.Conn
}

/* Init establishes a reusable connection */
func (db *DB) Init() (error) {
	// The hard-coded params should be replaced by the config.yml file
	// You'll need to update the interface to accept a param also
	initConfigFile := flag.String("config", "config.yaml", "Config file for nsl backend")
	initconfig, err := ReadConfig(*initConfigFile)
	if err != nil {
		log.Fatalf("Could not configure server, %v", err)
		return err
	}
	connConfig := pgx.ConnConfig{
		User:              initconfig.User,
		Password:          initconfig.Password,
		Host:              initconfig.Host,
		Port:              initconfig.Port,
		Database:          initconfig.Dbname,
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

/* Connection fetches an establish pgx connection */
func (db *DB) Connection() (*pgx.Conn) {
	return db.conn
}

func ReadConfig(file string) (*InitConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// Create a ServerConfig struct and read file info into the struct
	var initConfig InitConfig
	yaml.Unmarshal(data, &initConfig)

	if initConfig.Port == 0 {
		initConfig.Port = 4000
	}
	return &initConfig, nil
}
