package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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
	router.HandleFunc("/createEvent", CreateEvent).Methods("GET")
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

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	t1 := r1.Intn(100)

	sqlStatement := `
	INSERT INTO jobin212.users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, t1, "xxyy@email.com", "Joe", "Smith")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Jos Smith of age %d inserted into database", t1)
}
