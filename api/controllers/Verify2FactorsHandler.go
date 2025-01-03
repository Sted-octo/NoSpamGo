package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pquerna/otp/totp"
)

func Verify2FactorsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req Verify2FARequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	verify2Factors(w, req, dbConnector, userByMailLoader)

}

func verify2Factors(w http.ResponseWriter,
	req Verify2FARequest,
	dbConnector usecases.IDatabaseConnector[*sql.DB],
	userByMailLoader usecases.IUserByMailLoader[*sql.DB]) {

	user := userByMailLoader.Load(req.Mail, dbConnector)

	if user == nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	valid := totp.Validate(req.Token, user.Secret)

	response := struct {
		Valid bool `json:"valid"`
	}{
		Valid: valid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
