package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// TODO move consts to docker file
const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "p@ssword1"
	dbname   = "testdb"
)

var db *sql.DB

func main() {
	a := App{}
	a.Initialize(host, port, user, password, dbname)

	a.Run(":8080")
}
