package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	// "strconv"
	"database/sql"
	"log"
	"strings"

	"fmt"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

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
		&companyInfoBytes, &createdBy,
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

	// Parse requirements, responsibilities, and benefits from JSON
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

	// Parse company info if present and not an empty array
	if len(companyInfoBytes) > 0 && string(companyInfoBytes) != "[]" {
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

func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("📩 Received request to create a job")

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
	if job.Title == "" || job.Company == "" || job.Location == "" || job.Type == "" ||
		job.Salary == "" || job.Description == "" || len(job.Requirements) == 0 || job.Category == "" {
		fmt.Println("❌ Missing required fields:", job)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
		return
	}

	// Generate ID and set timestamps
	job.ID = uuid.New().String()
	job.PostedDate = time.Now()
	job.Status = "active"

	// Get createdBy from context
	userID, ok := r.Context().Value("userID").(string)
	if ok {
		job.CreatedBy = userID
	} else {
		fmt.Println("⚠️ Could not get userID from context")
	}

	// Prepare requirementsJSON (required field)
	requirementsJSON, err := json.Marshal(job.Requirements)
	if err != nil {
		fmt.Println("❌ Error marshaling requirements:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid requirements format"})
		return
	}

	// Prepare responsibilitiesJSON
	var responsibilitiesJSON []byte
	if len(job.Responsibilities) > 0 {
		responsibilitiesJSON, err = json.Marshal(job.Responsibilities)
		if err != nil {
			fmt.Println("❌ Error marshaling responsibilities:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid responsibilities format"})
			return
		}
	} else {
		responsibilitiesJSON = []byte("[]") // Use empty JSON array instead of nil
	}

	// Prepare benefitsJSON
	var benefitsJSON []byte
	if len(job.Benefits) > 0 {
		benefitsJSON, err = json.Marshal(job.Benefits)
		if err != nil {
			fmt.Println("❌ Error marshaling benefits:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid benefits format"})
			return
		}
	} else {
		benefitsJSON = []byte("[]") // Use empty JSON array instead of nil
	}

	// Prepare companyInfoJSON
	var companyInfoJSON []byte
	if job.CompanyInfo != nil {
		companyInfoJSON, err = json.Marshal(job.CompanyInfo)
		if err != nil {
			fmt.Println("❌ Error marshaling company info:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid company info format"})
			return
		}
	} else {
		companyInfoJSON = []byte("[]") // Set to nil instead of []byte("[]")
	}

	fmt.Println("🔧 Prepared data for insertion:")
	fmt.Println("  ID:", job.ID)
	fmt.Println("  Title:", job.Title)
	fmt.Println("  Company:", job.Company)
	fmt.Println("  Location:", job.Location)
	fmt.Println("  Type:", job.Type)
	fmt.Println("  Salary:", job.Salary)
	fmt.Println("  Description:", job.Description)
	fmt.Println("  Requirements:", string(requirementsJSON))
	if responsibilitiesJSON == nil {
		fmt.Println("  Responsibilities: <nil>")
	} else {
		fmt.Println("  Responsibilities:", string(responsibilitiesJSON))
	}
	if benefitsJSON == nil {
		fmt.Println("  Benefits: <nil>")
	} else {
		fmt.Println("  Benefits:", string(benefitsJSON))
	}
	fmt.Println("  PostedDate:", job.PostedDate)
	fmt.Println("  Category:", job.Category)
	fmt.Println("  Status:", job.Status)
	if companyInfoJSON == nil {
		fmt.Println("  CompanyInfo: <nil>")
	} else {
		fmt.Println("  CompanyInfo:", string(companyInfoJSON))
	}
	fmt.Println("  CreatedBy:", job.CreatedBy)

	// Insert into database
	_, err = db.DB.Exec(`
    INSERT INTO jobs (id, title, company, location, type, salary, description, requirements, responsibilities, benefits, posted_date, category, status, company_info, created_by)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
`, job.ID, job.Title, job.Company, job.Location, job.Type, job.Salary, job.Description,
		requirementsJSON, responsibilitiesJSON, benefitsJSON, job.PostedDate, job.Category, job.Status, companyInfoJSON, job.CreatedBy)
	if err != nil {
		fmt.Println("❌ Database error while inserting job:", err)
		fmt.Println("  Requirements:", string(requirementsJSON))
		if responsibilitiesJSON == nil {
			fmt.Println("  Responsibilities: <nil>")
		} else {
			fmt.Println("  Responsibilities:", string(responsibilitiesJSON))
		}
		if benefitsJSON == nil {
			fmt.Println("  Benefits: <nil>")
		} else {
			fmt.Println("  Benefits:", string(benefitsJSON))
		}
		if companyInfoJSON == nil {
			fmt.Println("  CompanyInfo: <nil>")
		} else {
			fmt.Println("  CompanyInfo:", string(companyInfoJSON))
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to create job"})
		return
	}

	fmt.Println("✅ Job created successfully - ID:", job.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(job)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if user is admin
	userID, role, err := utils.GetUserFromToken(r)
	if err != nil || role != "admin" {
		fmt.Println("❌ Unauthorized: Invalid token or role")
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
		fmt.Println("❌ Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	fmt.Printf("🔍 Job Update Data Received: %+v\n", job)

	// Check if job exists and was created by this admin
	var createdBy string
	err = db.DB.QueryRow("SELECT created_by FROM jobs WHERE id = $1", id).Scan(&createdBy)
	if err != nil {
		fmt.Println("❌ Error checking job existence:", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
		return
	}

	if createdBy != userID {
		fmt.Println("❌ User not authorized to update this job - CreatedBy:", createdBy, "UserID:", userID)
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "You can only update jobs you created"})
		return
	}

	// Marshal JSON fields
	requirementsJSON, err := json.Marshal(job.Requirements)
	if err != nil {
		fmt.Println("❌ Error marshaling requirements:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid requirements format"})
		return
	}

	var responsibilitiesJSON []byte
	if len(job.Responsibilities) > 0 {
		responsibilitiesJSON, err = json.Marshal(job.Responsibilities)
		if err != nil {
			fmt.Println("❌ Error marshaling responsibilities:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid responsibilities format"})
			return
		}
	} else {
		responsibilitiesJSON = []byte("[]")
	}

	var benefitsJSON []byte
	if len(job.Benefits) > 0 {
		benefitsJSON, err = json.Marshal(job.Benefits)
		if err != nil {
			fmt.Println("❌ Error marshaling benefits:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid benefits format"})
			return
		}
	} else {
		benefitsJSON = []byte("[]")
	}

	var companyInfoJSON []byte
	if job.CompanyInfo != nil {
		companyInfoJSON, err = json.Marshal(job.CompanyInfo)
		if err != nil {
			fmt.Println("❌ Error marshaling company info:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid company info format"})
			return
		}
	} else {
		companyInfoJSON = nil // Set to nil to store NULL in the database
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
		job.Description, requirementsJSON, responsibilitiesJSON, benefitsJSON,
		job.Category, job.Status, companyInfoJSON, id)

	if err != nil {
		fmt.Println("❌ Database error while updating job:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to update job"})
		return
	}

	// Simplified response to avoid unmarshalling issues
	fmt.Println("✅ Job updated successfully - ID:", id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Job updated successfully",
		"jobId":   id,
	})
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

