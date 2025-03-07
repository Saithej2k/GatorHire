import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
        // Add error handling for proxy
        configure: (proxy, _options) => {
          proxy.on('error', (err, _req, _res) => {
            // Suppress error output to console
            console.log('API server unavailable, using mock data instead');
          });
        }
      }
    }
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true
  }
});