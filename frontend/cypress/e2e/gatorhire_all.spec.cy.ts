describe('GatorHire Simplified Tests', () => {
  // Mock API responses for all tests
  beforeEach(() => {
    // Mock login API (aligned with backend endpoint)
    cy.intercept('POST', '/auth/login', {
      statusCode: 200,
      body: { success: true, token: 'mock-token', user: { role: 'user', fullName: 'Test User' } },
    }).as('loginRequest');

    // Mock admin login API (aligned with backend endpoint)
    cy.intercept('POST', '/auth/admin/login', {
      statusCode: 200,
      body: { success: true, token: 'mock-admin-token', user: { role: 'admin', fullName: 'Admin User' } },
    }).as('adminLoginRequest');

    // Mock signup API (aligned with backend endpoint)
    cy.intercept('POST', '/auth/register', {
      statusCode: 201,
      body: { success: true },
    }).as('signupRequest');

    // Mock job fetch API for JobDetailsPage, JobListingsPage, and ApplicationPage
    cy.intercept('GET', '/jobs/*', {
      statusCode: 200,
      body: {
        id: '1',
        title: 'Test Job',
        company: 'Test Corp',
        location: 'Test City',
        type: 'Full-time',
        salary: '$100k',
        description: 'Test Description',
        requirements: ['Skill 1'],
        responsibilities: ['Task 1'],
        benefits: ['Benefit 1'],
        category: 'Technology',
        companyInfo: { name: 'Test Corp', description: 'A test company' },
        postedDate: '2025-01-01',
      },
    }).as('fetchJob');

    // Mock job creation API
    cy.intercept('POST', '/jobs', {
      statusCode: 201,
      body: { success: true },
    }).as('createJob');

    // Mock user applications API for DashboardPage
    cy.intercept('GET', '/applications/user', {
      statusCode: 200,
      body: [],
    }).as('fetchUserApplications');

    // Mock saved jobs API for DashboardPage
    cy.intercept('GET', '/saved-jobs', {
      statusCode: 200,
      body: [],
    }).as('fetchSavedJobs');

    // Mock recommended jobs API for DashboardPage and JobListingsPage
    cy.intercept('GET', '/jobs', {
      statusCode: 200,
      body: [],
    }).as('fetchJobs');

    // Mock admin jobs API for AdminDashboardPage
    cy.intercept('GET', '/admin/jobs', {
      statusCode: 200,
      body: [],
    }).as('fetchAdminJobs');

    // Mock application submission API for ApplicationPage
    cy.intercept('POST', '/applications/*', {
      statusCode: 201,
      body: { success: true },
    }).as('submitApplication');
  });

  // Login Page Tests
  it('should load the login page', () => {
    cy.visit('/login');
    cy.contains('Log in to GatorHire').should('be.visible');
  });

  it('should allow users to type in login credentials', () => {
    cy.visit('/login');
    cy.get('#email').type('user@example.com');
    cy.get('#password').type('Password123!');
    cy.get('#email').should('have.value', 'user@example.com');
  });

  // Signup Page Tests
  it('should load the signup page', () => {
    cy.visit('/signup');
    cy.contains('Create your account').should('be.visible');
  });

  it('should allow users to enter signup details', () => {
    cy.visit('/signup');
    cy.get('input[name="fullName"]').type('Test User');
    cy.get('input[name="email"]').type('test@example.com');
    cy.get('input[name="fullName"]').should('have.value', 'Test User');
  });

  // Home Page Tests
  it('should load the home page', () => {
    cy.visit('/');
    cy.contains('GatorHire').should('be.visible');
  });

  // Admin Login Page Tests
  it('should load the admin login page', () => {
    cy.visit('/admin/login');
    cy.contains('Employer/Admin Login').should('be.visible');
  });



  // Not Found Page Test
  it('should display 404 page for invalid route', () => {
    cy.visit('/invalid-route');
    cy.contains('404').should('be.visible');
    cy.contains('Page Not Found').should('be.visible');
  });
});