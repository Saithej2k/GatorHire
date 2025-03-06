# Sprint 1: GatorHire Platform Development

## User Stories

### 1. User Authentication (Signup & Login)
As a user, I want to be able to **sign up** for a new account and **log in** with my credentials so that I can access my personalized dashboard and job listings.

### 2. Job Listings Management
As a user, I want to be able to **view all available job listings**, **search for jobs** by title or keyword, and **view detailed information** about each job so that I can find suitable job opportunities.

### 3. Admin Features
As an admin, I want to be able to **create new job listings** in the system so that job seekers can apply to new opportunities.

---

## Planned Issues for Sprint 1

### Frontend Development:
1. **Home Page Layout** - Design the homepage with job categories and a "Get Started" button.
2. **Signup Page** - Implement the signup form for creating new accounts.
3. **Login Page** - Implement the login form for user authentication.
4. **Dashboard Page** - Implement a dashboard page that fetches job listings from the backend.
5. **Navbar** - Implement a dynamic navbar that switches between Login/Signup or Logout based on authentication state.

### Backend Development:
1. **User Authentication (Signup & Login)** - Implement API endpoints to handle user registration and authentication.
2. **Job Listings API (Create, Get, Search, Get by ID)** - Implement API endpoints to create, retrieve, search, and view job listings.

### Database Setup:
1. **PostgreSQL Database Configuration** - Set up and configure the PostgreSQL database for storing user and job data.
2. **Database Models** - Define models for users and jobs in the database.

---

## Completed Issues

### Frontend:
1. **Home Page Layout**: Completed and implemented with dynamic categories.
2. **Signup Page**: Completed the form and integrated it with the backend.
3. **Login Page**: Completed the form and integrated it with the backend.

### Backend:
1. **User Authentication (Signup & Login)**: Implemented signup and login API endpoints with JWT token generation.
2. **Job Listings API**: Completed the job creation, retrieval, and search functionality.
3. **Database Configuration**: Successfully set up PostgreSQL and connected the backend with it.

### Database:
1. **PostgreSQL Database Setup**: Completed the database configuration and verified the connection using `DB.Ping()`.
2. **Database Models**: Successfully defined and implemented the models for users and jobs.

---

## Incomplete Issues and Challenges

### Frontend:
1. **Dashboard Page**: The dashboard page was partially completed but was not fully integrated with the job listing API. The integration with the backend is delayed and will be carried over to the next sprint.

### Backend:
1. **Job Creation and Job Details Fetching**: The CreateJob endpoint is functional, but there were issues in the UI integration for dynamic job listing updates. Further testing and improvements are required.

### Challenges:
1. **API Integration**: There were some challenges while integrating the backend APIs with the frontend, especially with dynamic data fetching and ensuring the frontend updates correctly based on backend responses.
2. **Testing and Debugging**: Extensive testing was required for the authentication system to ensure JWT tokens were properly handled, and some errors in user validation were identified and fixed during testing.

---

## Conclusion

In **Sprint 1**, we successfully completed most of the backend and frontend foundations for **user authentication**, **job listing management**, and **database setup**. The **frontend integration** with the backend is partially complete, and further work will be done in **Sprint 2** to finalize the dashboard and ensure proper API integration.
