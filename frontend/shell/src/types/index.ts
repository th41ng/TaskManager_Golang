// TypeScript interfaces cho Project và Task

export interface Project {
  id: number
  name: string
  owner_id: number
  created_at?: string
  updated_at?: string
}

export interface Task {
  id: number
  title: string
  done: boolean
  priority: number // 1=Low, 2=Medium, 3=High
  project_id: number
  created_at?: string
  updated_at?: string
}

export interface CreateProjectRequest {
  name: string
  owner_id: number
}

export interface UpdateProjectRequest {
  id: number
  name: string
  owner_id: number
}

export interface CreateTaskRequest {
  title: string
  done: boolean
  priority: number
  project_id: number
}

export interface UpdateTaskRequest {
  id: number
  title: string
  done: boolean
  priority: number
  project_id: number
}

export interface ProjectWithTasks extends Project {
  tasks?: Task[]
  completedTasks?: number
  totalTasks?: number
}
