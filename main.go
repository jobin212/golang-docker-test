package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	a := App{}

	// get environmental variables
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	a.Initialize(host, port, user, password, dbname)

	log.Println("running on port 8080...")
	a.Run(":8080")
}
