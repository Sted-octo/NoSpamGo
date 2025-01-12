package controllers

import (
	dataprovider "NoSpamGo/dataProvider"
	"NoSpamGo/tools"
	"NoSpamGo/usecases"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	isMailConfigOk := !(user.ImapUsername == "" && user.ImapPassword == "" && user.ImapServerUrl == "" && user.ImapServerPort == 0)

	response := struct {
		Valid          bool   `json:"valid"`
		Token          string `json:"token,omitempty"`
		IsMailConfigOk bool   `json:"ismailconfigok"`
	}{
		Valid:          valid,
		Token:          "",
		IsMailConfigOk: isMailConfigOk,
	}

	if valid {
		jwtKey := os.Getenv("JWT_KEY")

		token, err := tools.GenerateToken(user.Mail, jwtKey)
		if err != nil {
			log.Printf("Error generating jwt token : %s \r\n", err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		response.Token = token
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
