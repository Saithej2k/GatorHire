describe('Job Details Page', () => {
    it('should display appropriate job details', () => {
      cy.visit('/jobs/e8754a32-9ca5-476c-8fb6-bd6062ab958b'); // Assuming job ID 1
      cy.contains('Apply Now').should('be.visible');
    });
  });
  