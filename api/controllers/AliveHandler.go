package controllers

import (
	"encoding/json"
	"net/http"
)

func AliveHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{\"Alive\":true}")
}
