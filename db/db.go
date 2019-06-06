package db

import (
	"database/sql"
	"fmt"

	// postgres package
	_ "github.com/lib/pq"
)

// DBconnect connection pointer
var DBconnect *sql.DB

// ConnectToDB connect to database
func ConnectToDB(host, user, password, dbname, driver string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DBconnect, err = sql.Open(driver, psqlInfo)
	fmt.Println(DBconnect)
	if err != nil {
		return nil, err
	}
	return DBconnect, nil
}

// DBconnection return the connection
func DBconnection() *sql.DB {
	return DBconnect
}
