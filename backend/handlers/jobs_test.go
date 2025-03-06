package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
)

// ✅ Helper function to set up the router
func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	router.HandleFunc("/jobs/{id}", handlers.GetJobByID).Methods("GET")
	router.HandleFunc("/jobs", handlers.CreateJob).Methods("POST")
	router.HandleFunc("/jobs/{id}", handlers.UpdateJob).Methods("PUT")
	router.HandleFunc("/jobs/{id}", handlers.DeleteJob).Methods("DELETE")
	return router
}

// ✅ Initialize the database before running tests
func TestMain(m *testing.M) {
	db.InitDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}

// ✅ Test Creating a Job
// func TestCreateJob(t *testing.T) {
// 	router := setupRouter()

// 	reqBody := models.Job{
// 		Title:          "Software Engineer",
// 		Company:        "TechCorp",
// 		Location:       "New York",
// 		Type:           "Full-Time",
// 		Salary:         "$120,000",
// 		Description:    "Exciting role in tech",
// 		Requirements:   []string{"Go", "Microservices"},
// 		Responsibilities: []string{"Develop software", "Review code"},
// 		Benefits:       []string{"Health Insurance", "Stock Options"},
// 		Category:       "Technology",
// 	}

// 	body, _ := json.Marshal(reqBody)
// 	req, _ := http.NewRequest("POST", "/jobs", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)

// 	var response map[string]interface{}
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response["id"])
// }
func TestCreateJob(t *testing.T) {
	router := setupRouter()

	// ✅ Correct request body including "companyInfo"
	reqBody := `{
		"title": "Software Engineer",
		"company": "TechCorp",
		"location": "San Francisco, CA",
		"type": "Full-time",
		"salary": "$120,000 - $150,000",
		"description": "Building cutting-edge applications",
		"requirements": ["5+ years of experience", "Go", "PostgreSQL"],
		"responsibilities": ["Develop backend services", "Optimize database queries"],
		"benefits": ["Health insurance", "Remote work options"],
		"category": "Technology",
		"companyInfo": {
			"name": "TechCorp",
			"description": "Leading tech company",
			"website": "https://techcorp.com",
			"industry": "Software",
			"size": "1000+"
		}
	}`

	// ✅ Create request
	req, _ := http.NewRequest("POST", "/jobs", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	// ✅ Execute request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// ✅ Validate response
	assert.Equal(t, http.StatusCreated, w.Code)

	// ✅ Decode JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// ✅ Ensure required fields are present in response
	assert.True(t, response["id"] != nil, "Response should contain job ID")
	assert.Equal(t, "Software Engineer", response["title"])
	assert.Equal(t, "TechCorp", response["company"])
	assert.Equal(t, "Technology", response["category"])

	// ✅ Validate nested company info
	companyInfo, ok := response["companyInfo"].(map[string]interface{})
	assert.True(t, ok, "companyInfo should be a JSON object")
	assert.Equal(t, "TechCorp", companyInfo["name"])
	assert.Equal(t, "https://techcorp.com", companyInfo["website"])
}


// ✅ Test Fetching All Jobs
func TestGetJobs(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/jobs", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// ✅ Test Fetching a Single Job
func TestGetJobByID(t *testing.T) {
	router := setupRouter()

	// Assume there is a job with ID "12345" (you may change this)
	req, _ := http.NewRequest("GET", "/jobs/08508892-4e07-48e0-a0c5-cf30ea8e531c", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Job might exist or not
	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound}, w.Code)
}
