import { useState, useEffect } from 'react'
import CreateProjectDialog from './components/CreateProjectDialog'
import EditProjectDialog from './components/EditProjectDialog'

// ============================================================================
// EXPORT DIALOG COMPONENTS - Để Shell có thể dùng riêng lẻ
// ============================================================================
export { CreateProjectDialog, EditProjectDialog }

// ============================================================================
// INTERFACE - Props từ Shell truyền xuống
// ============================================================================
interface ProjectAppProps {
  // Dữ liệu từ Shell
  shellData?: {
    user?: any
    theme?: string
    token?: string
  }
  // Callbacks để gửi events về Shell
  onProjectCreated?: (project: any) => void
  onProjectUpdated?: (project: any) => void
  onNavigate?: (path: string) => void
}

// ============================================================================
// PROJECT APP - CONTENT ONLY COMPONENT
// ============================================================================
// Component này CHỈ render nội dung, KHÔNG có:
// - <BrowserRouter> (Shell đã có)
// - <Routes> (Shell quản lý routing)
// - Layout/Sidebar (Shell cung cấp)
//
// Chỉ hiển thị phần nội dung project management trong vùng trống của Shell
export default function ProjectApp({ shellData, onProjectCreated, onProjectUpdated }: ProjectAppProps) {
  const [projects, setProjects] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [selectedProject, setSelectedProject] = useState<any>(null)

  // Load projects from API
  useEffect(() => {
    loadProjects()
  }, [])

  const loadProjects = async () => {
    try {
      setLoading(true)
      const token = shellData?.token || localStorage.getItem('token')
      const response = await fetch('http://172.21.223.107:8080/api/projects', {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      })
      const data = await response.json()
      setProjects(data.projects || [])
    } catch (error) {
      console.error('Failed to load projects:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleCreateSuccess = (project: any) => {
    setCreateDialogOpen(false)
    loadProjects()
    onProjectCreated?.(project) // Notify Shell
  }

  const handleEditSuccess = (project: any) => {
    setEditDialogOpen(false)
    loadProjects()
    onProjectUpdated?.(project) // Notify Shell
  }

  const handleProjectClick = (project: any) => {
    setSelectedProject(project)
    setEditDialogOpen(true)
  }

  if (loading) {
    return (
      <div id="project-app-root" className="flex items-center justify-center h-64">
        <div className="text-gray-500">Loading projects...</div>
      </div>
    )
  }

  return (
    <div id="project-app-root" className="space-y-6">
      {/* Header Section */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
            Projects Management
          </h1>
          <p className="text-gray-600 dark:text-gray-400 mt-2">
            Manage all your projects in one place
          </p>
        </div>
        <button
          onClick={() => setCreateDialogOpen(true)}
          className="px-4 py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 transition-colors"
        >
          + New Project
        </button>
      </div>

      {/* Project Grid */}
      <div className="bg-white dark:bg-gray-800 rounded-lg shadow">
        <div className="p-6">
          <h2 className="text-xl font-semibold mb-4 text-gray-900 dark:text-white">
            Project List ({projects.length})
          </h2>
          
          {projects.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500 dark:text-gray-400">No projects yet</p>
              <button
                onClick={() => setCreateDialogOpen(true)}
                className="mt-4 text-purple-600 hover:text-purple-700"
              >
                Create your first project
              </button>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {projects.map((project) => (
                <div
                  key={project.id}
                  className="p-6 border border-gray-200 dark:border-gray-700 rounded-lg hover:shadow-lg hover:border-purple-300 dark:hover:border-purple-600 transition-all"
                >
                  <div className="flex items-start justify-between mb-3">
                    <div className="w-12 h-12 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center">
                      <span className="text-2xl">📁</span>
                    </div>
                  </div>
                  <h3 className="font-semibold text-lg text-gray-900 dark:text-white mb-2">
                    {project.name}
                  </h3>
                  <p className="text-sm text-gray-600 dark:text-gray-400 line-clamp-2">
                    {project.description || 'No description'}
                  </p>
                  <div className="mt-4 flex gap-3">
                    <button
                      type="button"
                      onClick={() => handleProjectClick(project)}
                      className="project-btn project-btn-secondary"
                    >
                      Edit
                    </button>
                    <button
                      type="button"
                      onClick={() => {
                        // Emit openTaskList so Shell can react and show Task app
                        window.dispatchEvent(new CustomEvent('openTaskList', { detail: project.id }))
                      }}
                      className="project-btn project-btn-primary"
                    >
                      View Tasks
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Dialogs */}
      {createDialogOpen && (
        <CreateProjectDialog
          open={createDialogOpen}
          onOpenChange={setCreateDialogOpen}
          onSuccess={handleCreateSuccess}
        />
      )}

      {editDialogOpen && selectedProject && (
        <EditProjectDialog
          open={editDialogOpen}
          onOpenChange={setEditDialogOpen}
          onSuccess={handleEditSuccess}
          project={selectedProject}
        />
      )}

      {/* Debug Info (Optional) */}
      {shellData && (
        <div className="mt-8 p-4 bg-gray-100 dark:bg-gray-900 rounded-lg text-xs">
          <p className="font-semibold text-gray-700 dark:text-gray-300">
            📡 Data from Shell:
          </p>
          <pre className="mt-2 text-gray-600 dark:text-gray-400">
            {JSON.stringify(shellData, null, 2)}
          </pre>
        </div>
      )}
    </div>
  )
}
