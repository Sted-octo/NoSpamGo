package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AliveHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("AliveHandler")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{\"Alive\":true}")
}
