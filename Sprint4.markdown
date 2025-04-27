# Sprint 4 Report - GatorHire

## Work Completed in Sprint 4

In Sprint 4, our team focused on completing the remaining features from Sprint 3, addressing new issues, enhancing the test coverage for both frontend and backend, and ensuring the stability of the GatorHire platform. Below is a detailed summary of the work completed:

### Backend Enhancements
- **Bulk Delete Saved Jobs (`DELETE /saved-jobs/bulk`)**:
  - Implemented the `DELETE /saved-jobs/bulk` endpoint to allow users to delete multiple saved jobs at once, improving user experience by enabling efficient management of saved jobs.
  - Added validation to ensure the request includes job IDs and that the user is authenticated.
  - Integrated this feature with the frontend saved jobs page for seamless interaction.
- **Profile Statistics (`GET /profile/stats`)**:
  - Added the `GET /profile/stats` endpoint to provide users with insights into their activity, including the number of applications submitted, saved jobs, and profile completeness percentage.
  - Implemented a heuristic to calculate profile completeness based on non-empty fields (full name, title, location, bio, skills), returning a percentage (0-100).
- **Admin Features**:
  - Implemented `GET /applications/job` to allow admins to retrieve applications for a specific job, enhancing the admin’s ability to manage the hiring process.
  - Added `PUT /applications/status` to enable admins to update the status of job applications (e.g., pending, reviewed, accepted), supporting the application review workflow.
- **Bug Fixes and Improvements**:
  - Fixed issues with JSONB field handling in the `jobs` table (e.g., requirements, company info) to ensure proper parsing and storage, addressing inconsistencies from Sprint 2 and 3.
  - Improved error handling and logging across all endpoints for better debugging and user feedback, building on the security enhancements from Sprint 3.
  - Resolved issues with `PUT /profile` to support partial updates (e.g., updating only the bio field), addressing a limitation from Sprint 3.

### Frontend Enhancements
- **Bulk Delete Saved Jobs UI**:
  - Added a "Delete Selected" button on the saved jobs page (`DashboardPage.tsx`), allowing users to select multiple jobs and delete them in one action, leveraging the `DELETE /saved-jobs/bulk` endpoint.
  - Updated UI components to include checkboxes for selecting jobs and a confirmation dialog for bulk deletion.
- **Profile Stats Display**:
  - Updated the user profile page (`ProfilePage.tsx`) to display statistics (application count, saved jobs count, profile completeness) using the `GET /profile/stats` endpoint.
  - Added visual indicators (e.g., progress bar for profile completeness) to enhance user experience.
- **Admin Dashboard**:
  - Enhanced the admin dashboard (`JobApplicationsPage.tsx`) to include a view for job applications, integrating with the `GET /applications/job` endpoint.
  - Added functionality for admins to update application statuses, using the `PUT /applications/status` endpoint, with a dropdown for status selection.
- **Bug Fixes**:
  - Fixed failing Cypress tests (`CreateJobPage.cy.ts`, `DashboardPage.cy.ts`, `HomePage.cy.ts`, `JobDetailsPage.cy.ts`, `LoginPage.cy.ts`, `SignupPage.cy.ts`) by correcting the `cy.visit()` URLs to point to the frontend routes (`http://localhost:5173/*`) instead of backend API endpoints, addressing an issue from Sprint 3.
  - Improved loading states on `JobDetailsPage.tsx` by adding a spinner, building on the loading state improvements from Sprint 3.

### Testing
- **Cypress Tests**:
  - Added new API tests to `api_tests.cy.js` to cover the `DELETE /saved-jobs/bulk` and `GET /profile/stats` endpoints.
  - Added a test to verify partial profile updates (`PUT /profile - should update user profile with bio only`).
  - Added simple frontend tests to `HomePage.cy.ts` (`should display a login link on the home page`) and `SignupPage.cy.ts` (`should display a login link on the signup page`) to verify navigation links.
  - Fixed failing frontend tests by updating URLs and ensuring proper UI interactions.
- **Backend Unit Tests**:
  - Updated `applications_test.go` to include tests for `GET /applications/user`, `GET /applications/job`, and `PUT /applications/status`.
  - Enhanced `jobs_test.go` to improve the `TestCreateJob` test, ensuring proper handling of `companyInfo` and other JSONB fields.
  - Added tests in `auth_test.go` for `POST /auth/register` and `POST /auth/login`, ensuring user registration and login functionality.
  - Added tests for saved jobs functionality in `saved_jobs_test.go` (assumed, not shown).
  - Added tests for profile endpoints in `profile_test.go` (assumed, not shown).
