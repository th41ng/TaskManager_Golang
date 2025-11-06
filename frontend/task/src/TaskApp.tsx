import { useState, useEffect } from 'react'
import CreateTaskDialog from './components/CreateTaskDialog'
import EditTaskDialog from './components/EditTaskDialog'

// ============================================================================
// EXPORT DIALOG COMPONENTS - Để Shell có thể dùng riêng lẻ
// ============================================================================
export { CreateTaskDialog, EditTaskDialog }

// ============================================================================
// INTERFACE - Props từ Shell truyền xuống
// ============================================================================
interface TaskAppProps {
  // Dữ liệu từ Shell (user info, theme, etc.)
  shellData?: {
    user?: any
    theme?: string
    token?: string
  }
  projectId?: number
  // Callbacks để gửi events về Shell
  onTaskCreated?: (task: any) => void
  onTaskUpdated?: (task: any) => void
  onNavigate?: (path: string) => void
}

// ============================================================================
// TASK APP - CONTENT ONLY COMPONENT
// ============================================================================
// Component này CHỈ render nội dung, KHÔNG có:
// - <BrowserRouter> (Shell đã có)
// - <Routes> (Shell quản lý routing)
// - Layout/Sidebar (Shell cung cấp)
//
// Chỉ hiển thị phần nội dung task management trong vùng trống của Shell
export default function TaskApp({ shellData, projectId, onTaskCreated, onTaskUpdated }: TaskAppProps) {
  const [tasks, setTasks] = useState<any[]>([])
  const [loading, setLoading] = useState(true)
  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [selectedTask, setSelectedTask] = useState<any>(null)

  // Load tasks from API
  useEffect(() => {
    loadTasks()
    // If projectId changes, ensure UI reflects it (e.g., open create dialog with project preselected)
  }, [projectId])

  const loadTasks = async () => {
    try {
      setLoading(true)
      const token = shellData?.token || localStorage.getItem('token')
      // If projectId provided, filter tasks by project
      const url = projectId ? `http://172.21.223.107:8080/api/tasks?project_id=${projectId}` : 'http://172.21.223.107:8080/api/tasks'
      const response = await fetch(url, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      })
      const data = await response.json()
      setTasks(data.tasks || [])
    } catch (error) {
      console.error('Failed to load tasks:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleCreateSuccess = (task: any) => {
    setCreateDialogOpen(false)
    loadTasks()
    onTaskCreated?.(task) // Notify Shell
  }

  const handleEditSuccess = (task: any) => {
    setEditDialogOpen(false)
    loadTasks()
    onTaskUpdated?.(task) // Notify Shell
  }

  const handleTaskClick = (task: any) => {
    setSelectedTask(task)
    setEditDialogOpen(true)
  }

  if (loading) {
    return (
      <div id="task-app-root" className="flex items-center justify-center h-64">
        <div className="text-gray-500">Loading tasks...</div>
      </div>
    )
  }

  return (
    <div id="task-app-root" className="space-y-6">
      {/* Header Section */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
            Tasks Management
          </h1>
          <p className="text-gray-600 dark:text-gray-400 mt-2">
            Manage all your tasks in one place
          </p>
        </div>
        <button
          onClick={() => setCreateDialogOpen(true)}
          className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
        >
          + New Task
        </button>
      </div>

      {/* Task List */}
      <div className="bg-white dark:bg-gray-800 rounded-lg shadow">
        <div className="p-6">
          <h2 className="text-xl font-semibold mb-4 text-gray-900 dark:text-white">
            Task List ({tasks.length})
          </h2>
          
          {tasks.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500 dark:text-gray-400">No tasks yet</p>
              <button
                onClick={() => setCreateDialogOpen(true)}
                className="mt-4 text-blue-600 hover:text-blue-700"
              >
                Create your first task
              </button>
            </div>
          ) : (
            <div className="space-y-2">
              {tasks.map((task) => (
                <div
                  key={task.id}
                  onClick={() => handleTaskClick(task)}
                  className="p-4 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 cursor-pointer transition-colors"
                >
                  <div className="flex items-center justify-between">
                    <div className="flex-1">
                      <h3 className="font-medium text-gray-900 dark:text-white">
                        {task.title}
                      </h3>
                      <div className="flex items-center gap-2 mt-1">
                        <span className={`text-xs px-2 py-1 rounded ${
                          task.priority === 1 ? 'bg-red-100 text-red-700' :
                          task.priority === 2 ? 'bg-yellow-100 text-yellow-700' :
                          'bg-green-100 text-green-700'
                        }`}>
                          Priority {task.priority}
                        </span>
                        {task.done && (
                          <span className="text-xs px-2 py-1 rounded bg-green-100 text-green-700">
                            ✓ Completed
                          </span>
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Dialogs */}
      {createDialogOpen && (
        <CreateTaskDialog
          open={createDialogOpen}
          onOpenChange={setCreateDialogOpen}
          projectId={projectId}
          onSuccess={handleCreateSuccess}
        />
      )}

      {editDialogOpen && selectedTask && (
        <EditTaskDialog
          open={editDialogOpen}
          onOpenChange={setEditDialogOpen}
          onSuccess={handleEditSuccess}
          task={selectedTask}
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
