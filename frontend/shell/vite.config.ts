import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import federation from '@originjs/vite-plugin-federation'
import path from 'path'

export default defineConfig(({ mode }) => {
  // load .env files
  const env = loadEnv(mode, process.cwd(), '')

  // Default to the Nomad production remoteEntry locations (served under /assets).
  // If you need to override for special testing you can still set VITE_REMOTE_*.
  const userRemote = env.VITE_REMOTE_USER || env.VITE_REMOTE_USER_URL || 'http://172.21.223.107:4001/assets/remoteEntry.js'
  const projectRemote = env.VITE_REMOTE_PROJECT || env.VITE_REMOTE_PROJECT_URL || 'http://172.21.223.107:4002/assets/remoteEntry.js'
  const taskRemote = env.VITE_REMOTE_TASK || env.VITE_REMOTE_TASK_URL || 'http://172.21.223.107:4003/assets/remoteEntry.js'

  return {
    plugins: [
      react(),
      federation({
        name: 'shell',
        remotes: {
          userApp: userRemote,
          projectApp: projectRemote,
          taskApp: taskRemote,
        },
        shared: ['react', 'react-dom', 'react-router-dom'],
      }),
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
  }
})
