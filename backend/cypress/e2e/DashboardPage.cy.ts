describe('Dashboard Page', () => {
    it('should load the dashboard page and show buttons to navigate', () => {
      cy.visit('/dashboard');
      cy.contains('Dashboard').should('be.visible');
    });
  });
  