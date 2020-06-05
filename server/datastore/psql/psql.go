package psql

import (
	"fmt"
	"os"
	"github.com/jackc/pgx/v4"
)

type DataBase struct {
	host, password, dbname string
	port int32
}


