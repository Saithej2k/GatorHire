package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"gatorhire/internal/db"
	"gatorhire/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize database connection
	if err := db.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Auth routes
	api.HandleFunc("/auth/signup", handlers.SignUp).Methods("POST")
	api.HandleFunc("/auth/login", handlers.Login).Methods("POST")

	// Job routes
	api.HandleFunc("/jobs", handlers.CreateJob).Methods("POST")
	api.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	api.HandleFunc("/jobs/{id}", handlers.GetJob).Methods("GET")
	api.HandleFunc("/jobs/search", handlers.SearchJobs).Methods("GET")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}

	handler := c.Handler(r)
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
