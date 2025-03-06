package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"io"
)

// Login handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
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
	err = db.DB.QueryRow(`
		SELECT id, email, password, full_name, title, location, bio, skills, role, created_at 
		FROM profiles 
		WHERE email = $1
	`, credentials.Email).Scan(
		&user.ID, &user.Email, &user.Password, &user.FullName, 
		&user.Title, &user.Location, &user.Bio, &user.Skills, 
		&user.Role, &user.CreatedAt,
	)
	
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid email or password"})
		return
	}
	
	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid email or password"})
		return
	}
	
	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to generate token"})
		return
	}
	
	// Return user data (excluding password)
	user.Password = ""
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.AuthResponse{
		Success: true,
		User: user,
		Token: token,
	})
}

// Register handles user registration
// func Register(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	fmt.Println("Received request on /api/auth/register")
	
// 	var newUser models.User
// 	err := json.NewDecoder(r.Body).Decode(&newUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
// 		return
// 	}

// 	fmt.Println("Raw request body:", newUser.Password)
	
// 	// Validate required fields
// 	if newUser.Email == "" || newUser.Password == "" || newUser.FullName == "" {
// 		fmt.Println("Missing required fields:", newUser)
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
// 		return
// 	}
	
// 	// Check if email already exists
// 	var count int
// 	err = db.DB.QueryRow("SELECT COUNT(*) FROM profiles WHERE email = $1", newUser.Email).Scan(&count)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
// 		return
// 	}
	
// 	if count > 0 {
// 		w.WriteHeader(http.StatusConflict)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Email already in use"})
// 		return
// 	}
	
// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to hash password"})
// 		return
// 	}
	
// 	// Generate ID
// 	newUser.ID = uuid.New().String()
// 	newUser.CreatedAt = time.Now()
// 	newUser.Role = "user" // Default role
	
// 	// Insert user into database
// 	_, err = db.DB.Exec(`
// 		INSERT INTO profiles (id, email, password, full_name, title, location, bio, skills, role, created_at)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
// 	`, newUser.ID, newUser.Email, string(hashedPassword), newUser.FullName, 
// 	   newUser.Title, newUser.Location, newUser.Bio, newUser.Skills, 
// 	   newUser.Role, newUser.CreatedAt)
	
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to create user"})
// 		return
// 	}
	
// 	// Generate JWT token
// 	token, err := utils.GenerateToken(newUser.ID, newUser.Email, newUser.Role)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to generate token"})
// 		return
// 	}
	
// 	// Return success response (excluding password)
// 	newUser.Password = ""
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(models.AuthResponse{
// 		Success: true,
// 		User: newUser,
// 		Token: token,
// 	})
// }

func Register(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Println("Received request on /api/auth/register")

    // Read and print raw request body for debugging
    body, err := io.ReadAll(r.Body)
    if err != nil {
        fmt.Println("Error reading request body:", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
        return
    }

    fmt.Println("Raw request body:", string(body)) // Print exact JSON being received

    // Decode JSON manually
    var newUser models.User
    err = json.Unmarshal(body, &newUser)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
        return
    }

    fmt.Printf("Decoded User Struct: %+v\n", newUser) // Print struct to check missing fields

    // Validate required fields
    if newUser.Email == "" || newUser.Password == "" || newUser.FullName == "" {
        fmt.Println("Missing required fields:", newUser)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Missing required fields"})
        return
    }

    // Convert `skills` to JSON format (because PostgreSQL expects JSONB)
    skillsJSON, err := json.Marshal(newUser.Skills)
    if err != nil {
        fmt.Println("Error converting skills to JSON:", err)
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to process skills data"})
        return
    }

    // Check if email already exists
    var count int
    err = db.DB.QueryRow("SELECT COUNT(*) FROM profiles WHERE email = $1", newUser.Email).Scan(&count)
    if err != nil {
        fmt.Println("Database error while checking email:", err) // Print error
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
        return
    }

    if count > 0 {
        w.WriteHeader(http.StatusConflict)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Email already in use"})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error hashing password:", err) // Print error
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to hash password"})
        return
    }

    // Generate ID
    newUser.ID = uuid.New().String()
    newUser.CreatedAt = time.Now()
    newUser.Role = "user" // Default role

    // Insert user into database
    _, err = db.DB.Exec(`
        INSERT INTO profiles (id, email, password, full_name, title, location, bio, skills, role, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    `, newUser.ID, newUser.Email, string(hashedPassword), newUser.FullName, 
       newUser.Title, newUser.Location, newUser.Bio, skillsJSON, 
       newUser.Role, newUser.CreatedAt)

    if err != nil {
        fmt.Println("Database error while inserting user:", err) // Print error
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Database error"})
        return
    }

    // Generate JWT token
    token, err := utils.GenerateToken(newUser.ID, newUser.Email, newUser.Role)
    if err != nil {
        fmt.Println("Error generating token:", err) // Print error
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to generate token"})
        return
    }

    // Return success response (excluding password)
    newUser.Password = ""
    fmt.Println("User registered successfully:", newUser.Email) // Confirm registration
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(models.AuthResponse{
        Success: true,
        User: newUser,
        Token: token,
    })
}
