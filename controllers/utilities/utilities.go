package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response json struct
type Response struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
}

// UtilityIndexLHandler index pagego
func UtilityIndexLHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{StatusCode: "200", Message: "indexPage"}
	fmt.Println("utilities controller")
	json.NewEncoder(w).Encode(response)
}

// Add
func Add(a, b int) int {
	return a + b
}
