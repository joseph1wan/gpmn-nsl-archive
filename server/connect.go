package main

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func main() {
	conn := Connect("connect test")
	defer conn.Close()
	fmt.Printf("Connection worked!\n")
}

func Connect(applicationName string) (conn *pgx.Conn) {
	var runtimeParams map[string]string
	runtimeParams = make(map[string]string)
	runtimeParams["application_name"] = applicationName
	connConfig := pgx.ConnConfig{
		User:              "postgres",
		Password:          "summer",
		Host:              "localhost",
		Port:              5432,
		Database:          "nsl_dev",
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
		RuntimeParams:     runtimeParams,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
		os.Exit(1)
	}
	return conn
}
