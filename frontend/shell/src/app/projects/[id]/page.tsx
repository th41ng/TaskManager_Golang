import { useEffect, useState, lazy, Suspense } from "react"
import { useParams } from "react-router"
import { 
  FolderIcon, 
  CalendarIcon, 
  UserIcon,
  TrendingUpIcon,
  MoreVerticalIcon,
  PencilIcon,
  TrashIcon,
} from "lucide-react"
import { projectsApi, tasksApi } from "@/lib/api"
import type { Project, Task } from "@/types"
import { TaskTable } from "@/components/task-table"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

// Import remote dialog
const EditProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.EditProjectDialog })))

export default function ProjectDetailPage() {
  const { id } = useParams<{ id: string }>()
  const [project, setProject] = useState<Project | null>(null)
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  const [editDialogOpen, setEditDialogOpen] = useState(false)

  const loadProjectData = async () => {
    if (!id) return
    
    setLoading(true)
    try {
      const [projectData, tasksData] = await Promise.all([
        projectsApi.get(Number(id)),
        tasksApi.list()
      ])
      
      setProject(projectData)
      setTasks(tasksData.filter(t => t.project_id === Number(id)))
    } catch (err) {
      console.error('Failed to load project:', err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadProjectData()
  }, [id])

  const handleDelete = async () => {
    if (!project || !confirm(`Delete project "${project.name}"?`)) return
    
    try {
      await projectsApi.delete(project.id)
      window.location.href = '/'
    } catch (err) {
      console.error('Failed to delete project:', err)
      alert('Failed to delete project')
    }
  }

  const handleEditSuccess = () => {
    setEditDialogOpen(false)
    loadProjectData()
  }

  if (loading) {
    return (
      <div className="p-6">
        <div className="animate-pulse space-y-4">
          <div className="h-8 bg-gray-200 rounded w-1/3"></div>
          <div className="h-48 bg-gray-200 rounded"></div>
        </div>
      </div>
    )
  }

  if (!project) {
    return (
      <div className="p-6">
        <div className="text-center">
          <h2 className="text-xl font-semibold mb-2">Project not found</h2>
          <p className="text-muted-foreground">The project you're looking for doesn't exist.</p>
          <Button asChild className="mt-4">
            <a href="/">Back to Dashboard</a>
          </Button>
        </div>
      </div>
    )
  }

  const completedTasks = tasks.filter(t => t.done).length
  const totalTasks = tasks.length
  const progressPercent = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0

  return (
    <div className="flex flex-col gap-6 p-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold flex items-center gap-3">
            <FolderIcon className="size-8" />
            {project.name}
          </h1>
          <p className="text-muted-foreground mt-1">
            Project #{project.id}
          </p>
        </div>
        
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" size="icon">
              <MoreVerticalIcon className="size-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem onClick={() => setEditDialogOpen(true)}>
              <PencilIcon className="mr-2 size-4" />
              Edit Project
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem onClick={handleDelete} className="text-red-600">
              <TrashIcon className="mr-2 size-4" />
              Delete Project
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <TrendingUpIcon className="size-5" />
            Project Overview
          </CardTitle>
          <CardDescription>Progress and statistics</CardDescription>
        </CardHeader>
        <CardContent className="space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="flex items-center gap-3">
              <div className="p-3 rounded-lg bg-blue-50">
                <UserIcon className="size-5 text-blue-600" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Owner ID</p>
                <p className="font-semibold">{project.owner_id}</p>
              </div>
            </div>
            
            {project.created_at && (
              <div className="flex items-center gap-3">
                <div className="p-3 rounded-lg bg-green-50">
                  <CalendarIcon className="size-5 text-green-600" />
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Created</p>
                  <p className="font-semibold">
                    {new Date(project.created_at).toLocaleDateString()}
                  </p>
                </div>
              </div>
            )}
            
            <div className="flex items-center gap-3">
              <div className="p-3 rounded-lg bg-purple-50">
                <FolderIcon className="size-5 text-purple-600" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Total Tasks</p>
                <p className="font-semibold">{totalTasks}</p>
              </div>
            </div>
          </div>

          <div className="space-y-2">
            <div className="flex items-center justify-between text-sm">
              <span className="text-muted-foreground">Completion Progress</span>
              <span className="font-semibold">{progressPercent}%</span>
            </div>
            <div className="relative w-full h-2 bg-secondary rounded-full overflow-hidden">
              <div 
                className="absolute h-full bg-primary transition-all duration-300"
                style={{ width: `${progressPercent}%` }}
              />
            </div>
            <div className="flex items-center justify-between text-xs text-muted-foreground">
              <span>{completedTasks} completed</span>
              <span>{totalTasks - completedTasks} remaining</span>
            </div>
          </div>
        </CardContent>
      </Card>

      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <h2 className="text-xl font-semibold">Project Tasks</h2>
          {/* TODO: Add Create Task button with projectId preset */}
        </div>
        
        <TaskTable 
          projectId={Number(id)} 
          onTaskUpdate={loadProjectData}
        />
      </div>

      {/* Remote Edit Project Dialog */}
      <Suspense fallback={null}>
        {editDialogOpen && project && (
          <EditProjectDialog
            open={editDialogOpen}
            onOpenChange={setEditDialogOpen}
            onSuccess={handleEditSuccess}
            project={project}
          />
        )}
      </Suspense>
    </div>
  )
}
