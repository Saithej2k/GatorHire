package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/google/uuid"
)

// SaveJobRequest represents the request body for saving a job
type SaveJobRequest struct {
	JobID string `json:"jobId"`
}

// SaveJob saves a job for the logged-in user
func SaveJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to save a job...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Parse request body
	var req SaveJobRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("❌ Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if req.JobID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job ID is required"})
		return
	}

	// Check if the job exists
	var jobExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM jobs WHERE id = $1)", req.JobID).Scan(&jobExists)
	if err != nil {
		log.Println("❌ Error checking job existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to verify job"})
		return
	}
	if !jobExists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		return
	}

	// Check if the job is already saved by the user
	var alreadySaved bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM saved_jobs WHERE user_id = $1 AND job_id = $2)", userID, req.JobID).Scan(&alreadySaved)
	if err != nil {
		log.Println("❌ Error checking if job is already saved:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to check saved status"})
		return
	}
	if alreadySaved {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job already saved"})
		return
	}

	// Save the job
	savedJobID := uuid.New().String()
	_, err = db.DB.Exec(`
        INSERT INTO saved_jobs (id, user_id, job_id, saved_date)
        VALUES ($1, $2, $3, NOW())
    `, savedJobID, userID, req.JobID)
	if err != nil {
		log.Println("❌ Database error while saving job:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to save job"})
		return
	}

	log.Println("✅ Job saved successfully - ID:", savedJobID)
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}

// UnsaveJob removes a saved job for the logged-in user
func UnsaveJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to unsave a job...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Parse query parameter for job ID
	jobID := r.URL.Query().Get("jobId")
	if jobID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job ID is required"})
		return
	}

	// Check if the saved job exists
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM saved_jobs WHERE user_id = $1 AND job_id = $2)", userID, jobID).Scan(&exists)
	if err != nil {
		log.Println("❌ Error checking saved job existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to verify saved job"})
		return
	}
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Saved job not found"})
		return
	}

	// Delete the saved job
	_, err = db.DB.Exec("DELETE FROM saved_jobs WHERE user_id = $1 AND job_id = $2", userID, jobID)
	if err != nil {
		log.Println("❌ Database error while unsaving job:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to unsave job"})
		return
	}

	log.Println("✅ Job unsaved successfully - JobID:", jobID)
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}

// GetSavedJobs retrieves all saved jobs for the logged-in user
// GetSavedJobs retrieves all saved jobs for the logged-in user
func GetSavedJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to fetch saved jobs...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Query to fetch saved jobs with job details
	query := `
        SELECT s.id, s.user_id, s.job_id, s.saved_date,
               j.id, j.title, j.company, j.location, j.type, j.salary, j.description,
               j.requirements, j.responsibilities, j.benefits, j.posted_date,
               j.category, j.status, j.company_info, COALESCE(j.created_by, '')
        FROM saved_jobs s
        JOIN jobs j ON s.job_id = j.id
        WHERE s.user_id = $1 AND j.status = 'active'
        ORDER BY s.saved_date DESC
    `
	log.Println("📡 Executing SQL Query:", query)
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		log.Println("❌ Database query error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		return
	}
	defer rows.Close()

	// Parse results
	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		var savedJobID, userID, jobID string
		var savedDate time.Time
		var requirementsJSON, responsibilitiesJSON, benefitsJSON, companyInfoJSON []byte
		var createdBy string

		err := rows.Scan(
			&savedJobID, &userID, &jobID, &savedDate,
			&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
			&job.Salary, &job.Description, &requirementsJSON, &responsibilitiesJSON,
			&benefitsJSON, &job.PostedDate, &job.Category, &job.Status,
			&companyInfoJSON, &createdBy,
		)
		if err != nil {
			log.Println("❌ Error scanning saved job data:", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing saved job data"})
			return
		}

		// Convert JSONB fields to slices
		if len(requirementsJSON) > 0 {
			if err := json.Unmarshal(requirementsJSON, &job.Requirements); err != nil {
				job.Requirements = []string{}
			}
		}

		if len(responsibilitiesJSON) > 0 {
			if err := json.Unmarshal(responsibilitiesJSON, &job.Responsibilities); err != nil {
				job.Responsibilities = []string{}
			}
		}

		if len(benefitsJSON) > 0 {
			if err := json.Unmarshal(benefitsJSON, &job.Benefits); err != nil {
				job.Benefits = []string{}
			}
		}

		// Parse company info
		if len(companyInfoJSON) > 0 && string(companyInfoJSON) != "[]" {
			var companyInfo models.CompanyInfo
			if err := json.Unmarshal(companyInfoJSON, &companyInfo); err != nil {
				log.Printf("⚠️ Error unmarshalling company info: %v", err)
			} else {
				job.CompanyInfo = &companyInfo
			}
		}

		job.CreatedBy = createdBy
		jobs = append(jobs, job)
	}

	// Ensure an empty array is returned if no jobs are found
	if jobs == nil {
		jobs = []models.Job{}
	}

	log.Printf("📦 Returning %d saved jobs\n", len(jobs))
	json.NewEncoder(w).Encode(jobs)
}

// BulkDeleteSavedJobsRequest represents the request body for bulk deleting saved jobs
type BulkDeleteSavedJobsRequest struct {
	JobIDs []string `json:"jobIds"`
}

// BulkDeleteSavedJobs deletes multiple saved jobs for the logged-in user
func BulkDeleteSavedJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to bulk delete saved jobs...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Parse request body
	var req BulkDeleteSavedJobsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("❌ Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if len(req.JobIDs) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "No job IDs provided"})
		return
	}

	// Build query to delete saved jobs
	query := "DELETE FROM saved_jobs WHERE user_id = $1 AND job_id IN ("
	args := []interface{}{userID}
	placeholders := []string{}
	for i, jobID := range req.JobIDs {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+2))
		args = append(args, jobID)
	}
	query += strings.Join(placeholders, ",") + ")"

	// Execute deletion
	log.Println("📡 Executing SQL Query:", query)
	result, err := db.DB.Exec(query, args...)
	if err != nil {
		log.Println("❌ Database error while deleting saved jobs:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to delete saved jobs"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error getting rows affected:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to confirm deletion"})
		return
	}

	log.Printf("✅ Deleted %d saved jobs\n", rowsAffected)
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}
