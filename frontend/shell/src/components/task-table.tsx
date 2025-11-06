import * as React from "react"
import { useState, useEffect, lazy, Suspense } from "react"
import { CheckIcon, PlusIcon } from "lucide-react"
import { tasksApi, projectsApi } from "@/lib/api"
import type { Task, Project } from "@/types"
import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"

// Import remote dialogs
const CreateTaskDialog = lazy(() => import("taskApp/TaskApp").then(m => ({ default: m.CreateTaskDialog })))
const EditTaskDialog = lazy(() => import("taskApp/TaskApp").then(m => ({ default: m.EditTaskDialog })))

interface TaskTableProps {
  projectId?: number
  onTaskUpdate?: () => void
}

export function TaskTable({ projectId, onTaskUpdate }: TaskTableProps) {
  const [tasks, setTasks] = useState<Task[]>([])
  const [projects, setProjects] = useState<Project[]>([])
  const [loading, setLoading] = useState(true)
  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [selectedTask, setSelectedTask] = useState<Task | null>(null)

  const loadData = async () => {
    setLoading(true)
    try {
      const [tasksData, projectsData] = await Promise.all([
        tasksApi.list(),
        projectsApi.list()
      ])
      
      console.log('Tasks loaded:', tasksData)
      console.log('Projects loaded:', projectsData)
      
      let filteredTasks = tasksData
      if (projectId) {
        filteredTasks = tasksData.filter(t => t.project_id === projectId)
      }
      
      setTasks(filteredTasks)
      setProjects(projectsData)
    } catch (err) {
      console.error('Failed to load data:', err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadData()
  }, [projectId])

  const handleToggleComplete = async (task: Task) => {
    try {
      const updatedTask = await tasksApi.toggleComplete(task)
      setTasks(prev => prev.map(t => 
        t.id === task.id ? updatedTask : t
      ))
      onTaskUpdate?.()
    } catch (err) {
      console.error('Failed to toggle task:', err)
    }
  }

  const handleCreateSuccess = () => {
    setCreateDialogOpen(false)
    loadData()
    onTaskUpdate?.()
  }

  const handleEditSuccess = () => {
    setEditDialogOpen(false)
    setSelectedTask(null)
    loadData()
    onTaskUpdate?.()
  }

  const handleEditTask = (task: Task) => {
    setSelectedTask(task)
    setEditDialogOpen(true)
  }

  const getProjectName = (projectId: number) => {
    const project = projects.find(p => p.id === projectId)
    return project?.name || `Project ${projectId}`
  }

  const getPriorityBadge = (priority: number) => {
    const variants = {
      1: { label: 'Low', className: 'bg-blue-50 text-blue-700 border-blue-200' },
      2: { label: 'Medium', className: 'bg-yellow-50 text-yellow-700 border-yellow-200' },
      3: { label: 'High', className: 'bg-red-50 text-red-700 border-red-200' },
    }
    const variant = variants[priority as keyof typeof variants] || variants[2]
    return (
      <Badge variant="outline" className={variant.className}>
        {variant.label}
      </Badge>
    )
  }

  // Group tasks by project
  const groupedTasks = tasks.reduce((acc, task) => {
    const key = task.project_id
    if (!acc[key]) acc[key] = []
    acc[key].push(task)
    return acc
  }, {} as Record<number, Task[]>)

  if (loading) {
    return (
      <div className="px-4 lg:px-6">
        <div className="rounded-lg border bg-card">
          <div className="p-6 text-center text-muted-foreground">
            Loading tasks...
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="px-4 lg:px-6">
      <div className="rounded-lg border bg-card">
        <div className="flex items-center justify-between p-4 border-b">
          <div>
            <h3 className="text-lg font-semibold">Tasks</h3>
            <p className="text-sm text-muted-foreground">
              {tasks.filter(t => !t.done).length} active, {tasks.filter(t => t.done).length} completed
            </p>
          </div>
          <Button onClick={() => setCreateDialogOpen(true)} size="sm">
            <PlusIcon className="mr-2 h-4 w-4" />
            New Task
          </Button>
        </div>

        {tasks.length === 0 ? (
          <div className="p-8 text-center text-muted-foreground">
            No tasks yet. Create one to get started!
          </div>
        ) : (
          <div>
            {Object.entries(groupedTasks).map(([projId, projectTasks]) => (
              <div key={projId}>
                {!projectId && (
                  <div className="px-4 py-3 bg-muted/30 border-b">
                    <h4 className="font-medium text-sm">{getProjectName(Number(projId))}</h4>
                  </div>
                )}
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead className="w-12"></TableHead>
                      <TableHead>Task</TableHead>
                      {!projectId && <TableHead>Project</TableHead>}
                      <TableHead className="w-32">Priority</TableHead>
                      <TableHead className="w-32">Status</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {projectTasks.map(task => (
                      <TableRow key={task.id} className="cursor-pointer hover:bg-muted/50" onClick={() => handleEditTask(task)}>
                        <TableCell onClick={(e) => e.stopPropagation()}>
                          <Checkbox
                            checked={task.done}
                            onCheckedChange={() => handleToggleComplete(task)}
                          />
                        </TableCell>
                        <TableCell>
                          <span className={task.done ? 'line-through text-muted-foreground' : ''}>
                            {task.title}
                          </span>
                        </TableCell>
                        {!projectId && (
                          <TableCell className="text-sm text-muted-foreground">
                            {getProjectName(task.project_id)}
                          </TableCell>
                        )}
                        <TableCell>
                          {getPriorityBadge(task.priority)}
                        </TableCell>
                        <TableCell>
                          {task.done ? (
                            <Badge variant="outline" className="bg-green-50 text-green-700 border-green-200">
                              <CheckIcon className="mr-1 size-3" />
                              Done
                            </Badge>
                          ) : (
                            <Badge variant="outline">Active</Badge>
                          )}
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Remote Task Dialogs */}
      <Suspense fallback={null}>
        {createDialogOpen && (
          <CreateTaskDialog
            open={createDialogOpen}
            onOpenChange={setCreateDialogOpen}
            onSuccess={handleCreateSuccess}
            projectId={projectId}
          />
        )}
      </Suspense>

      <Suspense fallback={null}>
        {editDialogOpen && selectedTask && (
          <EditTaskDialog
            open={editDialogOpen}
            onOpenChange={setEditDialogOpen}
            onSuccess={handleEditSuccess}
            task={selectedTask}
          />
        )}
      </Suspense>
    </div>
  )
}
