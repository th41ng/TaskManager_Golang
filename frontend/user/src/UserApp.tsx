import { Routes, Route } from "react-router-dom"
import { CookiesProvider } from "react-cookie"
import LoginPage from "./app/login/page"

// Component này được export cho Module Federation
// Không có BrowserRouter vì shell đã có rồi
function UserApp() {
  return (
    <CookiesProvider>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/login" element={<LoginPage />} />
      </Routes>
    </CookiesProvider>
  )
}

export default UserApp
