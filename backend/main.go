package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/handlers"
	"github.com/gatorhire/backend/middleware"
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

	// Public routes (no authentication required)
	api.HandleFunc("/jobs", handlers.GetJobs).Methods("GET", "OPTIONS")
	api.HandleFunc("/jobs/{id}", handlers.GetJobByID).Methods("GET", "OPTIONS")
	api.HandleFunc("/jobs/search", handlers.SearchJobs).Methods("GET", "OPTIONS") // New endpoint
	api.HandleFunc("/auth/login", handlers.Login).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/register", handlers.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/applications", handlers.CreateApplication).Methods("POST", "OPTIONS")

	// Authenticated routes (require JWT token)
	authAPI := api.PathPrefix("").Subrouter()
	authAPI.Use(middleware.AuthMiddleware)

	authAPI.HandleFunc("/applications/user", handlers.GetUserApplications).Methods("GET", "OPTIONS")
	authAPI.HandleFunc("/saved-jobs", handlers.SaveJob).Methods("POST", "OPTIONS")
	authAPI.HandleFunc("/saved-jobs", handlers.UnsaveJob).Methods("DELETE", "OPTIONS")
	authAPI.HandleFunc("/saved-jobs/bulk", handlers.BulkDeleteSavedJobs).Methods("DELETE", "OPTIONS") // New endpoint
	authAPI.HandleFunc("/saved-jobs", handlers.GetSavedJobs).Methods("GET", "OPTIONS")
	authAPI.HandleFunc("/profile", handlers.GetProfile).Methods("GET", "OPTIONS")
	authAPI.HandleFunc("/profile", handlers.UpdateProfile).Methods("PUT", "OPTIONS")
	authAPI.HandleFunc("/profile/stats", handlers.GetProfileStats).Methods("GET", "OPTIONS")              // New endpoint
	authAPI.HandleFunc("/jobs/recommendations", handlers.GetJobRecommendations).Methods("GET", "OPTIONS") // New endpoint

	// Admin routes (require JWT token and admin role)
	adminAPI := api.PathPrefix("").Subrouter()
	adminAPI.Use(middleware.AuthMiddleware, middleware.AdminMiddleware)

	adminAPI.HandleFunc("/jobs", handlers.CreateJob).Methods("POST", "OPTIONS")
	adminAPI.HandleFunc("/jobs/{id}", handlers.UpdateJob).Methods("PUT", "OPTIONS")
	adminAPI.HandleFunc("/jobs/{id}", handlers.DeleteJob).Methods("DELETE", "OPTIONS")
	adminAPI.HandleFunc("/applications/job", handlers.GetApplicationsByJob).Methods("GET", "OPTIONS")
	adminAPI.HandleFunc("/applications/status", handlers.UpdateApplicationStatus).Methods("PUT", "OPTIONS")

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
