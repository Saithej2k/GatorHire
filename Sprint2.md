# GatorHire - Sprint 2 Documentation

## Work Completed in Sprint 2

### Frontend
- Implemented responsive UI for all pages using React and Tailwind CSS
- Created the following pages:
  - Home page with job search functionality and category selection
  - Job listings page with filtering options by category, job type, and location
  - Job details page with comprehensive job information
  - Application submission form with validation
  - User profile page with editable fields
  - Dashboard page with application tracking
  - 404 Not Found page
- Added navigation with responsive mobile menu
- Implemented form validation for job applications
- Created reusable UI components (Button, Card, Badge, etc.)
- Added unit tests for components and services
- Implemented Cypress E2E test for application form functionality

### Backend
- Developed a RESTful API using Go
- Implemented the following endpoints:
  - GET /api/jobs - Retrieve all job listings
  - GET /api/jobs/{id} - Retrieve a specific job by ID
  - POST /api/applications - Submit a job application
  - GET /api/applications/user - Get user's applications
  - POST /api/auth/login - User login
  - POST /api/auth/register - User registration
  - GET /api/profile - Get user profile
  - PUT /api/profile - Update user profile
- Added CORS support for cross-origin requests
- Implemented unit tests for API endpoints
- Created mock data for development and testing

### Integration
- Connected frontend to backend API
- Implemented error handling for API requests
- Added loading states for asynchronous operations
- Created shared TypeScript interfaces for data models

## Frontend Unit Tests

### Component Tests
1. **JobCard.test.tsx**
   - Tests if job information renders correctly
   - Tests if category badge displays correctly
   - Tests if "View Details" link has the correct URL

2. **CategoryFilter.test.tsx**
   - Tests if all category buttons render correctly
   - Tests if category counts display correctly
   - Tests if selected category is highlighted
   - Tests if clicking a category calls the change handler

3. **SearchInput.test.tsx**
   - Tests if placeholder text renders correctly
   - Tests if input displays the current value
   - Tests if onChange handler is called when input changes
   - Tests if custom class is applied correctly

### Service Tests
1. **api.test.ts**
   - Tests if `fetchJobs()` retrieves all jobs correctly
   - Tests if `fetchJobById()` retrieves a specific job correctly
   - Tests if `submitApplication()` submits an application correctly
   - Tests error handling for all API functions

### Cypress E2E Test
1. **applicationForm.cy.ts**
   - Tests navigation to the application form from job details
   - Tests form validation for required fields
   - Tests successful submission of a completed application form

## Backend Unit Tests

1. **TestGetJobs**
   - Tests if the `/api/jobs` endpoint returns all jobs
   - Verifies the correct HTTP status code and response body

2. **TestGetJobByID**
   - Tests if the `/api/jobs/{id}` endpoint returns the correct job
   - Tests handling of invalid job IDs
   - Verifies the correct HTTP status code and response body

3. **TestCreateApplication**
   - Tests if the `/api/applications` endpoint creates a new application
   - Tests validation of required fields
   - Verifies the application is added to the database
   - Checks for the correct HTTP status code and response body

4. **TestLogin**
   - Tests if the `/api/auth/login` endpoint authenticates users correctly
   - Tests handling of invalid credentials
   - Verifies the correct HTTP status code and response body

5. **TestRegister**
   - Tests if the `/api/auth/register` endpoint creates a new user
   - Verifies the user is added to the database
   - Checks for the correct HTTP status code and response body

## Backend API Documentation

### Jobs API

#### GET /api/jobs
Retrieves all job listings.

**Response:**
```json
[
  {
    "id": "1",
    "title": "Senior Frontend Developer",
    "company": "TechCorp",
    "location": "San Francisco, CA",
    "type": "Full-time",
    "salary": "$120,000 - $150,000",
    "description": "We are looking for an experienced Frontend Developer to join our team...",
    "requirements": ["5+ years of experience with React", "Strong TypeScript skills", "Experience with state management"],
    "responsibilities": ["Develop new user-facing features", "Build reusable components", "Translate designs into code"],
    "benefits": ["Competitive salary", "Health insurance", "Unlimited PTO"],
    "postedDate": "2023-04-15",
    "category": "Technology",
    "companyInfo": {
      "name": "TechCorp",
      "description": "TechCorp is a leading technology company...",
      "website": "https://techcorp.example.com",
      "industry": "Software Development",
      "size": "500-1000 employees"
    }
  },
  // More jobs...
]
```

#### GET /api/jobs/{id}
Retrieves a specific job by ID.

**Parameters:**
- `id` (path parameter): The ID of the job to retrieve

**Response (Success):**
```json
{
  "id": "1",
  "title": "Senior Frontend Developer",
  "company": "TechCorp",
  "location": "San Francisco, CA",
  "type": "Full-time",
  "salary": "$120,000 - $150,000",
  "description": "We are looking for an experienced Frontend Developer to join our team...",
  "requirements": ["5+ years of experience with React", "Strong TypeScript skills", "Experience with state management"],
  "responsibilities": ["Develop new user-facing features", "Build reusable components", "Translate designs into code"],
  "benefits": ["Competitive salary", "Health insurance", "Unlimited PTO"],
  "postedDate": "2023-04-15",
  "category": "Technology",
  "companyInfo": {
    "name": "TechCorp",
    "description": "TechCorp is a leading technology company...",
    "website": "https://techcorp.example.com",
    "industry": "Software Development",
    "size": "500-1000 employees"
  }
}
```

