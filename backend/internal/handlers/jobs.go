package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gatorhire/internal/db"
	"gatorhire/internal/models"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateJob(db.DB, &job); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := models.GetJobs(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func SearchJobs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	jobs, err := models.SearchJobs(db.DB, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var job models.Job
	err := db.DB.QueryRow(`SELECT * FROM jobs WHERE id = $1`, id).Scan(
		&job.ID,
		&job.Title,
		&job.Company,
		&job.Location,
		&job.Salary,
		&job.Type,
		&job.Description,
		&job.Requirements,
		&job.PostedDate,
	)

	if err != nil {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}