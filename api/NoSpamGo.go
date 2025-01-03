package main

import (
	"NoSpamGo/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/setup-2fa", controllers.Setup2FactorsHandler).Methods("POST")
	router.HandleFunc("/verify-2fa", controllers.Verify2FactorsHandler).Methods("POST")
	router.HandleFunc("/spam-detector", controllers.SpamDetectorHandler).Methods("POST")
	router.HandleFunc("/alive", controllers.Setup2FactorsHandler).Methods("GET")

	fmt.Println("Serveur démarré sur :8080")
	http.ListenAndServe(":8080", router)
}
