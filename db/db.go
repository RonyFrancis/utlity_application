package db

import (
	"database/sql"
	"fmt"

	// postgres package
	_ "github.com/lib/pq"
)

// ConnectToDB connect to database
func ConnectToDB(host, user, password, dbname, driver string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open(driver, psqlInfo)
	fmt.Println(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
