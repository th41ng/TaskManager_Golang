import { BrowserRouter, Routes, Route, useNavigate } from "react-router-dom"
import { lazy, Suspense, useState, useEffect } from "react"
import { ThemeProvider } from "@/components/theme-provider"
import DashboardPage from "@/app/dashboard/page"
import ProjectDetailPage from "@/app/projects/[id]/page"
import Layout from "@/components/layout"

// ============================================================================
// MICRO FRONTEND IMPORTS - Remote Applications
// ============================================================================
// Kiến trúc: Shell owns Layout, Remote owns Content Only
// - Shell: Quản lý routing, layout (sidebar + header), truyền data xuống remote
// - Remote: Chỉ render nội dung, nhận props từ shell, gọi callbacks về shell
//
// Type declarations: src/types/remote-modules.d.ts

const UserApp = lazy(() => import("userApp/UserApp"))
const ProjectApp = lazy(() => import("projectApp/ProjectApp"))
const TaskApp = lazy(() => import("taskApp/TaskApp"))

export default function App() {
  // ============================================================================
  // SHELL STATE - Shared data để truyền xuống remote apps
  // ============================================================================
  const [shellData, setShellData] = useState({
    user: null as any,
    theme: 'light',
    token: localStorage.getItem('token') || '',
  })

  // Load user info khi app khởi động
  useEffect(() => {
    const token = localStorage.getItem('token')
    if (token) {
      // TODO: Fetch user info from API
      setShellData(prev => ({
        ...prev,
        token,
        user: { name: 'Admin User', email: 'admin@taskmanager.com' }
      }))
    }
  }, [])

  // target project id when Shell receives openTaskList event
  const [targetProjectId, setTargetProjectId] = useState<number | null>(null)

  // ============================================================================
  // CALLBACK HANDLERS - Remote apps gọi về Shell
  // ============================================================================
  const handleTaskEvent = (event: string, data?: any) => {
    console.log('📨 Task Event:', event, data)
    // TODO: Handle task events (refresh sidebar, show notification, etc.)
  }

  const handleProjectEvent = (event: string, data?: any) => {
    console.log('📨 Project Event:', event, data)
    // TODO: Handle project events (refresh sidebar, show notification, etc.)
  }

  // Inner component so we can use useNavigate() inside BrowserRouter
  function AppRoutes() {
    const navigate = useNavigate()

    useEffect(() => {
      const handler = (e: any) => {
        const projectId = e?.detail
        setTargetProjectId(projectId || null)
        // Navigate to tasks and include projectId as query param
        navigate(`/tasks${projectId ? `?projectId=${projectId}` : ''}`)
      }
      window.addEventListener('openTaskList', handler as EventListener)
      return () => window.removeEventListener('openTaskList', handler as EventListener)
    }, [navigate])

    return (
      <Routes>
        {/* ====================================================================
            DASHBOARD - Shell Local Page
            ==================================================================== */}
  <Route path="/" element={<DashboardPage />} />
        
        {/* ====================================================================
            LOGIN - Remote Full Page (No Layout)
            ====================================================================
            UserApp tự quản lý layout vì là trang đăng nhập độc lập
        */}
        <Route 
          path="/login/*" 
          element={
            <Suspense fallback={<div>Loading...</div>}>
              <UserApp />
            </Suspense>
          } 
        />
        
        {/* ====================================================================
            USERS - Remote với Layout
            ==================================================================== */}
        <Route 
          path="/users/*" 
          element={
            <Layout breadcrumbItems={[
              { label: "Dashboard", href: "/" },
              { label: "Users", isActive: true }
            ]}>
              <Suspense fallback={<div>Loading...</div>}>
                <UserApp />
              </Suspense>
            </Layout>
          } 
        />

        {/* ====================================================================
            PROJECT DETAIL - Shell Local Page
            ==================================================================== */}
        <Route 
          path="/projects/:id" 
          element={
            <Layout breadcrumbItems={[
              { label: "Dashboard", href: "/" },
              { label: "Projects", href: "/" },
              { label: "Detail", isActive: true }
            ]}>
              <Suspense fallback={<div>Loading Project...</div>}>
                <ProjectDetailPage />
              </Suspense>
            </Layout>
          } 
        />

        {/* ====================================================================
            PROJECTS MANAGEMENT - Remote Content Only
            ====================================================================
            Shell cung cấp:
            - Layout (sidebar + header)
            - Breadcrumb
            - shellData props
            - Callback handlers
            
            ProjectApp nhận:
            - shellData: { user, theme, token }
            - onProjectCreated, onProjectUpdated callbacks
        */}
        <Route 
          path="/projects-app/*" 
          element={
            <Layout breadcrumbItems={[
              { label: "Dashboard", href: "/" },
              { label: "Projects Management", isActive: true }
            ]}>
              <Suspense fallback={<div className="p-6">Loading Projects...</div>}>
                <ProjectApp 
                  shellData={shellData}
                  onProjectCreated={(project) => handleProjectEvent('created', project)}
                  onProjectUpdated={(project) => handleProjectEvent('updated', project)}
                />
              </Suspense>
            </Layout>
          } 
        />

        {/* ====================================================================
            TASKS MANAGEMENT - Remote Content Only
            ====================================================================
            Shell cung cấp:
            - Layout (sidebar + header)
            - Breadcrumb
            - shellData props
            - Callback handlers
            
            TaskApp nhận:
            - shellData: { user, theme, token }
            - onTaskCreated, onTaskUpdated callbacks
        */}
        <Route 
          path="/tasks/*" 
          element={
            <Layout breadcrumbItems={[
              { label: "Dashboard", href: "/" },
              { label: "Tasks Management", isActive: true }
            ]}>
              <Suspense fallback={<div className="p-6">Loading Tasks...</div>}>
                <TaskApp 
                  shellData={shellData}
                  projectId={targetProjectId ?? undefined}
                  onTaskCreated={(task) => handleTaskEvent('created', task)}
                  onTaskUpdated={(task) => handleTaskEvent('updated', task)}
                />
              </Suspense>
            </Layout>
          } 
        />
      </Routes>
    )
  }

  return (
    <div id="shell-root">
      <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
        <BrowserRouter>
          <AppRoutes />
        </BrowserRouter>
      </ThemeProvider>
    </div>
  )
}