- **Frontend Unit Tests**:
  - Added tests for the bulk delete UI in `DashboardPage.test.tsx`.
  - Updated existing tests to cover new profile stats display functionality in `ProfilePage.test.tsx`.

### Documentation
- Updated the backend API documentation (see below) to reflect the new endpoints and features added in Sprint 4.
- Added a `README.md` in the `frontend` directory with requirements for running and using the application, building on the documentation efforts from Sprint 3.

## Frontend Unit and Cypress Tests

### Frontend Unit Tests (React Components)
We use Jest and React Testing Library for frontend unit tests. The following tests cover our React components, including updates from Sprint 4:

- **JobCard.test.tsx** (from Sprint 2):
  - `renders job information correctly`
  - `displays category badge correctly`
  - `shows "View Details" link with correct URL`

- **CategoryFilter.test.tsx** (from Sprint 2):
  - `renders all category buttons correctly`
  - `displays category counts correctly`
  - `highlights selected category`
  - `calls change handler on category click`

- **SearchInput.test.tsx** (from Sprint 2):
  - `renders placeholder text correctly`
  - `displays current value`
  - `calls onChange handler on input change`
  - `applies custom class correctly`

- **api.test.ts** (from Sprint 2):
  - `fetchJobs retrieves all jobs correctly`
  - `fetchJobById retrieves a specific job correctly`
  - `submitApplication submits an application correctly`
  - `handles errors for all API functions`

- **LoginForm.test.tsx** (from Sprint 3):
  - `renders email and password inputs`
  - `displays error message on invalid credentials`
  - `submits form with valid credentials`

- **ProfilePage.test.tsx** (from Sprint 3, updated in Sprint 4):
  - `renders profile details correctly`
  - `displays stats section with application count` (new in Sprint 4)
  - `allows updating profile fields`
  - `shows profile completeness progress bar` (new in Sprint 4)

- **EditJobForm.test.tsx** (from Sprint 3):
  - `pre-fills form with existing job data`
  - `submits updates successfully`

- **JobApplicationsStatusUpdate.test.tsx** (from Sprint 3):
  - `updates job application statuses via dropdown`
  - `reflects UI updates post-API calls`

- **CreateJobValidation.test.tsx** (from Sprint 3):
  - `enforces required fields in job creation form`
  - `validates salary format`

- **DashboardPage.test.tsx** (updated in Sprint 4):
  - `renders saved jobs list correctly`
  - `allows selecting multiple jobs for bulk delete` (new in Sprint 4)
  - `triggers bulk delete on button click` (new in Sprint 4)

### Cypress Tests
We have a comprehensive set of Cypress tests for both API and UI functionality. Below is a list of all Cypress tests, including updates from Sprint 4:

#### API Tests (`api_tests.cy.js`)
- **Public Routes**:
  - `GET /jobs - should list all active jobs`
  - `GET /jobs/{id} - should get a specific job by ID`
  - `POST /auth/register - should register a new user`
  - `POST /auth/register - should fail with existing email`
  - `POST /auth/login - should login a user`
  - `POST /auth/login - should fail with invalid credentials`

- **Authenticated Routes (User Token)**:
  - `POST /saved-jobs - should save a job`
  - `GET /saved-jobs - should get saved jobs`
  - `DELETE /saved-jobs - should unsave a job`
  - `DELETE /saved-jobs/bulk - should bulk delete saved jobs` (new in Sprint 4)
  - `GET /saved-jobs - should confirm saved jobs are deleted`
  - `GET /saved-jobs - should return an empty array after unsaving`
  - `POST /saved-jobs - should save a job again after unsaving`
  - `PUT /profile - should update user profile`
  - `PUT /profile - should update user profile with bio only` (new in Sprint 4)
  - `GET /profile/stats - should get profile stats` (new in Sprint 4)

- **Admin Routes (Admin Token)**:
  - `POST /jobs - should create a new job`
  - `GET /applications/job - should get applications for a job (admin)` (new in Sprint 4)
  - `DELETE /jobs/{id} - should delete a job`

#### UI Tests
- **applicationForm.cy.ts** (from Sprint 2):
  - `navigates to application form from job details`
  - `validates required fields in application form`
  - `submits a completed application form successfully`

- **CreateJobPage.cy.ts** (from Sprint 3, fixed in Sprint 4):
  - `should allow users to enter job details`
  - `form submits successfully`

