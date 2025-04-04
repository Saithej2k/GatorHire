describe('Login Page', () => {
  it('should load the login page', () => {
    cy.visit('/login');
    cy.contains('Log in to GatorHire').should('be.visible');
    cy.get('#email').should('be.visible');
    cy.get('#password').should('be.visible');
    cy.get('button[type="submit"]').should('contain', 'Sign in');
  });

  it('should allow users to type in credentials', () => {
    cy.visit('/login');
    cy.get('#email').type('user@example.com');
    cy.get('#password').type('Password123!');
    cy.get('#email').should('have.value', 'user@example.com');
    cy.get('#password').should('have.value', 'Password123!');
  });

  it('should handle invalid credentials', () => {
    cy.visit('/login');
    cy.get('#email').type('wrong@example.com');
    cy.get('#password').type('wrongpass');
    // Mock the API response to simulate a failed login
    cy.intercept('POST', '/api/auth/login', {
      statusCode: 401,
      body: { success: false },
    }).as('loginFail');
    cy.get('button[type="submit"]').click();
    cy.wait('@loginFail');
    cy.contains('Invalid email or password').should('be.visible');
  });
});