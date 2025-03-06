# GatorHire Database Commands Reference

This document provides common PostgreSQL commands for managing the GatorHire database.

## Basic Database Commands

### Connect to Database
```bash
psql -d gatorhire
```

### List Tables
```sql
\dt
```

### View Table Structure
```sql
\d profiles
\d jobs
\d applications
\d saved_jobs
```

## Query Examples

### View All Users
```sql
SELECT id, email, full_name, role FROM profiles;
```

### View All Jobs
```sql
SELECT id, title, company, location, type, category, status FROM jobs;
```

### View All Applications
```sql
SELECT a.id, a.job_id, j.title, a.full_name, a.email, a.status 
FROM applications a
JOIN jobs j ON a.job_id = j.id;
```

### View Jobs by Category
```sql
SELECT id, title, company, location, type, salary 
FROM jobs 
WHERE category = 'Technology' AND status = 'active';
```

### View Applications for a Specific Job
```sql
SELECT a.id, a.full_name, a.email, a.status, a.created_at
FROM applications a
WHERE a.job_id = 'job_id_here';
```

### View Saved Jobs for a User
```sql
SELECT j.id, j.title, j.company, j.location, j.type, s.saved_date
FROM saved_jobs s
JOIN jobs j ON s.job_id = j.id
WHERE s.user_id = 'user_id_here';
```

## Data Manipulation

### Add a New User
```sql
INSERT INTO profiles (id, email, password, full_name, role, created_at)
VALUES (
    gen_random_uuid(),
    'user@example.com',
    'hashed_password_here',
    'User Name',
    'user',
    NOW()
);
```

### Add a New Job
```sql
INSERT INTO jobs (
    title, company, location, type, salary, description, 
    requirements, responsibilities, benefits, category, status
)
VALUES (
    'Job Title',
    'Company Name',
    'Location',
    'Full-time',
    'Salary Range',
    'Job Description',
    ARRAY['Requirement 1', 'Requirement 2'],
    ARRAY['Responsibility 1', 'Responsibility 2'],
    ARRAY['Benefit 1', 'Benefit 2'],
    'Category',
    'active'
);
```

### Update Job Status
```sql
UPDATE jobs SET status = 'closed' WHERE id = 'job_id_here';
```

### Update Application Status
```sql
UPDATE applications SET status = 'interview' WHERE id = 'application_id_here';
```

### Delete a Job
```sql
DELETE FROM jobs WHERE id = 'job_id_here';
```

## Advanced Queries

### Count Jobs by Category
```sql
SELECT category, COUNT(*) 
FROM jobs 
WHERE status = 'active' 
GROUP BY category 
ORDER BY COUNT(*) DESC;
```

### Find Recent Applications
```sql
SELECT a.id, j.title, a.full_name, a.created_at, a.status
FROM applications a
JOIN jobs j ON a.job_id = j.id
ORDER BY a.created_at DESC
LIMIT 10;
```

### Find Users with Most Applications
```sql
SELECT p.full_name, COUNT(a.id) as application_count
FROM profiles p
JOIN applications a ON p.id = a.user_id
GROUP BY p.id, p.full_name
ORDER BY application_count DESC
LIMIT 5;
```

### Search Jobs by Keyword
```sql
SELECT id, title, company, location
FROM jobs
WHERE 
    title ILIKE '%keyword%' OR 
    description ILIKE '%keyword%' OR
    company ILIKE '%keyword%'
    AND status = 'active';
```

## Database Maintenance

### Backup Database
```bash
pg_dump -d gatorhire > gatorhire_backup.sql
```

### Restore Database
```bash
psql -d gatorhire < gatorhire_backup.sql
```

### Reset Database
```bash
dropdb gatorhire
createdb gatorhire
psql -d gatorhire -f backend/db/schema.sql
```