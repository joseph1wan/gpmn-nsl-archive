package main

import(
	"fmt"
	"database/sql"
	//"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "summer2020"
	dbname = "nsl_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port =%d user=%s " +
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}