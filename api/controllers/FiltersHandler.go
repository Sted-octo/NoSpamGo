package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var email string = ps.ByName("email")

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnector.Close()

	var filtersGetter usecases.IFiltersGetter[*sql.DB] = new(dataprovider.FiltersGetter)

	filters := filtersGetter.Get(email, dbConnector)
	if filters == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("[]")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filters)
}
