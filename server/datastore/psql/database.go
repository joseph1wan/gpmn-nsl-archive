package psql

import (
	"fmt"
	"os"
	"github.com/jackc/pgx"
)

/* Create a Struct that can connect with the database.*/
type DataBase struct {
	host, password, dbname string
	port int32
}
