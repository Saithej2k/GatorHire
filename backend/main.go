package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	db.InitDB()
	defer db.CloseDB()

	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// Jobs endpoints
	api.HandleFunc("/jobs", handlers.GetJobs).Methods("GET", "OPTIONS")
	api.HandleFunc("/jobs/{id}", handlers.GetJobByID).Methods("GET", "OPTIONS")
	api.HandleFunc("/jobs", handlers.CreateJob).Methods("POST", "OPTIONS")
	api.HandleFunc("/jobs/{id}", handlers.UpdateJob).Methods("PUT", "OPTIONS")
	api.HandleFunc("/jobs/{id}", handlers.DeleteJob).Methods("DELETE", "OPTIONS")
	
	// Applications endpoints
	api.HandleFunc("/applications", handlers.CreateApplication).Methods("POST", "OPTIONS")
	api.HandleFunc("/applications/user", handlers.GetUserApplications).Methods("GET", "OPTIONS")
	api.HandleFunc("/applications/job", handlers.GetApplicationsByJob).Methods("GET", "OPTIONS")
	api.HandleFunc("/applications/status", handlers.UpdateApplicationStatus).Methods("PUT", "OPTIONS")
	
	// Saved jobs endpoints
	api.HandleFunc("/saved-jobs", handlers.SaveJob).Methods("POST", "OPTIONS")
	api.HandleFunc("/saved-jobs", handlers.UnsaveJob).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/saved-jobs", handlers.GetSavedJobs).Methods("GET", "OPTIONS")
	
	// Auth endpoints
	api.HandleFunc("/auth/login", handlers.Login).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/register", handlers.Register).Methods("POST", "OPTIONS")
	
	// Profile endpoints
	api.HandleFunc("/profile", handlers.GetProfile).Methods("GET", "OPTIONS")
	api.HandleFunc("/profile", handlers.UpdateProfile).Methods("PUT", "OPTIONS")

	// Set up CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	// Start server with CORS middleware
	handler := corsMiddleware.Handler(r)
	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}