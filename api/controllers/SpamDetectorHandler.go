package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/emersion/go-imap/client"
	"github.com/julienschmidt/httprouter"
)

func SpamDetectorHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var email dataprovider.Email
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnector.Close()

	var userByMailLoader usecases.IUserByMailLoader[*sql.DB] = new(dataprovider.UserByMailLoader)
	var unseenMessagesGetter usecases.IUnseenMessagesGetter[*client.Client] = new(dataprovider.ImapClientUnseenMessagesGetter)
	var spamMover usecases.ISpamMover[*client.Client] = new(dataprovider.ImapClientSpamMover)
	var clientConnector usecases.IClientConnector[*client.Client] = new(dataprovider.ImapClientConnector)
	var filtersGetter usecases.IFiltersGetter[*sql.DB] = new(dataprovider.FiltersGetter)
	var filterByNameForUserMailLoader usecases.IFilterByNameForUserMailLoader[*sql.DB] = new(dataprovider.FilterByNameForUserMailLoader)
	var filterSaver usecases.IFilterSaver[*sql.DB] = new(dataprovider.FilterSaver)

	saved := false
	if email.Mail != "" {
		user := userByMailLoader.Load(email.Mail, dbConnector)
		if user != nil {

			err = clientConnector.Connect(user.ImapServerUrl, user.ImapServerPort, user.ImapUsername, user.ImapPassword)
			if err != nil {
				log.Fatal(err)
			}
			defer clientConnector.Close()

			usecases.SpamDetector[*client.Client, *sql.DB](email.Mail, clientConnector, unseenMessagesGetter, spamMover, filtersGetter, dbConnector, filterSaver, filterByNameForUserMailLoader)
			saved = true
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
