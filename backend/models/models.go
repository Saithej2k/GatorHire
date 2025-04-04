package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// User represents a user in the system
type User struct {
	ID        string           `json:"id"`
	Email     string           `json:"email"`
	Password  string           `json:"password"` // Don't expose password in JSON
	FullName  string           `json:"fullName"`
	Title     *string          `json:"title,omitempty"`    // Changed to *string
	Location  *string          `json:"location,omitempty"` // Changed to *string
	Bio       *string          `json:"bio,omitempty"`      // Changed to *string
	Skills    *json.RawMessage `json:"skills"`             // Changed to *json.RawMessage
	Role      string           `json:"role,omitempty"`
	CreatedAt time.Time        `json:"createdAt"`
}

// Job represents a job posting
type Job struct {
	ID               string       `json:"id"`
	Title            string       `json:"title"`
	Company          string       `json:"company"`
	Location         string       `json:"location"`
	Type             string       `json:"type"`
	Salary           string       `json:"salary"`
	Description      string       `json:"description"`
	Requirements     []string     `json:"requirements"`
	Responsibilities []string     `json:"responsibilities,omitempty"`
	Benefits         []string     `json:"benefits,omitempty"`
	PostedDate       time.Time    `json:"postedDate"`
	Category         string       `json:"category"`
	Status           string       `json:"status"`
	CompanyInfo      *CompanyInfo `json:"companyInfo,omitempty"`
	CreatedBy        string       `json:"createdBy,omitempty"`
}

// CompanyInfo represents information about a company
type CompanyInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website,omitempty"`
	Industry    string `json:"industry,omitempty"`
	Size        string `json:"size,omitempty"`
}

// Value implements the driver.Valuer interface for CompanyInfo
func (c CompanyInfo) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implements the sql.Scanner interface for CompanyInfo
func (c *CompanyInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}

// Application represents a job application
type Application struct {
	ID          string    `json:"id"`
	JobID       string    `json:"jobId"`
	UserID      string    `json:"userId"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone,omitempty"`
	CoverLetter string    `json:"coverLetter,omitempty"`
	ResumeURL   string    `json:"resumeUrl,omitempty"`
	LinkedIn    string    `json:"linkedIn,omitempty"`
	Portfolio   string    `json:"portfolio,omitempty"`
	HeardFrom   string    `json:"heardFrom,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	Status      string    `json:"status,omitempty"`
}

// SavedJob represents a job saved by a user
type SavedJob struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	JobID     string    `json:"jobId"`
	SavedDate time.Time `json:"savedDate"`
}

// AuthResponse represents the response for authentication endpoints
type AuthResponse struct {
	Success bool   `json:"success"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Success bool `json:"success"`
}

// Initialize mock data slices
var (
	// Jobs is a slice of Job for easier handling of collections
	Jobs = []Job{
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

	// Users is a slice of User for easier handling of collections
	Users = []User{
		{
			ID:       "1",
			Email:    "john.doe@example.com",
			Password: "hashed_password", // In a real app, this would be properly hashed
			FullName: "John Doe",
			Role:     "user",
		},
	}

	// Applications is a slice of Application for easier handling of collections
	Applications = []Application{}

	// SavedJobs is a slice of SavedJob for easier handling of collections
	SavedJobs = []SavedJob{}
)
