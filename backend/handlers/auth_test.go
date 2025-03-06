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
	"github.com/gatorhire/backend/models"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Register auth routes
	router.POST("/register", gin.WrapF(handlers.Register))
	router.POST("/login", gin.WrapF(handlers.Login))

	return router
}

func TestMain(m *testing.M) {
	// Initialize database connection before running tests
	db.InitDB()

	// Run tests
	exitVal := m.Run()

	// Exit with the test run status
	os.Exit(exitVal)
}

func TestRegister(t *testing.T) {
	router := setupRouter()

	// Ensure the user does not exist before testing registration
	_, _ = db.DB.Exec("DELETE FROM profiles WHERE email = $1", "testuser@example.com")

	// Use raw JSON request to match expected API structure
	reqBody := `{
		"email": "testuser@example.com",
		"password": "password123",
		"fullName": "Test User"
	}`

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// ✅ Check HTTP status code
	assert.Equal(t, http.StatusCreated, w.Code, "Expected 201 Created but got %d", w.Code)

	// ✅ Decode response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err, "Failed to decode response JSON")

	// 🛠 DEBUG: Print full response for debugging
	t.Logf("Response Body: %v", response)

	// ✅ Check if "success" field exists and is true
	success, successExists := response["success"]
	assert.True(t, successExists, "Response should contain 'success' field")
	assert.True(t, success.(bool), "Expected 'success' field to be true")

	// ✅ Ensure a token is returned
	token, tokenExists := response["token"]
	assert.True(t, tokenExists, "Response should contain 'token' field")
	assert.NotEmpty(t, token, "Token should not be empty")

	// ✅ Ensure user object exists in response
	user, userExists := response["user"]
	assert.True(t, userExists, "Response should contain 'user' field")

	// ✅ Ensure user email matches
	userMap := user.(map[string]interface{})
	assert.Equal(t, "testuser@example.com", userMap["email"], "Email should match the registered user")
}





func TestLogin(t *testing.T) {
	router := setupRouter()

	// First, register the user
	reqBody := models.User{
		Email:    "testusser@example.com",
		Password: "password123",
		FullName: "Test User",
	}

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Now test login
	loginBody := `{"email":"testuser@example.com","password":"password123"}`
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(loginBody)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// ✅ FIX: Remove duplicate `response` declaration
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.True(t, response["success"].(bool))
	assert.NotEmpty(t, response["token"])
}
