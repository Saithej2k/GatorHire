const { defineConfig } = require('cypress')

module.exports = defineConfig({
  e2e: {
    baseUrl: 'http://localhost:8083/api', // Ensure this is correct
    supportFile: false, // Disable the support file requirement
    setupNodeEvents(on, config) {
      // Implement node event listeners here if needed
    },
  },
})