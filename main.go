package main

import (
	"fmt"
	"net/http"

	"github.com/RonyFrancis/utlity_application/controllers/utilities"
	"github.com/RonyFrancis/utlity_application/db"
	"github.com/gorilla/mux"
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
	dbConnect, err := db.ConnectToDB(host, user, password, dbname, driver, port)
	if err != nil {
		fmt.Println("DB connection failed!")
	}
	defer dbConnect.Close()
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/show/{id}", utilities.UtilityShowHandler).Methods("GET")
	myRouter.HandleFunc("/edit/{id}", utilities.UtilityEditHandler).Methods("GET")
	myRouter.HandleFunc("/index", utilities.UtilityIndexHandler).Methods("GET")
	myRouter.HandleFunc("/update/{id}", utilities.UtilityUpdateHandler).Methods("POST")
	myRouter.HandleFunc("/new", utilities.UtilityNewHandler).Methods("GET")
	myRouter.HandleFunc("/create", utilities.UtilityCreateHandler).Methods("POST")
	http.ListenAndServe(":8001", myRouter)
}