// SearchJobs searches for jobs based on a keyword and filters
// SearchJobs searches for jobs based on a keyword and filters
// SearchJobs searches for jobs based on a keyword and filters
func SearchJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to search jobs...")

	// Parse query parameters
	keyword := r.URL.Query().Get("keyword")
	category := r.URL.Query().Get("category")
	jobType := r.URL.Query().Get("type")
	location := r.URL.Query().Get("location")

	log.Printf("🔍 Search Filters - Keyword: %s | Category: %s | Type: %s | Location: %s\n", keyword, category, jobType, location)

	// Build the base query
	query := `
        SELECT id, title, company, location, type, salary, description, 
               requirements, responsibilities, benefits, posted_date, 
               category, status, company_info, COALESCE(created_by, '')
        FROM jobs
        WHERE status = 'active'
    `
	args := []interface{}{}
	conditions := []string{}

	// Add keyword search (search in title, company, and description)
	if keyword != "" {
		conditions = append(conditions, "(LOWER(title) LIKE LOWER($1) OR LOWER(company) LIKE LOWER($1) OR LOWER(description) LIKE LOWER($1))")
		args = append(args, "%"+keyword+"%")
	}

	// Add category filter
	if category != "" {
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(args)+1))
		args = append(args, category)
	}

	// Add job type filter
	if jobType != "" {
		conditions = append(conditions, fmt.Sprintf("type = $%d", len(args)+1))
		args = append(args, jobType)
	}

	// Add location filter
	if location != "" {
		conditions = append(conditions, fmt.Sprintf("LOWER(location) LIKE LOWER($%d)", len(args)+1))
		args = append(args, "%"+location+"%")
	}

	// Combine conditions
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	// Order by posted date
	query += " ORDER BY posted_date DESC"

	// Execute query
	log.Println("📡 Executing SQL Query:", query)
	rows, err := db.DB.Query(query, args...)
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
		var createdBy sql.NullString

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

		job.CreatedBy = createdBy.String
		jobs = append(jobs, job)
	}

	// Return response (empty list if no jobs match, not an error)
	log.Printf("📦 Returning %d jobs\n", len(jobs))
	json.NewEncoder(w).Encode(jobs)
}

