package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
	"github.com/google/uuid"
)

// SaveJob saves a job for the authenticated user
func SaveJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is authenticated
	userID, _, err := utils.GetUserFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Parse request body
	var request struct {
		JobID string `json:"jobId"`
	}
	
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Check if job exists
	var jobExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM jobs WHERE id = $1)", request.JobID).Scan(&jobExists)
	if err != nil || !jobExists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		return
	}
	
	// Check if job is already saved
	var alreadySaved bool
	err = db.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM saved_jobs 
			WHERE job_id = $1 AND user_id = $2
		)
	`, request.JobID, userID).Scan(&alreadySaved)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		return
	}
	
	if alreadySaved {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job already saved"})
		return
	}
	
	// Save job
	savedJob := models.SavedJob{
		ID:        uuid.New().String(),
		UserID:    userID,
		JobID:     request.JobID,
		SavedDate: time.Now(),
	}
	
	_, err = db.DB.Exec(`
		INSERT INTO saved_jobs (id, user_id, job_id, saved_date)
		VALUES ($1, $2, $3, $4)
	`, savedJob.ID, savedJob.UserID, savedJob.JobID, savedJob.SavedDate)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to save job"})
		return
	}
	
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}

// UnsaveJob removes a saved job for the authenticated user
func UnsaveJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is authenticated
	userID, _, err := utils.GetUserFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Get job ID from URL
	jobID := r.URL.Query().Get("jobId")
	if jobID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job ID is required"})
		return
	}
	
	// Delete saved job
	result, err := db.DB.Exec(`
		DELETE FROM saved_jobs
		WHERE user_id = $1 AND job_id = $2
	`, userID, jobID)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to unsave job"})
		return
	}
	
	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Saved job not found"})
		return
	}
	
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}

// GetSavedJobs returns all saved jobs for the authenticated user
func GetSavedJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is authenticated
	userID, _, err := utils.GetUserFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Query database for saved jobs
	rows, err := db.DB.Query(`
		SELECT s.id, s.job_id, s.saved_date,
		       j.title, j.company, j.location, j.type, j.posted_date, j.category
		FROM saved_jobs s
		JOIN jobs j ON s.job_id = j.id
		WHERE s.user_id = $1
		ORDER BY s.saved_date DESC
	`, userID)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		return
	}
	defer rows.Close()
	
	// Parse results
	var savedJobs []map[string]interface{}
	for rows.Next() {
		var savedJob struct {
			ID        string
			JobID     string
			SavedDate time.Time
			Title     string
			Company   string
			Location  string
			Type      string
			PostedDate time.Time
			Category  string
		}
		
		err := rows.Scan(
			&savedJob.ID, &savedJob.JobID, &savedJob.SavedDate,
			&savedJob.Title, &savedJob.Company, &savedJob.Location,
			&savedJob.Type, &savedJob.PostedDate, &savedJob.Category,
		)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing saved job data"})
			return
		}
		
		// Format for response
		savedJobs = append(savedJobs, map[string]interface{}{
			"id":        savedJob.JobID,
			"title":     savedJob.Title,
			"company":   savedJob.Company,
			"location":  savedJob.Location,
			"type":      savedJob.Type,
			"postedDate": savedJob.PostedDate,
			"savedDate": savedJob.SavedDate,
			"category":  savedJob.Category,
		})
	}
	
	json.NewEncoder(w).Encode(savedJobs)
}