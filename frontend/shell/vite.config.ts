import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import federation from '@originjs/vite-plugin-federation'
import path from 'path'

export default defineConfig({
  plugins: [
    react(),
    federation({
      name: 'shell',
      remotes: {
        userApp: 'http://localhost:4001/assets/remoteEntry.js',
      },
      shared: ['react', 'react-dom', 'react-router-dom']
    })
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: { 
    port: 5173,
    strictPort: true,
  },
  preview: {
    port: 5173,
  },
  build: { 
    target: 'esnext',
    minify: false,
    cssCodeSplit: false,
  },
})
