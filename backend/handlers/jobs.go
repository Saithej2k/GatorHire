package handlers

import (
	"encoding/json"
	"net/http"
	// "strconv"
	"strings"
	"log"
	"database/sql"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/lib/pq"
)

// GetJobs returns all job listings
// func GetJobs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	// Parse query parameters for filtering
// 	category := r.URL.Query().Get("category")
// 	searchTerm := r.URL.Query().Get("searchTerm")
// 	jobType := r.URL.Query().Get("jobType")
// 	location := r.URL.Query().Get("location")
	
// 	// Build query
// 	query := `
// 		SELECT id, title, company, location, type, salary, description, 
// 		       requirements, responsibilities, benefits, posted_date, 
// 		       category, status, company_info, created_by
// 		FROM jobs
// 		WHERE status = 'active'
// 	`
	
// 	// Add filters
// 	args := []interface{}{}
// 	argCount := 1
	
// 	if category != "" && category != "All" {
// 		query += " AND category = $" + strconv.Itoa(argCount)
// 		args = append(args, category)
// 		argCount++
// 	}
	
// 	if searchTerm != "" {
// 		query += " AND (title ILIKE $" + strconv.Itoa(argCount) + 
// 			" OR company ILIKE $" + strconv.Itoa(argCount) + 
// 			" OR location ILIKE $" + strconv.Itoa(argCount) + ")"
// 		args = append(args, "%"+searchTerm+"%")
// 		argCount++
// 	}
	
// 	if jobType != "" {
// 		query += " AND type = $" + strconv.Itoa(argCount)
// 		args = append(args, jobType)
// 		argCount++
// 	}
	
// 	if location != "" {
// 		query += " AND location ILIKE $" + strconv.Itoa(argCount)
// 		args = append(args, "%"+location+"%")
// 		argCount++
// 	}
	
// 	// Order by posted date
// 	query += " ORDER BY posted_date DESC"
	
// 	// Execute query
// 	rows, err := db.DB.Query(query, args...)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
// 		return
// 	}
// 	defer rows.Close()
	
// 	// Parse results
// 	var jobs []models.Job
// 	for rows.Next() {
// 		var job models.Job
// 		var companyInfoBytes []byte
		
// 		err := rows.Scan(
// 			&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
// 			&job.Salary, &job.Description, &job.Requirements, &job.Responsibilities,
// 			&job.Benefits, &job.PostedDate, &job.Category, &job.Status,
// 			&companyInfoBytes, &job.CreatedBy,
// 		)
		
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing job data"})
// 			return
// 		}
		
// 		// Parse company info if present
// 		if len(companyInfoBytes) > 0 {
// 			var companyInfo models.CompanyInfo
// 			if err := json.Unmarshal(companyInfoBytes, &companyInfo); err == nil {
// 				job.CompanyInfo = &companyInfo
// 			}
// 		}
		
// 		jobs = append(jobs, job)
// 	}
	
// 	json.NewEncoder(w).Encode(jobs)


