package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	router.HandleFunc("/events", CreateEventFromBody).Methods("POST")
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
	_, err := db.Exec(sqlStatement, t1, "xxyy"+strconv.Itoa(t1)+"@email.com", "Joe", "Smith")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Jos Smith of age %d inserted into database", t1)
}

type user_creation_request struct {
	Age       int
	Email     string
	FirstName string
	LastName  string
}

func CreateEventFromBody(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user user_creation_request
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// first get by email, then write if empty
	sqlStatement := `SELECT id FROM jobin212.users WHERE email=$1;`
	var id int
	row := db.QueryRow(sqlStatement, user.Email)

	switch err = row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned, creating new user...")
	case nil:
		fmt.Println("Email already exists")
		http.Error(w, "Email aleady in use", http.StatusMethodNotAllowed)
		return
	default:
		panic(err)
	}

	// write if email does not exist
	sqlStatement = `
	INSERT INTO jobin212.users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, user.Age, user.Email, user.FirstName, user.LastName)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "First Name: %s, Last Name: %s, Email: %s, Age: %d", user.FirstName, user.LastName, user.Email, user.Age)
}
