import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { PublicClientApplication } from '@azure/msal-browser'
import { MsalProvider } from '@azure/msal-react'
import { msalConfig } from './authConfig'
import './index.css'
import UserApp from './UserApp'

// Khởi tạo MSAL instance
const msalInstance = new PublicClientApplication(msalConfig)

// Standalone mode - wrap UserApp với BrowserRouter
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <MsalProvider instance={msalInstance}>
      <BrowserRouter>
        <UserApp />
      </BrowserRouter>
    </MsalProvider>
  </StrictMode>,
)
