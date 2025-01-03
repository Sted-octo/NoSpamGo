package controllers

import (
	"NoSpamGo/domain"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pquerna/otp/totp"
)

func Setup2FactorsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Générer un nouveau secret pour l'utilisateur
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NoSpam",
		AccountName: user.Username,
	})
	if err != nil {
		http.Error(w, "Erreur lors de la génération du secret", http.StatusInternalServerError)
		return
	}

	// Sauvegarder l'utilisateur avec son secret
	/*users[user.Username] = &domain.User{
		Username:     user.Username,
		Secret:       key.Secret(),
		IsEnabled2FA: true,
	}*/

	// Préparer la réponse
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
