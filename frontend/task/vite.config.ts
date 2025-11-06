import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import federation from '@originjs/vite-plugin-federation'
import path from 'path'

export default defineConfig({
  plugins: [
    react(),
    federation({
      name: 'taskApp',
      filename: 'remoteEntry.js',
      exposes: {
        './TaskApp': './src/TaskApp.tsx',
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
    port: 4003,
    strictPort: true,
    cors: true,
  },
  preview: {
    port: 4003,
  },
  build: { 
    target: 'esnext',
    minify: false,
    cssCodeSplit: false,
  },
})