// GetJobs retrieves job listings based on filters
func GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to fetch jobs...")

	// Parse query parameters for filtering
	category := r.URL.Query().Get("category")
	searchTerm := r.URL.Query().Get("searchTerm")
	jobType := r.URL.Query().Get("jobType")
	location := r.URL.Query().Get("location")

	log.Printf("🔍 Filters - Category: %s | Search: %s | Type: %s | Location: %s\n", category, searchTerm, jobType, location)

	// Build query
	query := `
		SELECT id, title, company, location, type, salary, description, 
		       requirements, responsibilities, benefits, posted_date, 
		       category, status, company_info, COALESCE(created_by, '') -- Fix NULL issue
		FROM jobs
		WHERE status = 'active'
		ORDER BY posted_date DESC
	`

	// Execute query
	log.Println("📡 Executing SQL Query:", query)
	rows, err := db.DB.Query(query)
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
		var requirementsJSON, responsibilitiesJSON, benefitsJSON, companyInfoJSON []byte
		var createdBy sql.NullString // Handle NULL values for created_by

		err := rows.Scan(
			&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
			&job.Salary, &job.Description, &requirementsJSON, &responsibilitiesJSON,
			&benefitsJSON, &job.PostedDate, &job.Category, &job.Status,
			&companyInfoJSON, &createdBy,
		)

		if err != nil {
			log.Println("❌ Error scanning job data:", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing job data"})
			return
		}

		// Convert JSONB fields to slices
		if len(requirementsJSON) > 0 {
			if err := json.Unmarshal(requirementsJSON, &job.Requirements); err != nil {
				job.Requirements = []string{} // Default to empty slice
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
		if len(companyInfoJSON) > 0 {
			var companyInfo models.CompanyInfo
			if err := json.Unmarshal(companyInfoJSON, &companyInfo); err != nil {
			} else {
				job.CompanyInfo = &companyInfo
			}
		}

		// ✅ Fix NULL issue for created_by
		job.CreatedBy = createdBy.String // If NULL, it will default to ""

		jobs = append(jobs, job)
	}

	// Return response
	log.Printf("📦 Returning %d jobs\n", len(jobs))
	json.NewEncoder(w).Encode(jobs)
}



// GetJobByID returns a specific job by ID
// func GetJobByID(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	// Get job ID from URL
// 	params := mux.Vars(r)
// 	id := params["id"]
	
// 	// Query database
// 	var job models.Job
// 	var companyInfoBytes []byte
	
// 	err := db.DB.QueryRow(`
// 		SELECT id, title, company, location, type, salary, description, 
// 		       requirements, responsibilities, benefits, posted_date, 
// 		       category, status, company_info, created_by
// 		FROM jobs
// 		WHERE id = $1
// 	`, id).Scan(
// 		&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
// 		&job.Salary, &job.Description, &job.Requirements, &job.Responsibilities,
// 		&job.Benefits, &job.PostedDate, &job.Category, &job.Status,
// 		&companyInfoBytes, &job.CreatedBy,
// 	)
	
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
// 		return
// 	}
	
// 	// Parse company info if present
// 	if len(companyInfoBytes) > 0 {
// 		var companyInfo models.CompanyInfo
// 		if err := json.Unmarshal(companyInfoBytes, &companyInfo); err == nil {
// 			job.CompanyInfo = &companyInfo
// 		}
// 	}
	
// 	json.NewEncoder(w).Encode(job)
// }
func GetJobByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get job ID from URL
	params := mux.Vars(r)
	id := params["id"]
	
	// Log received ID for debugging
	log.Printf("🔍 Received request to fetch job with ID: %s", id)
	
	// Query database to fetch the job details
	var job models.Job
	var companyInfoBytes []byte
	var requirementsBytes []byte
	var responsibilitiesBytes []byte
	var benefitsBytes []byte
	var createdBy sql.NullString // Use sql.NullString for nullable field
	
	// Log the query being executed for debugging
	query := `
		SELECT id, title, company, location, type, salary, description, 
		       requirements, responsibilities, benefits, posted_date, 
		       category, status, company_info, created_by
		FROM jobs
		WHERE id = $1
	`
	log.Printf("📡 Executing query: %s", query)
	
	err := db.DB.QueryRow(query, id).Scan(
		&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
		&job.Salary, &job.Description, &requirementsBytes, &responsibilitiesBytes,
		&benefitsBytes, &job.PostedDate, &job.Category, &job.Status,
		&companyInfoBytes, &createdBy, // Use sql.NullString for nullable created_by
	)
	
	// Check for errors while querying the database
	if err != nil {
		if err == sql.ErrNoRows {
			// Specific error for no matching rows
			log.Printf("❌ No job found with ID: %s", id)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		} else {
			// Other errors
			log.Printf("❌ Error querying database: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		}
		return
	}
	
	// Parse requirements, responsibilities, and benefits from JSON or CSV (depending on your storage format)
	if len(requirementsBytes) > 0 {
		if err := json.Unmarshal(requirementsBytes, &job.Requirements); err != nil {
			// If it's not valid JSON, treat it as a CSV and split by commas
			job.Requirements = strings.Split(string(requirementsBytes), ",")
		}
	}
	
	if len(responsibilitiesBytes) > 0 {
		if err := json.Unmarshal(responsibilitiesBytes, &job.Responsibilities); err != nil {
			// If it's not valid JSON, treat it as a CSV and split by commas
			job.Responsibilities = strings.Split(string(responsibilitiesBytes), ",")
		}
	}
	
	if len(benefitsBytes) > 0 {
		if err := json.Unmarshal(benefitsBytes, &job.Benefits); err != nil {
			// If it's not valid JSON, treat it as a CSV and split by commas
			job.Benefits = strings.Split(string(benefitsBytes), ",")
		}
	}
	
	// Parse company info if present
	if len(companyInfoBytes) > 0 {
		var companyInfo models.CompanyInfo
		if err := json.Unmarshal(companyInfoBytes, &companyInfo); err == nil {
			job.CompanyInfo = &companyInfo
		} else {
			log.Printf("⚠️ Error unmarshalling company info: %v", err)
		}
	}
	
	// Handle the createdBy field from sql.NullString
	if createdBy.Valid {
		job.CreatedBy = createdBy.String
	} else {
		job.CreatedBy = "" // Set to empty string if NULL
	}
	
	// Log the job data for debugging before sending the response
	log.Printf("🔍 Found job: %+v", job)
	
	// Return the job details in the response
	json.NewEncoder(w).Encode(job)
}




// CreateJob creates a new job listing (admin only)
// func CreateJob(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	// Check if user is admin
// 	userID, role, err := utils.GetUserFromToken(r)
// 	if err != nil || role != "admin" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
// 		return
// 	}
	
// 	// Parse request body
// 	var job models.Job
// 	err = json.NewDecoder(r.Body).Decode(&job)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
// 		return
// 	}
	
// 	// Validate required fields
// 	if job.Title == "" || job.Company == "" || job.Location == "" || 
// 	   job.Type == "" || job.Salary == "" || job.Description == "" || 
// 	   len(job.Requirements) == 0 || job.Category == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
// 		return
// 	}
	
// 	// Set default values
// 	job.Status = "active"
// 	job.CreatedBy = userID
	
// 	// Convert company info to JSON
// 	var companyInfoJSON []byte
// 	if job.CompanyInfo != nil {
// 		companyInfoJSON, err = json.Marshal(job.CompanyInfo)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error processing company info"})
// 			return
// 		}
// 	}
	
// 	// Insert job into database
// 	err = db.DB.QueryRow(`
// 		INSERT INTO jobs (
// 			title, company, location, type, salary, description, 
// 			requirements, responsibilities, benefits, category, 
// 			status, company_info, created_by
// 		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
// 		RETURNING id, posted_date
// 	`, job.Title, job.Company, job.Location, job.Type, job.Salary, 
// 	   job.Description, job.Requirements, job.Responsibilities, job.Benefits, 
// 	   job.Category, job.Status, companyInfoJSON, job.CreatedBy,
// 	).Scan(&job.ID, &job.PostedDate)
	
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to create job"})
// 		return
// 	}
	
// 	// Return created job
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(job)
// }

func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	fmt.Println("📩 Received request to create a job")

	// Parse request body
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		fmt.Println("❌ Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	fmt.Printf("🔍 Job Data Received: %+v\n", job)

	// Validate required fields
	if job.Title == "" || job.Company == "" || job.Location == "" || 
	   job.Type == "" || job.Salary == "" || job.Description == "" || 
	   len(job.Requirements) == 0 || job.Category == "" {
		fmt.Println("❌ Missing required fields")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
		return
	}

	// Convert company info to JSON (if provided)
	var companyInfoJSON []byte
	if job.CompanyInfo != nil {
		companyInfoJSON, err = json.Marshal(job.CompanyInfo)
		if err != nil {
			fmt.Println("❌ Error converting company info to JSON:", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error processing company info"})
			return
		}
	}

	// Insert job into database
	err = db.DB.QueryRow(`
		INSERT INTO jobs (
			title, company, location, type, salary, description, 
			requirements, responsibilities, benefits, category, 
			status, company_info
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, posted_date
	`, job.Title, job.Company, job.Location, job.Type, job.Salary, 
	   job.Description, pq.Array(job.Requirements), pq.Array(job.Responsibilities), pq.Array(job.Benefits), 
	   job.Category, "active", companyInfoJSON).Scan(&job.ID, &job.PostedDate)

	if err != nil {
		fmt.Println("❌ Database error while inserting job:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to create job"})
		return
	}

	fmt.Println("✅ Job successfully inserted into the database:", job.ID)

	// Return created job
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}





// UpdateJob updates an existing job (admin only)
func UpdateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is admin
	userID, role, err := utils.GetUserFromToken(r)
	if err != nil || role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Get job ID from URL
	params := mux.Vars(r)
	id := params["id"]
	
	// Parse request body
	var job models.Job
	err = json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Check if job exists and was created by this admin
	var createdBy string
	err = db.DB.QueryRow("SELECT created_by FROM jobs WHERE id = $1", id).Scan(&createdBy)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		return
	}
	
	if createdBy != userID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "You can only update jobs you created"})
		return
	}
	
	// Convert company info to JSON
	var companyInfoJSON []byte
	if job.CompanyInfo != nil {
		companyInfoJSON, err = json.Marshal(job.CompanyInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error processing company info"})
			return
		}
	}
	
	// Update job in database
	_, err = db.DB.Exec(`
		UPDATE jobs SET
			title = $1, company = $2, location = $3, type = $4, 
			salary = $5, description = $6, requirements = $7, 
			responsibilities = $8, benefits = $9, category = $10, 
			status = $11, company_info = $12
		WHERE id = $13
	`, job.Title, job.Company, job.Location, job.Type, job.Salary, 
	   job.Description, job.Requirements, job.Responsibilities, job.Benefits, 
	   job.Category, job.Status, companyInfoJSON, id)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to update job"})
		return
	}
	
	// Get updated job
	err = db.DB.QueryRow(`
		SELECT id, title, company, location, type, salary, description, 
		       requirements, responsibilities, benefits, posted_date, 
		       category, status, company_info, created_by
		FROM jobs
		WHERE id = $1
	`, id).Scan(
		&job.ID, &job.Title, &job.Company, &job.Location, &job.Type,
		&job.Salary, &job.Description, &job.Requirements, &job.Responsibilities,
		&job.Benefits, &job.PostedDate, &job.Category, &job.Status,
		&companyInfoJSON, &job.CreatedBy,
	)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to retrieve updated job"})
		return
	}
	
	// Parse company info if present
	if len(companyInfoJSON) > 0 {
		var companyInfo models.CompanyInfo
		if err := json.Unmarshal(companyInfoJSON, &companyInfo); err == nil {
			job.CompanyInfo = &companyInfo
		}
	}
	
	json.NewEncoder(w).Encode(job)
}

// DeleteJob deletes a job (admin only)
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is admin
	userID, role, err := utils.GetUserFromToken(r)
	if err != nil || role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Get job ID from URL
	params := mux.Vars(r)
	id := params["id"]
	
	// Check if job exists and was created by this admin
	var createdBy string
	err = db.DB.QueryRow("SELECT created_by FROM jobs WHERE id = $1", id).Scan(&createdBy)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		return
	}
	
	if createdBy != userID {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "You can only delete jobs you created"})
		return
	}
	
	// Delete job from database
	_, err = db.DB.Exec("DELETE FROM jobs WHERE id = $1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to delete job"})
		return
	}
	
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}