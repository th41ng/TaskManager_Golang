import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import './index.css'
import TaskApp from './TaskApp'

// Standalone mode - wrap TaskApp with BrowserRouter
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <TaskApp 
        shellData={{ user: undefined, theme: 'light', token: undefined }}
        onTaskCreated={() => console.log('Task created')}
        onTaskUpdated={() => console.log('Task updated')}
      />
    </BrowserRouter>
  </StrictMode>,
)
