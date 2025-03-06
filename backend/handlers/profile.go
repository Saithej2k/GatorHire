package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
	"github.com/gatorhire/backend/utils"
)

// GetProfile returns the user's profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get user ID from token
	userID, _, err := utils.GetUserFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	// Query database for user profile
	var user models.User
	err = db.DB.QueryRow(`
		SELECT id, email, full_name, title, location, bio, skills, role, created_at 
		FROM profiles 
		WHERE id = $1
	`, userID).Scan(
		&user.ID, &user.Email, &user.FullName, &user.Title, 
		&user.Location, &user.Bio, &user.Skills, &user.Role, 
		&user.CreatedAt,
	)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "User not found"})
		return
	}
	
	json.NewEncoder(w).Encode(user)
}

// UpdateProfile updates the user's profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get user ID from token
	userID, _, err := utils.GetUserFromToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}
	
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	
	// Update user in database
	_, err = db.DB.Exec(`
		UPDATE profiles 
		SET full_name = $1, title = $2, location = $3, bio = $4, skills = $5
		WHERE id = $6
	`, updatedUser.FullName, updatedUser.Title, updatedUser.Location, 
	   updatedUser.Bio, updatedUser.Skills, userID)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to update profile"})
		return
	}
	
	// Get updated user
	err = db.DB.QueryRow(`
		SELECT id, email, full_name, title, location, bio, skills, role, created_at 
		FROM profiles 
		WHERE id = $1
	`, userID).Scan(
		&updatedUser.ID, &updatedUser.Email, &updatedUser.FullName, 
		&updatedUser.Title, &updatedUser.Location, &updatedUser.Bio, 
		&updatedUser.Skills, &updatedUser.Role, &updatedUser.CreatedAt,
	)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to retrieve updated profile"})
		return
	}
	
	json.NewEncoder(w).Encode(updatedUser)
}