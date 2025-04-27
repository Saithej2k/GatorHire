describe('Admin Dashboard Page', () => {
    it('should load the admin dashboard with correct navigations', () => {
      cy.visit('/admin/dashboard', { timeout: 10000 }); // 10s timeout
      cy.wait(5000);
      cy.contains('Admin Dashboard', { timeout: 10000 }).should('be.visible');

    });
  });
  