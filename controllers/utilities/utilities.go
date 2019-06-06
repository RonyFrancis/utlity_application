package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RonyFrancis/utlity_application/db"
	utilitiyOrder "github.com/RonyFrancis/utlity_application/models/utility_order"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rony"
	password = "rony"
	dbname   = "testdb"
	driver   = "postgres"
)

// Response json struct
type Response struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

// UtilityIndexLHandler index pagego
func UtilityIndexLHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{StatusCode: "200", Message: "indexPage"}
	json.NewEncoder(w).Encode(response)
}

// UtilityShowHandler return utility by ID
func UtilityShowHandler(w http.ResponseWriter, r *http.Request) {
	// response := Response{StatusCode: "200", Message: "indexPage"}
	vars := mux.Vars(r)
	order := utilitiyOrder.CreateUtilityOrder()
	dbConnect := db.DBconnection()
	i1, err := strconv.Atoi(vars["id"])
	if err == nil {
		fmt.Println(i1)
	}
	rec, err := order.GetByID(dbConnect, i1)
	if err != nil {
		fmt.Println("fetch failed")
	}
	response := utilitiyOrder.Response(rec)
	json.NewEncoder(w).Encode(response)
}
