package psql

import (
	"fmt"
	"os"
	"github.com/jackc/pgx"
)

type DataBase struct {
	host, password, dbname string
	port int32
}
