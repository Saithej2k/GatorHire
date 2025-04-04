describe('Signup Page', () => {
    it('should load the signup page', () => {
      cy.visit('/signup');
      cy.contains('Sign Up').should('be.visible');
    });
  
    it('should allow users to enter signup details', () => {
      cy.visit('/signup');
      // cy.get('input[name="name"]').type('Test User');
      // cy.get('input[name="email"]').type('test@example.com');
      // cy.get('input[name="password"]').type('password123');
    });
  
    it('should submit the signup form', () => {
      cy.visit('/signup');
      // cy.get('input[name="name"]').type('Test User');
      // cy.get('input[name="email"]').type('test@example.com');
      // cy.get('input[name="password"]').type('password123');
      // cy.get('button[type="submit"]').click();
    });
  });
  