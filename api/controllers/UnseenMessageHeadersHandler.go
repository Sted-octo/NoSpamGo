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

func UnseenMessageHeadersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var email string = ps.ByName("email")

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnector.Close()

	var userByMailLoader usecases.IUserByMailLoader[*sql.DB] = new(dataprovider.UserByMailLoader)

	user := userByMailLoader.Load(email, dbConnector)
	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("[]")
		return
	}

	var unseenMessagesGetter usecases.IUnseenMessagesGetter[*client.Client] = new(dataprovider.ImapClientUnseenMessagesGetter)
	var clientConnector usecases.IClientConnector[*client.Client] = new(dataprovider.ImapClientConnector)

	err = clientConnector.Connect(user.ImapServerUrl, user.ImapServerPort, user.ImapUsername, user.ImapPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnector.Close()

	messages := unseenMessagesGetter.Get(clientConnector)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
