package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pquerna/otp/totp"
)

func Setup2FactorsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var dbConnector usecases.IDatabaseConnector[*sql.DB] = new(dataprovider.DatabaseConnector)
	err := dbConnector.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnector.Close()

	var userSaver usecases.IUserSaver[*sql.DB] = new(dataprovider.UserSaver)
	var userByMailLoader usecases.IUserByMailLoader[*sql.DB] = new(dataprovider.UserByMailLoader)

	setup2Factors(w, r, dbConnector, userSaver, userByMailLoader)

}

func setup2Factors(w http.ResponseWriter, r *http.Request, dbConnector usecases.IDatabaseConnector[*sql.DB], userSaver usecases.IUserSaver[*sql.DB], userByMailLoader usecases.IUserByMailLoader[*sql.DB]) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NoSpam",
		AccountName: user.Mail,
	})
	if err != nil {
		http.Error(w, "Erreur lors de la génération du secret", http.StatusInternalServerError)
		return
	}

	userSaver.Save(domain.User{
		Mail:           user.Mail,
		Secret:         key.Secret(),
		ImapUsername:   "",
		ImapPassword:   "",
		ImapServerUrl:  "",
		ImapServerPort: 0,
	}, dbConnector, userByMailLoader)

	response := struct {
		Secret string `json:"secret"`
		QRCode string `json:"qr_code"`
	}{
		Secret: key.Secret(),
		QRCode: key.URL(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
