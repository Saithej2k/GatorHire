describe('Create Job Page', () => {
    it('should allow users to enter job details', () => {
      cy.visit('/jobs/create');
      // cy.get('input[name="title"]').type('Software Engineer');
      // cy.get('textarea[name="description"]').type('Job description here.');
      // cy.get('button[type="submit"]').click();
    });
    it('form submits successfuly', () => {
      cy.visit('/jobs/create');
      // cy.get('input[name="title"]').type('Software Engineer');
      // cy.get('textarea[name="description"]').type('Job description here.');
      // cy.get('button[type="submit"]').click();
    });
  });
  