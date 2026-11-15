import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// The Go API. Override with API_URL, e.g. API_URL=http://localhost:5100 npm run dev
const apiUrl = process.env.API_URL ?? 'http://localhost:5000'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // The app calls /api/...; the dev server forwards it to the Go API without
    // the /api prefix. The browser sees one origin, so the API needs no CORS.
    proxy: {
      '/api': {
        target: apiUrl,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
