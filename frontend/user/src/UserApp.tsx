import { Routes, Route } from "react-router-dom"
import { PublicClientApplication } from '@azure/msal-browser'
import { MsalProvider } from '@azure/msal-react'
import { msalConfig } from './authConfig'
import LoginPage from "./app/login/page"

// Khởi tạo MSAL instance
const msalInstance = new PublicClientApplication(msalConfig)

// Component này được export cho Module Federation
// Không có BrowserRouter vì shell đã có rồi
function UserApp() {
  return (
    <MsalProvider instance={msalInstance}>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/login" element={<LoginPage />} />
      </Routes>
    </MsalProvider>
  )
}

export default UserApp
