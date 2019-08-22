package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "p@ssword1"
	dbname   = "testdb"
)

var db *sql.DB

func main() {
	router := mux.NewRouter()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router.HandleFunc("/ping", PingDB).Methods("GET")
	router.HandleFunc("/events", GetEvents).Methods("GET")
	router.HandleFunc("/", PingDB).Methods("GET")

	fmt.Println("Listening on port 8080-- hello world!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func PingDB(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
		panic(err)
	}

	fmt.Fprintf(w, "ping succeeded")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
		panic(err)
	}

	fmt.Fprintf(w, "events ping succeeded")
}
