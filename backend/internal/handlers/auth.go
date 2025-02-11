package handlers

import (
	"encoding/json"
	"gatorhire/internal/db"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// SignUpRequest struct for parsing the sign-up request body
type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest struct for parsing the login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUp function handles user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	// Parse the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert user into the database
	_, err = db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		req.Name, req.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Error creating account", http.StatusInternalServerError)
		log.Println("Error inserting user into database:", err)
		return
	}

	// Respond to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account created successfully"})
}

// Login function handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Parse the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Query the database to find the user by email
	var storedPassword string
	var userID int
	err := db.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", req.Email).Scan(&userID, &storedPassword)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// If user is not found
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			// Log other errors
			http.Error(w, "Error querying user", http.StatusInternalServerError)
		}
		log.Println("Error querying user:", err)
		return
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.Password))
	if err != nil {
		// Passwords don't match
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                // Store the user ID in the token for future requests
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours)
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the generated token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
