package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB_HOST     = os.Getenv("POSTGRES_PORT_5432_TCP_ADDR")
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "snsrobot"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbinfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME,
	)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	InitDB(db)

	defer db.Close()
}
