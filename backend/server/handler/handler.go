package handler

import (
	"backend/server/cors"
	"log"
	"net/http"
	"os"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}

func AuthorizeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.SetCors(&w)
		next.ServeHTTP(w, r)
	}
}

func RequestValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.SetCors(&w)
		next.ServeHTTP(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// Ouverture du fichier de journalisation en mode append
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Création d'un logger avec le fichier ouvert
	logger := log.New(file, "", log.LstdFlags)
	return func(w http.ResponseWriter, r *http.Request) {
		// Log des informations sur la requête entrante
		startTime := time.Now()
		logger.Printf("Incoming request: %s %s %s", r.Method, r.URL.Path, r.Proto)
		logger.Printf("Headers: %v", r.Header)

		// Création d'un responseWriter personnalisé pour capturer le code de statut
		rl := &responseLogger{ResponseWriter: w}

		// Appel du handler suivant dans la chaîne
		next.ServeHTTP(rl, r)

		// Log des informations sur la réponse sortante
		elapsedTime := time.Since(startTime)
		logger.Printf("Outgoing response: Status %d, took %s", rl.status, elapsedTime)
	}
}

// responseLogger est un http.ResponseWriter personnalisé qui capture le code de statut
type responseLogger struct {
	http.ResponseWriter
	status int
}

// WriteHeader capture le code de statut avant d'écrire l'en-tête
func (rl *responseLogger) WriteHeader(status int) {
	rl.status = status
	rl.ResponseWriter.WriteHeader(status)
}
