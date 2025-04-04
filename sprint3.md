# Sprint 3 Summary

## Work Completed
In Sprint 3, our team focused on resolving outstanding issues from Sprint 2 while implementing new features to improve the GatorHire platform. Here’s a detailed breakdown of the work completed:

- **Resolved Uncompleted Issues from Sprint 2**:
  - **Issue #45: Login Validation Bug** - Corrected a flaw in the login form that permitted invalid credentials to bypass validation. The system now displays appropriate error messages for incorrect email or password entries.
  - **Issue #52: Profile Page Loading Error** - Fixed a bug where the user profile page failed to load for some users due to missing data fields. The page now handles such cases smoothly.

- **New Functionality Implemented**:
  - **User Profile Enhancement** - Enabled users to edit and save their profile details (name, email, bio) dynamically using the `fetchUserProfile` and `updateUserProfile` APIs.
  - **Admin Job Management** - Developed the `EditJobPage.tsx` component, allowing admins to modify existing job postings through the `updateJob` API.
  - **Job Application Management** - Built the `JobApplicationsPage.tsx` component, which lets admins view and update job application statuses via the `fetchJobApplications` and `updateApplicationStatus` APIs.

- **Bug Fixes**:
  - **Issue #60: Salary Validation in Job Creation** - Implemented validation to ensure the salary field in job postings is either a number or a valid range (e.g., "50000-60000").
  - **Issue #63: Loading State Improvement** - Enhanced user experience by replacing basic "Loading..." text with a spinner and message on pages like `DashboardPage.tsx` and `EditJobPage.tsx`.

- **Security Enhancements**:
  - Added authentication checks to admin routes (`EditJobPage.tsx`, `JobApplicationsPage.tsx`, `CreateJobPage.tsx`) to restrict access to authenticated admin users only.

## Frontend Unit Tests
The following unit tests were written to validate the frontend components and features developed in Sprint 3:

- **`testLoginFormValidation`**:
  - **Purpose**: Verifies that the login form accurately validates input and shows error messages for invalid credentials.
  - **Coverage**: Tests empty fields, invalid email formats, and incorrect passwords.

- **`testProfilePageLoad`**:
  - **Purpose**: Ensures the user profile page loads correctly and displays user data (name, email, bio).
  - **Coverage**: Includes tests for authenticated and unauthenticated users, with redirection to the login page when appropriate.

- **`testEditJobForm`**:
  - **Purpose**: Confirms that the job editing form in `EditJobPage.tsx` pre-fills with existing job data and submits updates successfully.
  - **Coverage**: Tests form field population, input modifications, and API submission.

- **`testJobApplicationsStatusUpdate`**:
  - **Purpose**: Validates that job application statuses can be updated via the dropdown in `JobApplicationsPage.tsx`.
  - **Coverage**: Tests status changes and UI updates post-API calls.

- **`testCreateJobValidation`**:
  - **Purpose**: Ensures the job creation form in `CreateJobPage.tsx` enforces required fields and valid salary formats.
  - **Coverage**: Tests missing fields and improper salary inputs.

## Backend Unit Tests
The following unit tests were created to verify the backend APIs and services implemented or updated in Sprint 3:

- **`testGetUserProfile`**:
  - **Purpose**: Ensures the `GET /api/profile` endpoint returns accurate user profile data for authenticated users.
  - **Coverage**: Tests successful data retrieval and error handling for unauthenticated requests.

- **`testUpdateUserProfile`**:
  - **Purpose**: Verifies that the `PUT /api/profile` endpoint updates user profile information correctly.
  - **Coverage**: Tests successful updates and validation errors (e.g., missing fields).

- **`testUpdateJob`**:
  - **Purpose**: Confirms that the `PUT /api/jobs/:id` endpoint updates job postings with provided data.
  - **Coverage**: Tests successful updates, unauthorized access attempts, and invalid job IDs.

- **`testFetchJobApplications`**:
  - **Purpose**: Ensures the `GET /api/applications/job?jobId=<id>` endpoint returns the correct list of job applications.
  - **Coverage**: Tests valid and invalid job IDs, plus empty application lists.

- **`testUpdateApplicationStatus`**:
  - **Purpose**: Verifies that the `PUT /api/applications/status` endpoint updates job application statuses correctly.
  - **Coverage**: Tests valid status updates, invalid application IDs, and unauthorized access.

## Updated Backend API Documentation
The following API endpoints were updated or added in Sprint 3, complete with descriptions, request formats, and response structures.

### **GET /api/profile**
- **Description**: Retrieves the authenticated user's profile data.
- **Authentication**: Requires a valid JWT token.
- **Response**:
  ```json
  {
    "id": "user-1",
    "fullName": "John Doe",
    "email": "john@example.com",
    "bio": "Experienced frontend developer."
  }
  ```
- **Error Response**:
  - `401 Unauthorized`: Returned if no valid token is provided.

### **PUT /api/profile**
- **Description**: Updates the authenticated user's profile information.
- **Authentication**: Requires a valid JWT token.
- **Request Body**:
  ```json
  {
    "fullName": "John Doe",
    "email": "john@example.com",
    "bio": "Updated bio."
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Profile updated successfully."
  }
  ```
- **Error Response**:
  - `400 Bad Request`: Returned if required fields are missing.
  - `401 Unauthorized`: Returned if no valid token is provided.

### **PUT /api/jobs/:id**
- **Description**: Updates an existing job posting (admin-only).
- **Authentication**: Requires a valid admin JWT token.
- **Request Body**:
  ```json
  {
    "title": "Updated Job Title",
    "company": "TechCorp",
    "location": "San Francisco",
    "type": "Full-time",
    "salary": "100000-120000",
    "description": "Updated description.",
    "requirements": ["React", "TypeScript"],
    "category": "Technology"
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Job updated successfully."
  }
  ```
- **Error Response**:
  - `400 Bad Request`: Returned if required fields are missing.
  - `401 Unauthorized`: Returned if no valid admin token is provided.
  - `404 Not Found`: Returned if the job ID does not exist.

### **GET /api/applications/job?jobId=<id>**
- **Description**: Retrieves all applications for a specific job (admin-only).
- **Authentication**: Requires a valid admin JWT token.
- **Response**:
  ```json
  [
    {
      "id": "app-1",
      "jobId": "job-1",
      "fullName": "John Doe",
      "email": "john@example.com",
      "status": "pending"
    }
  ]
  ```
- **Error Response**:
  - `401 Unauthorized`: Returned if no valid admin token is provided.
  - `404 Not Found`: Returned if the job ID does not exist.

### **PUT /api/applications/status**
- **Description**: Updates the status of a job application (admin-only).
- **Authentication**: Requires a valid admin JWT token.
- **Request Body**:
  ```json
  {
    "applicationId": "app-1",
    "status": "reviewed"
  }
  ```
- **Response**:
  ```json
  {
    "success": true,
    "message": "Application status updated."
  }
  ```
- **Error Response**:
  - `400 Bad Request`: Returned if the status is invalid.
  - `401 Unauthorized`: Returned if no valid admin token is provided.
  - `404 Not Found`: Returned if the application ID does not exist.