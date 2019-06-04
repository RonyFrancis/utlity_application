package db

import (
	"database/sql"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rony"
	password = "rony"
	dbname   = "testdb"
	driver   = "postgres"
)

func TestConnectToDB(t *testing.T) {
	db, err := ConnectToDB(host, user, password, dbname, driver, port)
	if err != nil {
		t.Errorf("DB connection failed")
	}
	defer CloseDB(db)
	db1, err := ConnectToDB("", "", "", "", "", 0)
	if err != nil {
		t.Errorf("DB connection failed")
	}
	defer CloseDB(db1)
}

func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
