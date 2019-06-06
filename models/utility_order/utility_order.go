package utilitiyOrder

import (
	"database/sql"
	"errors"
	"fmt"
)

// UtilityOrder model
type UtilityOrder struct {
	ID      int    `json:"id"`
	TrxDate string `json:"trxdate"`
	TrxTime string `json:"trxtime"`
	UserID  int    `json:"userid"`
}

// GetCount returns no of records
func (u *UtilityOrder) GetCount(db *sql.DB) int {
	rows, err := db.Query("SELECT COUNT(*) as count FROM  utilityorder")
	checkErr(err)
	return checkCount(rows)
}

// InsertRecord insert to table
func (u *UtilityOrder) InsertRecord(db *sql.DB) int {
	psqlInfo := fmt.Sprintf("INSERT INTO utilityorder VALUES (nextval('utilityorder_id_seq'), %s,%s,%d)", u.TrxDate, u.TrxTime, u.UserID)
	insert, err := db.Query(psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	fmt.Println(insert)
	return 10
}

// GetByID fetch record by ID
func (u *UtilityOrder) GetByID(db *sql.DB, id int) ([]*UtilityOrder, error) {
	sqlQuery := fmt.Sprintf("select * from utilityorder where id = %d ", id)
	insert, err := db.Query(sqlQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	payload := make([]*UtilityOrder, 0)
	for insert.Next() {
		data := new(UtilityOrder)
		err := insert.Scan(
			&data.ID,
			&data.TrxDate,
			&data.TrxTime,
			&data.UserID,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	if len(payload) == 0 {
		return nil, errors.New("No records found")
	}
	return payload, nil
}

// CreateUtilityOrder create empty struct
func CreateUtilityOrder() *UtilityOrder {
	return &UtilityOrder{
		TrxDate: "21",
		TrxTime: "32",
		UserID:  1,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Response(rec []*UtilityOrder) *UtilityOrder {
	if len(rec) > 0 {
		return rec[0]
	}
	return &UtilityOrder{}
}
