package main

import (
	"NoSpamGo/controllers"
	"NoSpamGo/tools"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"

	"github.com/julienschmidt/httprouter"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.RequestURI, "/setup-2fa") ||
			strings.HasPrefix(r.RequestURI, "/verify-2fa") ||
			strings.HasPrefix(r.RequestURI, "/alive") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, tools.Concat(http.StatusText(http.StatusBadRequest), "Token string is empty"), http.StatusBadRequest)
			return
		}
		jwtKey := os.Getenv("JWT_KEY")

		claims, err := tools.IsTokenValide(tokenString, jwtKey)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		if claims == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		context.Set(r, "claimsID", claims.Email)
		next.ServeHTTP(w, r)
	})
}

func main() {

	router := httprouter.New()
	router.POST("/setup-2fa", controllers.Setup2FactorsHandler)
	router.POST("/verify-2fa", controllers.Verify2FactorsHandler)
	router.POST("/spam-detector", controllers.SpamDetectorHandler)
	router.GET("/alive", controllers.AliveHandler)
	router.POST("/update-mail-access", controllers.UpdateMailAccessHandler)

	middleware := middlewareOne(router)

	var origins []string
	origin, ok := os.LookupEnv("NOSPAM_ORIGINS")
	if ok {
		origins = strings.Split(origin, ",")
	}

	fmt.Println("Serveur démarré sur :8070")
	err := http.ListenAndServe(":8070",
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins(origins))(middleware))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
