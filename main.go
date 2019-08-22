package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/lib/pq"
)

type Document struct {
	Metadata postgres.Jsonb
	Secrets  postgres.Hstore
	Body     string
	ID       int
}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "p@ssword1"
	dbname   = "testdb"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			fmt.Println("ping failed")
			panic(err)
		}

		fmt.Fprintf(w, "__")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			fmt.Println("ping failed")
			panic(err)
		}

		fmt.Fprintf(w, "ping succeeded")
	})

	fmt.Println("Listening on port 8080-- hello world!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
