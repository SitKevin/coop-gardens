package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	host     = os.Getenv("DB_HOST")
)

func main() {
	// convert port to integer
	portInt, err := strconv.Atoi(port)
	CheckError(err)

	// connect to the database
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close the database connection
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected to the database!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
