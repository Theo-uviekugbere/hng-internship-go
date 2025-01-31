package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response structure
type Response struct {
	Email           string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GitHubURL       string `json:"github_url"`
}

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Handler function
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Email:           "uviekugbere.theo@gmail.com",
		CurrentDatetime: time.Now().UTC().Format(time.RFC3339),
		GitHubURL:       "https://github.com/Theo-uviekugbere/hng-internship-go",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", enableCORS(http.HandlerFunc(handler)))

	http.ListenAndServe(":8080", mux)
}
