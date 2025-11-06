import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import './index.css'
import ProjectApp from './ProjectApp'

// Standalone mode - wrap ProjectApp with BrowserRouter
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <ProjectApp 
        shellData={{ user: undefined, theme: 'light', token: undefined }}
        onProjectCreated={() => console.log('Project created')}
        onProjectUpdated={() => console.log('Project updated')}
      />
    </BrowserRouter>
  </StrictMode>,
)
