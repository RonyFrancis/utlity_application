package main

import (
	"fmt"
	"net/http"

	"github.com/RonyFrancis/utlity_application/controllers/utilities"
	"github.com/RonyFrancis/utlity_application/db"
	utilitiyOrder "github.com/RonyFrancis/utlity_application/models/utility_order"
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
	fmt.Println("DB connection successful!")
	order := utilitiyOrder.CreateUtilityOrder()
	fmt.Println(order.GetCount(dbConnect))
	// fmt.Println(order.InsertRecord(dbConnect))
	// rec, err := order.GetByID(dbConnect, 10000)
	// if err != nil {
	// 	fmt.Println("fetch failed")
	// }
	// fmt.Println(rec)
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", utilities.UtilityIndexLHandler).Methods("POST")
	myRouter.HandleFunc("/show/{id}", utilities.UtilityShowHandler).Methods("GET")
	http.ListenAndServe(":8001", myRouter)
}
