package main

import (
	"fmt"

	"github.com/RonyFrancis/utlity_application/db"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rony"
	password = "rony"
	dbname   = "testdb"
	driver   = "postgres"
)

func main() {
	fmt.Println("utlity app")
	db, err := db.ConnectToDB(host, user, password, dbname, driver, port)
	if err != nil {
		fmt.Println("DB connection failed!")
	}
	defer db.Close()
	fmt.Println("DB connection successful!")
	// myRouter := mux.NewRouter()
	// myRouter.HandleFunc("/", utilities.UtilityIndexLHandler).Methods("POST")
	// http.ListenAndServe(":8001", myRouter)
}
