import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: 'dist'
  },
  server: {
    proxy: {
      '/.well-known': 'http://localhost:8089',
      '/a2a': 'http://localhost:8089',
    }
  }
}) 