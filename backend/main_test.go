package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gatorhire/backend/models"
	"github.com/gorilla/mux"
)

// Mock data for testing
var jobs = []models.Job{
	{
		ID:          "1",
		Title:       "Senior Frontend Developer",
		Company:     "TechCorp",
		Location:    "San Francisco, CA",
		Type:        "Full-time",
		Salary:      "$120,000 - $150,000",
		Description: "We are looking for an experienced Frontend Developer to join our team...",
		Requirements: []string{
			"5+ years of experience with React",
			"Strong TypeScript skills",
			"Experience with state management",
		},
		Category: "Technology",
	},
	{
		ID:          "2",
		Title:       "Backend Engineer",
		Company:     "DataSystems",
		Location:    "New York, NY",
		Type:        "Full-time",
		Salary:      "$130,000 - $160,000",
		Description: "Join our backend team to build scalable APIs and services...",
		Requirements: []string{
			"Experience with Go or similar languages",
			"Knowledge of SQL databases",
			"RESTful API design",
		},
		Category: "Technology",
	},
}

var users = []models.User{
	{
		ID:       "1",
		Email:    "john.doe@example.com",
		Password: "hashed_password", // In a real app, this would be properly hashed
		FullName: "John Doe",
		Role:     "user",
	},
}

var applications = []models.Application{}

// Mock handlers for testing
func getJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func getJobByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	id := params["id"]
	
	for _, job := range jobs {
		if job.ID == id {
			json.NewEncoder(w).Encode(job)
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Job not found"})
}

func createApplication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var app models.Application
	err := json.NewDecoder(r.Body).Decode(&app)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Validate required fields
	if app.JobID == "" || app.FullName == "" || app.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
		return
	}
	
	// Generate ID
	app.ID = "app-" + app.JobID + "-" + app.Email
	
	// Add to applications
	applications = append(applications, app)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"applicationId": app.ID,
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Find user by email
	var user models.User
	found := false
	for _, u := range users {
		if u.Email == credentials.Email {
			user = u
			found = true
			break
		}
	}
	
	if !found || user.Password != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid email or password"})
		return
	}
	
	// Generate mock token
	token := "mock_jwt_token"
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AuthResponse{
		Success: true,
		User:    user,
		Token:   token,
	})
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Validate required fields
	if newUser.Email == "" || newUser.Password == "" || newUser.FullName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
		return
	}
	
	// Check if email already exists
	for _, u := range users {
		if u.Email == newUser.Email {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Email already in use"})
			return
		}
	}
	
	// Generate ID
	newUser.ID = "user-" + newUser.Email
	newUser.Role = "user"
	
	// Add to users
	users = append(users, newUser)
	
	// Generate mock token
	token := "mock_jwt_token"
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.AuthResponse{
		Success: true,
		User:    newUser,
		Token:   token,
	})
}

func TestGetJobs(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/api/jobs", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getJobs)

	// Call the handler directly
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response []models.Job
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify we got the expected number of jobs
	if len(response) != len(jobs) {
		t.Errorf("Expected %d jobs, got %d", len(jobs), len(response))
	}
}

func TestGetJobByID(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()
	r.HandleFunc("/api/jobs/{id}", getJobByID)

	// Test cases
	testCases := []struct {
		name           string
		jobID          string
		expectedStatus int
		checkBody      bool
	}{
		{
			name:           "Valid Job ID",
			jobID:          "1",
			expectedStatus: http.StatusOK,
			checkBody:      true,
		},
		{
			name:           "Invalid Job ID",
			jobID:          "999",
			expectedStatus: http.StatusNotFound,
			checkBody:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request with the job ID
			req, err := http.NewRequest("GET", "/api/jobs/"+tc.jobID, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder
			rr := httptest.NewRecorder()

			// Serve the request
			r.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			// For valid job IDs, check the response body
			if tc.checkBody {
				var response models.Job
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}

				// Verify we got the expected job
				if response.ID != tc.jobID {
					t.Errorf("Expected job ID %s, got %s", tc.jobID, response.ID)
				}
			}
		})
	}
}

func TestCreateApplication(t *testing.T) {
	// Test application data
	application := models.Application{
		JobID:    "1",
		FullName: "Test User",
		Email:    "test@example.com",
		Phone:    "123-456-7890",
	}

	// Convert to JSON
	applicationJSON, err := json.Marshal(application)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the application data
	req, err := http.NewRequest("POST", "/api/applications", bytes.NewBuffer(applicationJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createApplication)

	// Initial count of applications
	initialCount := len(applications)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check that an application was added
	if len(applications) != initialCount+1 {
		t.Errorf("Expected %d applications, got %d", initialCount+1, len(applications))
	}

	// Check the response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify success is true
	success, ok := response["success"].(bool)
	if !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Verify application ID is present
	if _, ok := response["applicationId"]; !ok {
		t.Errorf("Expected applicationId in response, got %v", response)
	}
}

func TestLogin(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		credentials    map[string]string
		expectedStatus int
		checkSuccess   bool
	}{
		{
			name: "Valid Credentials",
			credentials: map[string]string{
				"email":    "john.doe@example.com",
				"password": "hashed_password", // This matches our mock data
			},
			expectedStatus: http.StatusOK,
			checkSuccess:   true,
		},
		{
			name: "Invalid Password",
			credentials: map[string]string{
				"email":    "john.doe@example.com",
				"password": "wrong_password",
			},
			expectedStatus: http.StatusUnauthorized,
			checkSuccess:   false,
		},
		{
			name: "User Not Found",
			credentials: map[string]string{
				"email":    "nonexistent@example.com",
				"password": "any_password",
			},
			expectedStatus: http.StatusUnauthorized,
			checkSuccess:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert credentials to JSON
			credentialsJSON, err := json.Marshal(tc.credentials)
			if err != nil {
				t.Fatal(err)
			}

			// Create a request with the credentials
			req, err := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(credentialsJSON))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(login)

			// Call the handler
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}

			// For successful login, check the response
			if tc.checkSuccess {
				var response map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}

				// Verify success is true
				success, ok := response["success"].(bool)
				if !ok || !success {
					t.Errorf("Expected success to be true, got %v", response["success"])
				}

				// Verify user and token are present
				if _, ok := response["user"]; !ok {
					t.Errorf("Expected user in response, got %v", response)
				}
				if _, ok := response["token"]; !ok {
					t.Errorf("Expected token in response, got %v", response)
				}
			}
		})
	}
}

func TestRegister(t *testing.T) {
	// Initial count of users
	initialCount := len(users)

	// Test user data
	newUser := models.User{
		Email:    "new.user@example.com",
		Password: "password123",
		FullName: "New User",
	}

	// Convert to JSON
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request with the user data
	req, err := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(register)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check that a user was added
	if len(users) != initialCount+1 {
		t.Errorf("Expected %d users, got %d", initialCount+1, len(users))
	}

	// Check the response
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify success is true
	success, ok := response["success"].(bool)
	if !ok || !success {
		t.Errorf("Expected success to be true, got %v", response["success"])
	}

	// Verify user and token are present
	if _, ok := response["user"]; !ok {
		t.Errorf("Expected user in response, got %v", response)
	}
	if _, ok := response["token"]; !ok {
		t.Errorf("Expected token in response, got %v", response)
	}
}