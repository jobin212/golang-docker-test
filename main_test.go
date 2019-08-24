package main // todo should this be main_test?

import (
	"log"
	"os"
	"testing"
)

// TODO move consts to docker file
const (
	test_host     = "postgres"
	test_port     = 5432
	test_user     = "postgres"
	test_password = "p@ssword1"
	test_dbname   = "testdb"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(host, port, user, password, dbname)

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)	
)`
