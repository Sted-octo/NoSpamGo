package controllers

import (
	"NoSpamGo/domain"
	"encoding/json"
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

	user := new(domain.User)
	// Récupérer l'utilisateur
	/*user, exists := users[req.Username]
	if !exists {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}*/

	// Vérifier le code TOTP
	valid := totp.Validate(req.Token, user.Secret)

	response := struct {
		Valid bool `json:"valid"`
	}{
		Valid: valid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
