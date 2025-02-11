package models

import (
	"time"
	"database/sql"
)

type Job struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Company      string    `json:"company"`
	Location     string    `json:"location"`
	Salary       string    `json:"salary"`
	Type         string    `json:"type"`
	Description  string    `json:"description"`
	Requirements []string  `json:"requirements"`
	PostedDate   time.Time `json:"postedDate"`
}

func CreateJob(db *sql.DB, job *Job) error {
	query := `
		INSERT INTO jobs (title, company, location, salary, type, description, requirements)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, posted_date`

	return db.QueryRow(
		query,
		job.Title,
		job.Company,
		job.Location,
		job.Salary,
		job.Type,
		job.Description,
		job.Requirements,
	).Scan(&job.ID, &job.PostedDate)
}

func GetJobs(db *sql.DB) ([]Job, error) {
	rows, err := db.Query(`SELECT * FROM jobs ORDER BY posted_date DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var job Job
		err := rows.Scan(
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
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func SearchJobs(db *sql.DB, query string) ([]Job, error) {
	rows, err := db.Query(`
		SELECT * FROM jobs 
		WHERE search_vector @@ plainto_tsquery('english', $1)
		ORDER BY posted_date DESC`,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var job Job
		err := rows.Scan(
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
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}