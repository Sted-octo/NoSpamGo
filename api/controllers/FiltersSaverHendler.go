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

func FiltersSaverHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var userFilter dataprovider.UserFilters
	if err := json.NewDecoder(r.Body).Decode(&userFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnector.Close()

	var filterByNameForUserMailLoader usecases.IFilterByNameForUserMailLoader[*sql.DB] = new(dataprovider.FilterByNameForUserMailLoader)
	var filterSaver usecases.IFilterSaver[*sql.DB] = new(dataprovider.FilterSaver)

	saved := true

	for _, filter := range userFilter.Filters {
		state := filterSaver.Save(userFilter.Mail, filter, dbConnector, filterByNameForUserMailLoader)
		if !state {
			saved = false
			break
		}
	}

	response := struct {
		Saved bool `json:"saved"`
	}{
		Saved: saved,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
