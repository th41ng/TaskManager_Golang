// API Service Layer for TaskManager
import type { Project, Task, CreateProjectRequest, CreateTaskRequest, UpdateTaskRequest } from '@/types'

const API_BASE_URL = (import.meta.env.VITE_API_BASE as string) || 'http://172.21.223.107:8080/api'

// Helper function to get auth headers
function getAuthHeaders(): HeadersInit {
  const token = localStorage.getItem('token')
  return {
    'Content-Type': 'application/json',
    ...(token && { 'Authorization': `Bearer ${token}` })
  }
}

// Helper function to handle API errors
async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const error = await response.json().catch(() => ({ message: 'Unknown error' }))
    throw new Error(error.message || `HTTP Error: ${response.status}`)
  }
  return response.json()
}

// Projects API
export const projectsApi = {
  async list(): Promise<Project[]> {
    const response = await fetch(`${API_BASE_URL}/projects`, {
      headers: getAuthHeaders()
    })
    const data = await handleResponse<{ projects: Project[] }>(response)
    return data.projects || []
  },

  async get(id: number): Promise<Project> {
    const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
      headers: getAuthHeaders()
    })
    return handleResponse<Project>(response)
  },

  async create(data: CreateProjectRequest): Promise<Project> {
    const response = await fetch(`${API_BASE_URL}/projects`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify(data),
    })
    return handleResponse<Project>(response)
  },

  async update(id: number, data: Partial<CreateProjectRequest>): Promise<Project> {
    const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify({ id, ...data }),
    })
    return handleResponse<Project>(response)
  },

  async delete(id: number): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
      method: 'DELETE',
      headers: getAuthHeaders(),
    })
    await handleResponse<{ success: boolean }>(response)
  },
}

// Tasks API
export const tasksApi = {
  async list(): Promise<Task[]> {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      headers: getAuthHeaders()
    })
    const data = await handleResponse<{ tasks: Task[] }>(response)
    return data.tasks || []
  },

  async get(id: number): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      headers: getAuthHeaders()
    })
    return handleResponse<Task>(response)
  },

  async create(data: CreateTaskRequest): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify(data),
    })
    return handleResponse<Task>(response)
  },

  async update(id: number, data: Partial<UpdateTaskRequest>): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify({ id, ...data }),
    })
    return handleResponse<Task>(response)
  },

  async toggleComplete(task: Task): Promise<Task> {
    const response = await fetch(`${API_BASE_URL}/tasks/${task.id}`, {
      method: 'PUT',
      headers: getAuthHeaders(),
      body: JSON.stringify({ 
        id: task.id, 
        title: task.title,
        done: !task.done,
        priority: task.priority,
        project_id: task.project_id
      }),
    })
    return handleResponse<Task>(response)
  },

  async delete(id: number): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
      method: 'DELETE',
      headers: getAuthHeaders(),
    })
    await handleResponse<{ success: boolean }>(response)
  },
}