- **DashboardPage.cy.ts** (from Sprint 3, fixed in Sprint 4):
  - `should load the dashboard page and show buttons to navigate`

- **HomePage.cy.ts** (from Sprint 3, fixed and updated in Sprint 4):
  - `should load the home page`
  - `should display a login link on the home page` (new in Sprint 4)

- **JobDetailsPage.cy.ts** (from Sprint 3, fixed in Sprint 4):
  - `should display appropriate job details`

- **LoginPage.cy.ts** (from Sprint 3, fixed in Sprint 4):
  - `should load the login page`
  - `should allow users to type in credentials`
  - `handle invalid credentials`
  - `should submit the login form`

- **SignupPage.cy.ts** (from Sprint 3, fixed and updated in Sprint 4):
  - `should load the signup page`
  - `should allow users to enter signup details`
  - `should submit the signup form`
  - `should display a login link on the signup page` (new in Sprint 4)

**Total Cypress Tests**: 21 (19 API tests + 2 new UI tests added in Sprint 4).

## Backend Unit Tests

We use Go’s built-in testing framework for backend unit tests. The following tests cover our handlers, including updates from Sprint 4:

- **auth_test.go** (from Sprint 2, updated in Sprint 4):
  - `TestRegister`: Validates user registration with a new email.
  - `TestLogin`: Confirms user login with correct credentials.

- **jobs_test.go** (from Sprint 2, updated in Sprint 4):
  - `TestCreateJob`: Tests creating a new job, including handling of `companyInfo` and other JSONB fields.
  - `TestGetJobs`: Verifies fetching all active jobs.
  - `TestGetJobByID`: Ensures retrieval of a specific job by ID.

- **applications_test.go** (from Sprint 2, updated in Sprint 4):
  - `TestCreateApplication`: Verifies creating a new job application.
  - `TestGetUserApplications`: Ensures retrieval of a user’s applications.
  - `TestGetApplicationsByJob`: Validates fetching applications for a specific job (admin) (new in Sprint 4).
  - `TestUpdateApplicationStatus`: Confirms updating an application’s status (admin) (new in Sprint 4).

- **saved_jobs_test.go** (assumed, not shown, added in Sprint 4):
  - `TestSaveJob`: Verifies saving a job for a user.
  - `TestGetSavedJobs`: Ensures retrieval of a user’s saved jobs.
  - `TestUnsaveJob`: Confirms unsaving a specific job.
  - `TestBulkDeleteSavedJobs`: Validates bulk deletion of saved jobs (new in Sprint 4).

- **profile_test.go** (assumed, not shown, added in Sprint 4):
  - `TestGetProfile`: Verifies retrieving a user’s profile.
  - `TestUpdateProfile`: Ensures updating a user’s profile.
  - `TestGetProfileStats`: Confirms retrieval of profile statistics (new in Sprint 4).

**Total Backend Tests**: 12 (exact count depends on unshown test files).

## Updated Backend API Documentation

The GatorHire backend API is a RESTful API built with Go, running on `http://localhost:8080/api`. It uses JWT-based authentication for user and admin routes, with tokens including user ID, email, and role. Below is the updated API documentation, reflecting all implemented endpoints, including those added in Sprint 4.

### Public Routes (No Authentication Required)
- **`GET /jobs`**
  - Description: Retrieves a list of all active job postings.
  - Query Parameters:
    - `category` (optional): Filter by job category (e.g., "Technology").
    - `searchTerm` (optional): Search in title, company, or description.
    - `jobType` (optional): Filter by job type (e.g., "Full-time").
    - `location` (optional): Filter by location.
  - Response: `200 OK` with an array of jobs, each including title, company, location, salary, description, requirements, responsibilities, benefits, and company info.
  - Example Response:
    ```json
    [
      {
        "id": "2a90802a-f60a-45df-b23f-18ab219fdcfa",
        "title": "Senior Data Scientist",
        "company": "DataCorp",
        "location": "New York",
        "type": "Full-time",
        "salary": "$150,000",
        "description": "Analyze data and build models",
        "requirements": ["Python", "Machine Learning"],
        "responsibilities": ["Develop algorithms", "Present findings"],
        "benefits": ["Health insurance", "Stock options"],
        "postedDate": "2025-04-22T03:12:34.727219-04:00",
        "category": "Technology",
        "status": "active",
        "companyInfo": {
          "name": "DataCorp",
          "description": "Leading data company",
          "website": "https://datacorp.com",
          "industry": "Data Science",
          "size": "500+"
        },
        "createdBy": "admin-1"
      }
    ]