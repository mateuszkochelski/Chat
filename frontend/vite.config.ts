import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    port: 4173,
    host: true,
  },
  preview: {
    allowedHosts: ['frontend_container'],
    port: 4173,
    host: true,
  },
});
