import { BrowserRouter, Routes, Route } from "react-router-dom"
import { lazy, Suspense } from "react"
import Page from "@/app/dashboard/page"

// Import UserApp từ remote
const UserApp = lazy(() => import("userApp/UserApp"))

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Page />} />
        <Route 
          path="/login" 
          element={
            <Suspense fallback={<div>Loading...</div>}>
              <UserApp />
            </Suspense>
          } 
        />
        <Route 
          path="/users/*" 
          element={
            <Suspense fallback={<div>Loading...</div>}>
              <UserApp />
            </Suspense>
          } 
        />
      </Routes>
    </BrowserRouter>
  )
}