**Response (Error):**
```json
{
  "error": "Job not found"
}
```

### Applications API

#### POST /api/applications
Submits a job application.

**Request Body:**
```json
{
  "jobId": "1",
  "fullName": "John Doe",
  "email": "john.doe@example.com",
  "phone": "123-456-7890",
  "coverLetter": "I am writing to express my interest in the Senior Frontend Developer position...",
  "resumeUrl": "https://example.com/resume.pdf",
  "linkedIn": "https://linkedin.com/in/johndoe",
  "portfolio": "https://johndoe.com",
  "heardFrom": "job-board"
}
```

**Response (Success):**
```json
{
  "success": true,
  "applicationId": "1"
}
```

**Response (Error):**
```json
{
  "error": "Missing required fields"
}
```

#### GET /api/applications/user
Retrieves all applications for the authenticated user.

**Headers:**
- `Authorization`: Bearer token for authentication

**Response (Success):**
```json
[
  {
    "id": "1",
    "jobId": "1",
    "fullName": "John Doe",
    "email": "john.doe@example.com",
    "phone": "123-456-7890",
    "coverLetter": "I am writing to express my interest...",
    "resumeUrl": "https://example.com/resume.pdf",
    "linkedIn": "https://linkedin.com/in/johndoe",
    "portfolio": "https://johndoe.com",
    "heardFrom": "job-board",
    "createdAt": "2023-04-20T10:30:00Z",
    "status": "pending"
  },
  // More applications...
]
```

### Authentication API

#### POST /api/auth/login
Authenticates a user.

**Request Body:**
```json
{
  "email": "john.doe@example.com",
  "password": "password123"
}
```

**Response (Success):**
```json
{
  "success": true,
  "user": {
    "id": "1",
    "email": "john.doe@example.com",
    "fullName": "John Doe",
    "title": "Senior Frontend Developer",
    "location": "San Francisco, CA",
    "bio": "Experienced frontend developer...",
    "skills": ["React", "TypeScript", "JavaScript", "HTML/CSS"],
    "role": "user",
    "createdAt": "2023-04-01T12:00:00Z"
  },
  "token": "jwt_token_here"
}
```

**Response (Error):**
```json
{
  "error": "Invalid email or password"
}
```

#### POST /api/auth/register
Registers a new user.

**Request Body:**
```json
{
  "email": "new.user@example.com",
  "password": "password123",
  "fullName": "New User",
  "title": "Frontend Developer",
  "location": "New York, NY",
  "bio": "Frontend developer with 3 years of experience...",
  "skills": ["React", "JavaScript", "CSS"]
}
```

**Response (Success):**
```json
{
  "success": true,
  "user": {
    "id": "2",
    "email": "new.user@example.com",
    "fullName": "New User",
    "title": "Frontend Developer",
    "location": "New York, NY",
    "bio": "Frontend developer with 3 years of experience...",
    "skills": ["React", "JavaScript", "CSS"],
    "role": "user",
    "createdAt": "2023-04-20T10:30:00Z"
  },
  "token": "jwt_token_here"
}
```

**Response (Error):**
```json
{
  "error": "Email already in use"
}
```

### Profile API

#### GET /api/profile
Retrieves the user's profile.

**Headers:**
- `Authorization`: Bearer token for authentication

**Response (Success):**
```json
{
  "id": "1",
  "email": "john.doe@example.com",
  "fullName": "John Doe",
  "title": "Senior Frontend Developer",
  "location": "San Francisco, CA",
  "bio": "Experienced frontend developer...",
  "skills": ["React", "TypeScript", "JavaScript", "HTML/CSS"],
  "role": "user",
  "createdAt": "2023-04-01T12:00:00Z"
}
```

**Response (Error):**
```json
{
  "error": "User not found"
}
```

#### PUT /api/profile
Updates the user's profile.

**Headers:**
- `Authorization`: Bearer token for authentication

**Request Body:**
```json
{
  "fullName": "John Doe",
  "title": "Lead Frontend Developer",
  "location": "San Francisco, CA",
  "bio": "Updated bio information...",
  "skills": ["React", "TypeScript", "JavaScript", "HTML/CSS", "Redux"]
}
```

**Response (Success):**
```json
{
  "id": "1",
  "email": "john.doe@example.com",
  "fullName": "John Doe",
  "title": "Lead Frontend Developer",
  "location": "San Francisco, CA",
  "bio": "Updated bio information...",
  "skills": ["React", "TypeScript", "JavaScript", "HTML/CSS", "Redux"],
  "role": "user",
  "createdAt": "2023-04-01T12:00:00Z"
}
```

**Response (Error):**
```json
{
  "error": "User not found"
}
```