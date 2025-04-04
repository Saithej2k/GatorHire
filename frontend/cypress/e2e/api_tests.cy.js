describe('GatorHire API Tests', () => {
  // Tokens for testing (updated with new tokens)
  const userToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ1c2VyLTEiLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImlzcyI6ImdhdG9yaGlyZSIsImV4cCI6MTc0MzY2MzY2NSwiaWF0IjoxNzQzNTc3MjY1fQ.ZVP3FOin4Ey5y533-JT_eu4OXi0gmLdtGxmUjbN-jOE'; // Replace with the new user token
  const adminToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJhZG1pbi0xIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiIsImlzcyI6ImdhdG9yaGlyZSIsImV4cCI6MTc0MzY2MzY3MCwiaWF0IjoxNzQzNTc3MjcwfQ.kz71ky1jpCoWwvnYuqldhQYbmLFwfu8pV3gioGXQ55s'; // Replace with the new admin token
  const jobId = '2a90802a-f60a-45df-b23f-18ab219fdcfa'; // Existing job ID

  // Public Routes
  describe('Public Routes', () => {
    it('GET /jobs - should list all active jobs', () => {
      cy.request({
        method: 'GET',
        url: '/jobs',
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.be.an('array');
        expect(response.body).to.have.length.greaterThan(0);
        expect(response.body[0]).to.have.property('id');
        expect(response.body[0]).to.have.property('title');
      });
    });

    it('GET /jobs/{id} - should get a specific job by ID', () => {
      cy.request({
        method: 'GET',
        url: `/jobs/${jobId}`,
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('id', jobId);
        expect(response.body).to.have.property('title', 'Senior Data Scientist');
        expect(response.body).to.have.property('company', 'DataCorp');
      });
    });

    it('POST /auth/register - should register a new user', () => {
      cy.request({
        method: 'POST',
        url: '/auth/register',
        body: {
          email: `newuser${Date.now()}@example.com`, // Unique email
          password: 'newpass',
          fullName: 'New User',
        },
      }).then((response) => {
        expect(response.status).to.eq(200); // Fixed to expect 200
        expect(response.body).to.have.property('success', true);
        expect(response.body.user).to.have.property('email');
        expect(response.body).to.have.property('token');
      });
    });

    it('POST /auth/login - should login a user', () => {
      cy.request({
        method: 'POST',
        url: '/auth/login',
        body: {
          email: 'test@example.com',
          password: 'testpass',
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
        expect(response.body.user).to.have.property('email', 'test@example.com');
        expect(response.body).to.have.property('token');
      });
    });
  });

  // Authenticated Routes
  describe('Authenticated Routes (User Token)', () => {
    it('GET /applications/user - should get user applications', () => {
      cy.request({
        method: 'GET',
        url: '/applications/user',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.be.an('array');
      });
    });

    it('POST /saved-jobs - should save a job', () => {
      cy.request({
        method: 'POST',
        url: '/saved-jobs',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
        body: {
          jobId: jobId,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
      });
    });

    it('GET /saved-jobs - should get saved jobs', () => {
      cy.request({
        method: 'GET',
        url: '/saved-jobs',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.be.an('array');
        expect(response.body).to.have.length.greaterThan(0);
        expect(response.body[0]).to.have.property('id', jobId);
      });
    });

    it('DELETE /saved-jobs/bulk - should bulk delete saved jobs', () => {
      cy.request({
        method: 'DELETE',
        url: '/saved-jobs/bulk',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
        body: {
          jobIds: [jobId],
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
      });
    });

    it('GET /saved-jobs - should confirm saved jobs are deleted', () => {
      cy.request({
        method: 'GET',
        url: '/saved-jobs',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
      });
    });

    it('GET /profile - should get user profile', () => {
      cy.request({
        method: 'GET',
        url: '/profile',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('id', 'user-1');
        expect(response.body).to.have.property('email', 'test@example.com');
      });
    });

    it('PUT /profile - should update user profile', () => {
      cy.request({
        method: 'PUT',
        url: '/profile',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
        body: {
          fullName: 'Updated User',
          title: 'Developer',
          location: 'New York',
          bio: 'I am a developer',
          skills: ['Python', 'SQL'],
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('fullName', 'Updated User');
        expect(response.body).to.have.property('title', 'Developer');
        expect(response.body.skills).to.deep.equal(['Python', 'SQL']);
      });
    });

    it('GET /profile/stats - should get profile stats', () => {
      cy.request({
        method: 'GET',
        url: '/profile/stats',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('applicationsCount');
        expect(response.body).to.have.property('savedJobsCount');
        expect(response.body).to.have.property('profileCompleteness');
      });
    });
  });

  // Admin Routes
  describe('Admin Routes (Admin Token)', () => {
    let newJobId; // Declare at the top of the describe block

    it('POST /jobs - should create a new job', () => {
      cy.request({
        method: 'POST',
        url: '/jobs',
        headers: {
          Authorization: `Bearer ${adminToken}`,
        },
        body: {
          title: 'Software Engineer',
          company: 'TechCorp',
          location: 'San Francisco',
          type: 'Full-time',
          salary: '$140,000',
          description: 'Develop software',
          requirements: ['Go', 'SQL'],
          category: 'Technology',
        },
      }).then((response) => {
        expect(response.status).to.eq(201);
        expect(response.body).to.have.property('id');
        expect(response.body).to.have.property('title', 'Software Engineer');
        newJobId = response.body.id; // Capture the new job ID
      });
    });

    it('PUT /jobs/{id} - should update a job', () => {
      cy.request({
        method: 'PUT',
        url: `/jobs/${newJobId}`, // Use the newly created job ID
        headers: {
          Authorization: `Bearer ${adminToken}`,
        },
        body: {
          title: 'Senior Software Engineer',
          company: 'TechCorp',
          location: 'San Francisco',
          type: 'Full-time',
          salary: '$150,000',
          description: 'Updated description',
          requirements: ['Go', 'SQL'],
          category: 'Technology',
          status: 'active',
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
        expect(response.body).to.have.property('message', 'Job updated successfully');
        expect(response.body).to.have.property('jobId', newJobId);
      });
    });

    it('DELETE /jobs/{id} - should delete a job', () => {
      cy.request({
        method: 'DELETE',
        url: `/jobs/${newJobId}`,
        headers: {
          Authorization: `Bearer ${adminToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
      });
    });
  });
});