package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
	"github.com/google/uuid"
)

func CreateApplication(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Decode the request body into the application struct
    var app models.Application
    err := json.NewDecoder(r.Body).Decode(&app)
    if err != nil {
        log.Printf("❌ Error decoding request body: %v", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
        return
    }

    // Validate required fields
    if app.JobID == "" || app.FullName == "" || app.Email == "" {
        log.Printf("❌ Missing required fields: JobID: %s, FullName: %s, Email: %s", app.JobID, app.FullName, app.Email)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
        return
    }

    // Check if the job exists
    var jobExists bool
    query := "SELECT EXISTS(SELECT 1 FROM jobs WHERE id = $1)"
    log.Printf("📡 Executing query to check if job exists: %s with JobID: %s", query, app.JobID)
    err = db.DB.QueryRow(query, app.JobID).Scan(&jobExists)
    if err != nil {
        log.Printf("❌ Error checking job existence: %v", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
        return
    }

    if !jobExists {
        log.Printf("❌ Job not found with JobID: %s", app.JobID)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
        return
    }

    // Check if user has already applied to this job
    var alreadyApplied bool
    query = `
        SELECT EXISTS(
            SELECT 1 FROM applications 
            WHERE job_id = $1 AND email = $2
        )
    `
    log.Printf("📡 Executing query to check if user already applied: %s", query)
    err = db.DB.QueryRow(query, app.JobID, app.Email).Scan(&alreadyApplied)
    if err != nil {
        log.Printf("❌ Error checking if user has already applied: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
        return
    }

    if alreadyApplied {
        log.Printf("❌ User has already applied to the job with JobID: %s", app.JobID)
        w.WriteHeader(http.StatusConflict)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "You have already applied to this job"})
        return
    }

    // Generate ID and set creation time
    app.ID = uuid.New().String()
    app.Status = "pending"
    app.CreatedAt = time.Now()

    // Retrieve the user ID associated with the provided email
    var userID string
    userQuery := "SELECT id FROM profiles WHERE email = $1"
	log.Printf("userquery", userQuery);
    err = db.DB.QueryRow(userQuery, app.Email).Scan(&userID)
    if err != nil {
        log.Printf("❌ Error retrieving user ID: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "User not found"})
        return
    }
	log.Printf("userrid", userID);
    app.UserID = userID // Set the UserID field

    // Insert application into the database
    insertQuery := `
        INSERT INTO applications (
            id, job_id, user_id, full_name, email, phone, 
            cover_letter, resume_url, linkedin, portfolio, 
            heard_from, created_at, status
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    `
    log.Printf("📡 Executing insert query: %s", insertQuery)
    _, err = db.DB.Exec(insertQuery, app.ID, app.JobID, app.UserID, app.FullName, app.Email, app.Phone, 
        app.CoverLetter, app.ResumeURL, app.LinkedIn, app.Portfolio, 
        app.HeardFrom, app.CreatedAt, app.Status)

    if err != nil {
        log.Printf("❌ Error inserting application into database: %v", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to submit application"})
        return
    }

    // Return success response
    log.Printf("✅ Application submitted successfully with ID: %s", app.ID)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "applicationId": app.ID,
    })
}





// GetUserApplications returns all applications for the authenticated user
// func GetUserApplications(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
	
// 	// Check if user is authenticated
// 	userID, _, err := utils.GetUserFromToken(r)
// 	log.Printf("userid", userID, err)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
// 		return
// 	}
	
// 	// Query database for user's applications
// 	rows, err := db.DB.Query(`
// 		SELECT a.id, a.job_id, a.full_name, a.email, a.created_at, a.status,
// 		       j.title, j.company
// 		FROM applications a
// 		JOIN jobs j ON a.job_id = j.id
// 		WHERE a.user_id = $1
// 		ORDER BY a.created_at DESC
// 	`, userID)
	
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
// 		return
// 	}
// 	defer rows.Close()
	
// 	// Parse results
// 	var applications []map[string]interface{}
// 	for rows.Next() {
// 		var app struct {
// 			ID        string
// 			JobID     string
// 			FullName  string
// 			Email     string
// 			CreatedAt time.Time
// 			Status    string
// 			JobTitle  string
// 			Company   string
// 		}
		
// 		err := rows.Scan(
// 			&app.ID, &app.JobID, &app.FullName, &app.Email, 
// 			&app.CreatedAt, &app.Status, &app.JobTitle, &app.Company,
// 		)
		
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing application data"})
// 			return
// 		}
		
// 		// Format for response
// 		applications = append(applications, map[string]interface{}{
// 			"id":          app.ID,
// 			"jobId":       app.JobID,
// 			"jobTitle":    app.JobTitle,
// 			"company":     app.Company,
// 			"appliedDate": app.CreatedAt,
// 			"status":      app.Status,
// 		})
// 	}
	
// 	json.NewEncoder(w).Encode(applications)
// }
func GetUserApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Log the incoming request
	log.Println("📩 Received request to fetch user applications")

	// Check if user is authenticated
	userID, _, err := utils.GetUserFromToken(r)
	log.Printf("🔍 Extracted user ID: %s, Error: %v", userID, err)

	if err != nil {
		log.Println("❌ Unauthorized: No valid user found")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Query database for user's applications
	log.Printf("📡 Executing query to fetch applications for user ID: %s", userID)
	rows, err := db.DB.Query(`
		SELECT a.id, a.job_id, a.full_name, a.email, a.created_at, a.status,
		       j.title, j.company
		FROM applications a
		JOIN jobs j ON a.job_id = j.id
		WHERE a.user_id = $1
		ORDER BY a.created_at DESC
	`, userID)

	if err != nil {
		log.Printf("❌ Error querying database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		return
	}
	defer rows.Close()

	// Parse results
	var applications []map[string]interface{}
	for rows.Next() {
		var app struct {
			ID        string
			JobID     string
			FullName  string
			Email     string
			CreatedAt time.Time
			Status    string
			JobTitle  string
			Company   string
		}

		// Scan each row into the app struct
		err := rows.Scan(
			&app.ID, &app.JobID, &app.FullName, &app.Email,
			&app.CreatedAt, &app.Status, &app.JobTitle, &app.Company,
		)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing application data"})
			return
		}

		// Log each application found
		log.Printf("🔍 Fetched application: %+v", app)

		// Format for response
		applications = append(applications, map[string]interface{}{
			"id":          app.ID,
			"jobId":       app.JobID,
			"jobTitle":    app.JobTitle,
			"company":     app.Company,
			"appliedDate": app.CreatedAt,
			"status":      app.Status,
		})
	}

	// Log the number of applications found
	log.Printf("✅ Found %d applications for user ID: %s", len(applications), userID)
	log.Printf("apps", applications)
	// Send response
	json.NewEncoder(w).Encode(applications)
}


// GetApplicationsByJob returns all applications for a specific job (admin only)
func GetApplicationsByJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is admin
	_, role, err := utils.GetUserFromToken(r)
	if err != nil || role != "admin" {
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
	
	// Query database for job applications
	rows, err := db.DB.Query(`
		SELECT a.id, a.job_id, a.user_id, a.full_name, a.email, 
		       a.phone, a.cover_letter, a.resume_url, a.linkedin, 
		       a.portfolio, a.heard_from, a.created_at, a.status
		FROM applications a
		WHERE a.job_id = $1
		ORDER BY a.created_at DESC
	`, jobID)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
		return
	}
	defer rows.Close()
	
	// Parse results
	var applications []models.Application
	for rows.Next() {
		var app models.Application
		
		err := rows.Scan(
			&app.ID, &app.JobID, &app.UserID, &app.FullName, &app.Email,
			&app.Phone, &app.CoverLetter, &app.ResumeURL, &app.LinkedIn,
			&app.Portfolio, &app.HeardFrom, &app.CreatedAt, &app.Status,
		)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Error parsing application data"})
			return
		}
		
		applications = append(applications, app)
	}
	
	json.NewEncoder(w).Encode(applications)
}

// UpdateApplicationStatus updates the status of an application (admin only)
func UpdateApplicationStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is admin
	_, role, err := utils.GetUserFromToken(r)
	if err != nil || role != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Parse request body
	var request struct {
		ApplicationID string `json:"applicationId"`
		Status        string `json:"status"`
	}
	
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"reviewed":  true,
		"interview": true,
		"rejected":  true,
		"accepted":  true,
	}
	
	if !validStatuses[request.Status] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid status"})
		return
	}
	
	// Update application status
	_, err = db.DB.Exec(
		"UPDATE applications SET status = $1 WHERE id = $2",
		request.Status, request.ApplicationID,
	)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to update application status"})
		return
	}
	
	json.NewEncoder(w).Encode(models.SuccessResponse{Success: true})
}