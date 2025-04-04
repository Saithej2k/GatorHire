-- Create profiles table first (no dependencies)
CREATE TABLE profiles (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    full_name TEXT NOT NULL,
    title TEXT,
    location TEXT,
    bio TEXT,
    skills JSONB,
    role TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create jobs table (no dependencies)
CREATE TABLE jobs (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    company TEXT NOT NULL,
    location TEXT NOT NULL,
    type TEXT NOT NULL,
    salary TEXT NOT NULL,
    description TEXT NOT NULL,
    requirements JSONB NOT NULL,
    responsibilities JSONB,
    benefits JSONB,
    posted_date TIMESTAMP NOT NULL,
    category TEXT NOT NULL,
    status TEXT NOT NULL,
    company_info JSONB,
    created_by TEXT
);

-- Create applications table (depends on jobs and profiles)
CREATE TABLE applications (
    id TEXT PRIMARY KEY,
    job_id TEXT NOT NULL REFERENCES jobs(id),
    user_id TEXT NOT NULL REFERENCES profiles(id),
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    cover_letter TEXT,
    resume_url TEXT,
    linkedin TEXT,
    portfolio TEXT,
    heard_from TEXT,
    created_at TIMESTAMP NOT NULL,
    status TEXT NOT NULL
);

-- Create saved_jobs table (depends on profiles and jobs)
CREATE TABLE saved_jobs (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    job_id TEXT NOT NULL,
    saved_date TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES profiles(id),
    CONSTRAINT fk_job FOREIGN KEY (job_id) REFERENCES jobs(id)
);