// GetJobRecommendations provides job recommendations based on user skills
// GetJobRecommendations provides job recommendations based on user skills
func GetJobRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request for job recommendations...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Fetch user profile to get skills
	var user models.User
	var skillsJSON []byte
	err := db.DB.QueryRow(`
        SELECT id, email, password, full_name, title, location, bio, skills, role, created_at 
        FROM profiles 
        WHERE id = $1
    `, userID).Scan(
		&user.ID, &user.Email, &user.Password, &user.FullName,
		&user.Title, &user.Location, &user.Bio, &skillsJSON,
		&user.Role, &user.CreatedAt,
	)
	if err != nil {
		log.Println("❌ Error fetching user profile:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to fetch user profile"})
		return
	}

	// Parse user skills
	var userSkills []string
	if skillsJSON != nil {
		if err := json.Unmarshal(skillsJSON, &userSkills); err != nil {
			log.Println("❌ Error unmarshalling user skills:", err)
			userSkills = []string{}
		}
	}

	if len(userSkills) == 0 {
		log.Println("⚠️ User has no skills to base recommendations on")
		json.NewEncoder(w).Encode([]models.Job{})
		return
	}

	// Build query to find jobs matching user skills
	query := `
        SELECT id, title, company, location, type, salary, description, 
               requirements, responsibilities, benefits, posted_date, 
               category, status, company_info, COALESCE(created_by, '')
        FROM jobs
        WHERE status = 'active'
    `
	args := []interface{}{}
	conditions := []string{}

	// Add skill matching condition
	for i, skill := range userSkills {
		conditions = append(conditions, fmt.Sprintf("EXISTS (SELECT 1 FROM jsonb_array_elements_text(requirements) AS req WHERE LOWER(req) LIKE LOWER($%d))", i+1))
		args = append(args, "%"+skill+"%")
	}

	if len(conditions) > 0 {
		query += " AND (" + strings.Join(conditions, " OR ") + ")"
	}

	query += " ORDER BY posted_date DESC LIMIT 10"

	// Execute query
	log.Println("📡 Executing SQL Query:", query)
	rows, err := db.DB.Query(query, args...)
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
		var createdBy sql.NullString

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

		job.CreatedBy = createdBy.String
		jobs = append(jobs, job)
	}

	// Return response (empty list if no jobs match, not an error)
	log.Printf("📦 Returning %d recommended jobs\n", len(jobs))
	json.NewEncoder(w).Encode(jobs)
}
