const { defineConfig } = require('cypress')

module.exports = defineConfig({
  e2e: {
    baseUrl: 'http://localhost:8083/api', // Correct base URL for the backend API
    setupNodeEvents(on, config) {
      // Implement node event listeners here if needed
    },
  },
})