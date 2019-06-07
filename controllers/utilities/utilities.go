package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

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

// UtilityIndexHandler index page
func UtilityIndexHandler(w http.ResponseWriter, r *http.Request) {
	order := utilitiyOrder.CreateUtilityOrder()
	rec, err := order.All()
	if err != nil {
		fmt.Println("fetch failed")
	}
	json.NewEncoder(w).Encode(rec)
}

// UtilityEditHandler edit page
func UtilityEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i1, err := strconv.Atoi(vars["id"])
	if err == nil {
		fmt.Println(i1)
	}
	order := utilitiyOrder.CreateUtilityOrder()
	rec, err := order.GetByID(db.DBconnection(), i1)
	if err != nil {
		fmt.Println("fetch failed")
	}
	response := utilitiyOrder.Response(rec)
	lp := filepath.Join("templates", "layout.tmpl")
	fp := filepath.Join("templates", "edit.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "layout", response)
}

// UtilityUpdateHandler update page
func UtilityUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i1, err := strconv.Atoi(vars["id"])
	if err == nil {
		fmt.Println(i1)
	}
	r.ParseForm()
	order := utilitiyOrder.CreateUtilityOrder()
	err = order.UpdateRecord(i1, r.FormValue)
	if err != nil {
		fmt.Println(err)
	}
	redirectURL := fmt.Sprintf("/show/%d", i1)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// UtilityNewHandler new page
func UtilityNewHandler(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.tmpl")
	fp := filepath.Join("templates", "new.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}

// UtilityCreateHandler create new record
func UtilityCreateHandler(w http.ResponseWriter, r *http.Request) {
	order := utilitiyOrder.CreateUtilityOrder()
	r.ParseForm()
	order.TrxDate = r.FormValue("trxdate")
	order.TrxTime = r.FormValue("trxtime")
	i1, err := strconv.Atoi(r.FormValue("userid"))
	if err != nil {
		fmt.Println(err)
	}
	order.UserID = i1
	err = order.CreateRecord()
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
