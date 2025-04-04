describe('Login Page', () => {
    it('should load the login page', () => {
      cy.visit('/login');
      cy.contains('Login').should('be.visible');
    });
  
    it('should allow users to type in credentials', () => {
      cy.visit('/login');
      cy.get('input[name="email"]').type('test@example.com');
      cy.get('input[name="password"]').type('password123');
    });
    it('handle invalid credentials', () => {
      cy.visit('/login');
    });
  
    it('should submit the login form', () => {
      cy.visit('/login');
      cy.get('input[name="email"]').type('test@example.com');
      cy.get('input[name="password"]').type('password123');
      cy.get('button[type="submit"]').click();
    });
  });
  