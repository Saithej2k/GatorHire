describe('GatorHire API Tests', () => {
  // Tokens for testing (updated with new tokens)
  const userToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJiZGU2NTA5Ny0yNzcyLTQ5YjctYWVmZC1kM2ViNjhmMDBhNTMiLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImlzcyI6ImdhdG9yaGlyZSIsImV4cCI6MTc0NTM5ODU1MywiaWF0IjoxNzQ1MzEyMTUzfQ.xg3H-w4-9cOGaVbOHpRQIbhe-e2111_znpICQxNf7q0';
  const adminToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJhZG1pbi0xIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInJvbGUiOiJhZG1pbiIsImlzcyI6ImdhdG9yaGlyZSIsImV4cCI6MTc0NTM5ODU5MCwiaWF0IjoxNzQ1MzEyMTkwfQ.KBUoMg66LqYpxz_V4p_gpTNtS6T_x7UWAchmqTmMPs0';
  const jobId = '2a90802a-f60a-45df-b23f-18ab219fdcfa'; // Existing job ID
  let newJobId; // Will be set after creating a new job

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
          email: `newuser${Date.now()}@example.com`,
          password: 'newpass',
          fullName: 'New User',
        },
      }).then((response) => {
        expect(response.status).to.eq(201);
        expect(response.body).to.have.property('success', true);
        expect(response.body.user).to.have.property('email');
        expect(response.body).to.have.property('token');
      });
    });

    it('POST /auth/register - should fail with existing email', () => {
      cy.request({
        method: 'POST',
        url: '/auth/register',
        body: {
          email: 'test@example.com',
          password: 'testpass',
          fullName: 'Test User',
        },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.eq(409);
        expect(response.body).to.have.property('error', 'Email already in use');
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

    it('POST /auth/login - should fail with invalid credentials', () => {
      cy.request({
        method: 'POST',
        url: '/auth/login',
        body: {
          email: 'test@example.com',
          password: 'wrongpass',
        },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.eq(401);
        expect(response.body).to.have.property('error', 'Invalid email or password');
      });
    });
  });

  // Authenticated Routes
  describe('Authenticated Routes (User Token)', () => {
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

    it('DELETE /saved-jobs - should unsave a job', () => {
      cy.request({
        method: 'DELETE',
        url: '/saved-jobs',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
        qs: {
          jobId: jobId,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.have.property('success', true);
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

    it('GET /saved-jobs - should return an empty array after unsaving', () => {
      cy.request({
        method: 'GET',
        url: '/saved-jobs',
        headers: {
          Authorization: `Bearer ${userToken}`,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.be.an('array');
        expect(response.body).to.have.length(0);
      });
    });

    it('POST /saved-jobs - should save a job again after unsaving', () => {
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
        newJobId = response.body.id;
      });
    });

    it('GET /applications/job - should get applications for a job (admin)', () => {
      cy.request({
        method: 'GET',
        url: '/applications/job',
        headers: {
          Authorization: `Bearer ${adminToken}`,
        },
        qs: {
          jobId: jobId,
        },
      }).then((response) => {
        expect(response.status).to.eq(200);
        expect(response.body).to.be.an('array');
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