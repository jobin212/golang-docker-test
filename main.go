package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	router.HandleFunc("/events/{id:[0-9]+}", GetEvent).Methods("GET")
	router.HandleFunc("/events", CreateEvent).Methods("POST")
	router.HandleFunc("/", PingDB).Methods("GET")

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	// get mux vars
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	// first get by email, then write if empty
	sqlStatement := `SELECT age, first_name, last_name, email FROM jobin212.users WHERE id=$1;`

	row := db.QueryRow(sqlStatement, id)

	user_obj := user_model{ID: id}

	if err = row.Scan(&user_obj.Age, &user_obj.FirstName, &user_obj.LastName, &user_obj.Email); err != nil {
		switch err {
		case sql.ErrNoRows:
			http.Error(w, "There is no user associated with this id", http.StatusNotFound)
			return
		default:
			http.Error(w, "There is no user associated with this id", http.StatusInternalServerError)
		}
	}

	// fmt.Fprintf(w, "User exists %s, %s", user_obj.FirstName, user_obj.LastName)
	respondWithJSON(w, 200, user_obj)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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

type user_creation_request struct {
	Age       int
	Email     string
	FirstName string
	LastName  string
}

type user_model struct {
	Age       int    `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ID        int    `json:"id"`
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request user_creation_request
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	// first get by email, then write if empty
	sqlStatement := `SELECT id FROM jobin212.users WHERE email=$1;`
	var id int
	row := db.QueryRow(sqlStatement, request.Email)

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
	_, err = db.Exec(sqlStatement, request.Age, request.Email, request.FirstName, request.LastName)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "First Name: %s, Last Name: %s, Email: %s, Age: %d",
		request.FirstName, request.LastName, request.Email, request.Age)
}
