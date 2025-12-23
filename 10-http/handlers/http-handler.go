package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponseEntity struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Live    bool   `json:"live"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponseEntity{Live: true, Message: "Server online", Status: http.StatusOK}
	json.NewEncoder(w).Encode(response)
}
