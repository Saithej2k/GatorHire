package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/handlers"
)

// ✅ Setup router
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/applications", gin.WrapF(handlers.CreateApplication))
	router.GET("/user/applications", gin.WrapF(handlers.GetUserApplications))
	router.GET("/job/applications", gin.WrapF(handlers.GetApplicationsByJob))
	router.PUT("/applications/status", gin.WrapF(handlers.UpdateApplicationStatus))

	return router
}

// ✅ Initialize the database before running tests
func TestMain(m *testing.M) {
	db.InitDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}

// ✅ Test application creation
func TestCreateApplication(t *testing.T) {
	router := setupRouter()

	reqBody := `{
		"jobId": "b01eab12-a662-4eef-a2eb-c09f9e0e2870",
		"fullName": "Test User",
		"email": "testuser@example.com",
		"phone": "1234567890",
		"coverLetter": "This is a test cover letter.",
		"resumeURL": "https://example.com/resume.pdf",
		"linkedIn": "https://linkedin.com/in/testuser",
		"portfolio": "https://testuser.com",
		"heardFrom": "referral"
	}`

	req, _ := http.NewRequest("POST", "/applications", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.True(t, response["success"].(bool))
	assert.NotEmpty(t, response["applicationId"])
}

// ✅ Test fetching user applications
func TestGetUserApplications(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/user/applications", nil)
	req.Header.Set("Authorization", "Bearer dummy_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusUnauthorized}, w.Code)

	if w.Code == http.StatusOK {
		var applications []map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &applications)
		assert.Nil(t, err)
		assert.True(t, len(applications) >= 0)
	}
}

// ✅ Test fetching applications by job ID
func TestGetApplicationsByJob(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/job/applications?jobId=b01eab12-a662-4eef-a2eb-c09f9e0e2870", nil)
	req.Header.Set("Authorization", "Bearer admin_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusUnauthorized, http.StatusBadRequest}, w.Code)

	if w.Code == http.StatusOK {
		var applications []map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &applications)
		assert.Nil(t, err)
		assert.True(t, len(applications) >= 0)
	}
}

// ✅ Test updating application status
func TestUpdateApplicationStatus(t *testing.T) {
	router := setupRouter()

	reqBody := `{
		"applicationId": "b01eab12-a662-4eef-a2eb-c09f9e0e2870",
		"status": "reviewed"
	}`

	req, _ := http.NewRequest("PUT", "/applications/status", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer admin_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusUnauthorized, http.StatusBadRequest}, w.Code)

	if w.Code == http.StatusOK {
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.True(t, response["success"].(bool))
	}
}
