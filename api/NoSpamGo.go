package main

import (
	"NoSpamGo/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"

	"github.com/julienschmidt/httprouter"
)

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Executing middlewareTwo")
		fmt.Printf("Method: %s, URI: %s, Origin : %s\n", r.Method, r.RequestURI, r.Header.Get("Origin"))
		next.ServeHTTP(w, r)
		//log.Println("Executing middlewareTwo again")
	})
}

func main() {

	router := httprouter.New()
	router.POST("/setup-2fa", controllers.Setup2FactorsHandler)
	router.POST("/verify-2fa", controllers.Verify2FactorsHandler)
	router.POST("/spam-detector", controllers.SpamDetectorHandler)
	router.GET("/alive", controllers.AliveHandler)

	middleware := middlewareTwo(router)

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
