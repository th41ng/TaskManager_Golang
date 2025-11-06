import * as React from 'react'
import { useState, useEffect } from 'react'
import { API_BASE_URL, getAuthHeaders } from '../lib/api'

interface CreateTaskDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  projectId?: number
  onSuccess?: (task: any) => void
}

export default function CreateTaskDialog({ 
  open, 
  onOpenChange,
  projectId,
  onSuccess 
}: CreateTaskDialogProps) {
  const [title, setTitle] = useState('')
  const [selectedProjectId, setSelectedProjectId] = useState<number>(projectId || 0)
  const [priority, setPriority] = useState(2) // Default: Medium
  const [projects, setProjects] = useState<any[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    if (projectId) {
      setSelectedProjectId(projectId)
    }
  }, [projectId])

  useEffect(() => {
    if (open && !projectId) {
      // Fetch projects for dropdown
      fetch(`${API_BASE_URL}/projects`, { headers: getAuthHeaders() })
        .then(res => res.json())
        .then(data => {
          console.log('Fetched projects:', data)
          setProjects(data.projects || [])
        })
        .catch(err => console.error('Failed to fetch projects', err))
    }
  }, [open, projectId])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')

    try {
      const response = await fetch(`${API_BASE_URL}/tasks`, {
        method: 'POST',
        headers: getAuthHeaders(),
        body: JSON.stringify({ 
          title, 
          done: false, 
          priority, 
          project_id: selectedProjectId 
        }),
      })

      if (!response.ok) {
        throw new Error('Failed to create task')
      }

      const task = await response.json()
      onSuccess?.(task)
      onOpenChange(false)
      setTitle('')
      setPriority(2)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div
      className={`${open ? 'absolute' : 'hidden'} inset-0 z-50 flex items-center justify-center p-4`}
      onClick={() => onOpenChange(false)}
    >
      {/* Backdrop (scoped, absolute inside #task-app-root) */}
      <div className="task-backdrop" />

      {/* Dialog (scoped) */}
      <div
        className="task-dialog"
        onClick={(e) => e.stopPropagation()}
      >
        {/* Header */}
  <div className="flex items-center justify-between p-6 border-b border-gray-200 dark:border-gray-700">
          <div>
            <h2 className="text-2xl font-bold text-gray-900 dark:text-white">Create New Task</h2>
            <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">Add a new task to your project</p>
          </div>
          <button
            type="button"
            onClick={() => onOpenChange(false)}
            className="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300 transition-colors rounded-full p-2 hover:bg-gray-100 dark:hover:bg-gray-700"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
  <form onSubmit={handleSubmit} className="p-6 space-y-5">
          {/* Task Title */}
          <div>
            <label htmlFor="title" className="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
              Task Title <span className="text-red-500">*</span>
            </label>
            <input
              id="title"
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
              className="task-input"
              placeholder="e.g., Fix login bug"
            />
          </div>

          {/* Project Selector */}
          {!projectId && (
            <div>
              <label htmlFor="project" className="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                📁 Project <span className="text-red-500">*</span>
              </label>
              <select
                id="project"
                value={selectedProjectId}
                onChange={(e) => setSelectedProjectId(Number(e.target.value))}
                required
                className="task-select"
              >
                <option value="">Select a project</option>
                {projects.map(project => (
                  <option key={project.id} value={project.id}>
                    {project.name}
                  </option>
                ))}
              </select>
            </div>
          )}

          {/* Priority */}
          <div>
            <label className="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3">
              Priority Level
            </label>
            <div className="grid grid-cols-3 gap-3">
              {[
                { value: 1, label: 'Low', color: 'blue' },
                { value: 2, label: 'Medium', color: 'yellow' },
                { value: 3, label: 'High', color: 'red' }
              ].map(({ value, label, color }) => (
                <button
                  key={value}
                  type="button"
                  onClick={() => setPriority(value)}
                  className={`
                    px-4 py-3 rounded-lg border-2 font-medium transition-all
                    ${priority === value
                      ? `bg-${color}-50 border-${color}-500 text-${color}-700 dark:bg-${color}-900/30 dark:border-${color}-500 dark:text-${color}-400`
                      : 'border-gray-200 dark:border-gray-600 text-gray-600 dark:text-gray-400 hover:border-gray-300 dark:hover:border-gray-500'
                    }
                  `}
                >
                  {label}
                </button>
              ))}
            </div>
          </div>

          {/* Error Message */}
          {error && (
            <div className="flex items-start gap-3 p-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg">
              <svg className="w-5 h-5 text-red-600 dark:text-red-400 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
              </svg>
              <p className="text-sm text-red-700 dark:text-red-300">{error}</p>
            </div>
          )}

          {/* Actions */}
          <div className="flex justify-end gap-3 pt-4 border-t border-gray-200 dark:border-gray-700">
            <button
              type="button"
              onClick={() => onOpenChange(false)}
              className="px-6 py-2.5 border border-gray-300 dark:border-gray-600 rounded-lg font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-all"
              disabled={loading}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="px-6 py-2.5 bg-gradient-to-r from-blue-600 to-blue-700 text-white rounded-lg font-medium hover:from-blue-700 hover:to-blue-800 disabled:opacity-50 disabled:cursor-not-allowed transition-all shadow-lg shadow-blue-500/30"
              disabled={loading}
            >
              {loading ? (
                <span className="flex items-center gap-2">
                  <svg className="animate-spin h-4 w-4" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Creating...
                </span>
              ) : (
                'Create Task'
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}
