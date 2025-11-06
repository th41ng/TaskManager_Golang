import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import federation from '@originjs/vite-plugin-federation'
import path from 'path'

export default defineConfig({
  plugins: [
    react(),
    federation({
      name: 'projectApp',
      filename: 'remoteEntry.js',
      exposes: {
        './ProjectApp': './src/ProjectApp.tsx',
      },
      shared: ['react', 'react-dom', 'react-router-dom']
    }),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: { 
    port: 4002,
    strictPort: true,
    cors: true,
  },
  preview: {
    port: 4002,
  },
  build: { 
    target: 'esnext',
    minify: false,
    cssCodeSplit: false,
  },
})
