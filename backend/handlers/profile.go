package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gatorhire/backend/db"
	"github.com/gatorhire/backend/models"
)

// ProfileStats represents user profile statistics
type ProfileStats struct {
	ApplicationsCount   int     `json:"applicationsCount"`
	SavedJobsCount      int     `json:"savedJobsCount"`
	ProfileCompleteness float64 `json:"profileCompleteness"` // Percentage (0-100)
}

// GetProfile retrieves the logged-in user's profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to fetch user profile...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Fetch user profile
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

	// Assign skillsJSON directly to user.Skills as *json.RawMessage
	if skillsJSON != nil {
		rawSkills := json.RawMessage(skillsJSON)
		user.Skills = &rawSkills
	}

	// Remove password from response
	user.Password = ""

	log.Printf("📦 Returning user profile for user ID: %s\n", userID)
	json.NewEncoder(w).Encode(user)
}

// UpdateProfile updates the logged-in user's profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request to update user profile...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Parse request body
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		log.Println("❌ Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	// Validate required fields
	if updatedUser.FullName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Full name is required"})
		return
	}

	// Convert skills to JSON
	var skillsJSON []byte
	if updatedUser.Skills != nil {
		skillsJSON, err = json.Marshal(updatedUser.Skills)
		if err != nil {
			log.Println("❌ Error marshaling skills:", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid skills format"})
			return
		}
	}

	// Update user profile in database
	_, err = db.DB.Exec(`
        UPDATE profiles 
        SET full_name = $1, title = $2, location = $3, bio = $4, skills = $5
        WHERE id = $6
    `, updatedUser.FullName, updatedUser.Title, updatedUser.Location, updatedUser.Bio, skillsJSON, userID)
	if err != nil {
		log.Println("❌ Database error while updating profile:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to update profile"})
		return
	}

	// Fetch updated user profile
	var user models.User
	err = db.DB.QueryRow(`
        SELECT id, email, password, full_name, title, location, bio, skills, role, created_at 
        FROM profiles 
        WHERE id = $1
    `, userID).Scan(
		&user.ID, &user.Email, &user.Password, &user.FullName,
		&user.Title, &user.Location, &user.Bio, &skillsJSON,
		&user.Role, &user.CreatedAt,
	)
	if err != nil {
		log.Println("❌ Error fetching updated user profile:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to fetch updated profile"})
		return
	}

	// Assign skillsJSON directly to user.Skills as *json.RawMessage
	if skillsJSON != nil {
		rawSkills := json.RawMessage(skillsJSON)
		user.Skills = &rawSkills
	}

	// Remove password from response
	user.Password = ""

	log.Printf("✅ Profile updated successfully for user ID: %s\n", userID)
	json.NewEncoder(w).Encode(user)
}

// GetProfileStats retrieves statistics for the logged-in user
func GetProfileStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("📩 Received request for profile stats...")

	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Fetch user profile
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

	// Count applications
	var applicationsCount int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM applications WHERE user_id = $1", userID).Scan(&applicationsCount)
	if err != nil {
		log.Println("❌ Error counting applications:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to fetch application count"})
		return
	}

	// Count saved jobs
	var savedJobsCount int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM saved_jobs WHERE user_id = $1", userID).Scan(&savedJobsCount)
	if err != nil {
		log.Println("❌ Error counting saved jobs:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Failed to fetch saved jobs count"})
		return
	}

	// Calculate profile completeness (simple heuristic: check if fields are non-empty)
	totalFields := 5.0  // full_name, title, location, bio, skills
	filledFields := 1.0 // full_name is required
	if user.Title != nil && *user.Title != "" {
		filledFields++
	}
	if user.Location != nil && *user.Location != "" {
		filledFields++
	}
	if user.Bio != nil && *user.Bio != "" {
		filledFields++
	}
	if skillsJSON != nil {
		var skills []string
		if err := json.Unmarshal(skillsJSON, &skills); err == nil && len(skills) > 0 {
			filledFields++
		}
	}
	profileCompleteness := (filledFields / totalFields) * 100

	// Construct response
	stats := ProfileStats{
		ApplicationsCount:   applicationsCount,
		SavedJobsCount:      savedJobsCount,
		ProfileCompleteness: profileCompleteness,
	}

	log.Printf("📦 Returning profile stats: %+v\n", stats)
	json.NewEncoder(w).Encode(stats)
}
