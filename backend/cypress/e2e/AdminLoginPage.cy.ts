describe('Admin Login Page', () => {
    it('should allow admin to log in', () => {
      cy.visit('/admin/login');
      cy.get('input[name="email"]').type('admin@example.com');
      cy.get('input[name="password"]').type('adminpass');
      cy.get('button[type="submit"]').click();
    });
  });
  