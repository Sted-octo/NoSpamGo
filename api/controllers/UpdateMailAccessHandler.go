package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/domain"
	"NoSpamGo/tools"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func UpdateMailAccessHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dbConnector.Close()

	var userByMailLoader usecases.IUserByMailLoader[*sql.DB] = new(dataprovider.UserByMailLoader)
	var userSaver usecases.IUserSaver[*sql.DB] = new(dataprovider.UserSaver)
	var cryptoHelper usecases.ICryptoHelper = tools.NewCryptoHelper([]byte(os.Getenv("CRYPTO_KEY")))

	saved := userSaver.Save(user, dbConnector, userByMailLoader, cryptoHelper)

	response := struct {
		Saved bool `json:"saved"`
	}{
		Saved: saved,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
