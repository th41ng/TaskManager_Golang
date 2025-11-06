import { TrendingDownIcon, TrendingUpIcon, FolderIcon, CheckSquareIcon, ListTodoIcon, ClockIcon, PlusIcon } from "lucide-react"
import { useEffect, useState, lazy, Suspense } from "react"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { projectsApi, tasksApi } from "@/lib/api"
import type { Project, Task } from "@/types"

// Import remote dialogs
const CreateProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.CreateProjectDialog })))

export function SectionCards() {
  const [projects, setProjects] = useState<Project[]>([])
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  const [createProjectOpen, setCreateProjectOpen] = useState(false)

  useEffect(() => {
    loadData()
  }, [])

  const loadData = () => {
    Promise.all([
      projectsApi.list(),
      tasksApi.list()
    ])
      .then(([projectsData, tasksData]) => {
        console.log('SectionCards - Projects:', projectsData)
        console.log('SectionCards - Tasks:', tasksData)
        setProjects(projectsData)
        setTasks(tasksData)
      })
      .catch(err => console.error('Failed to load data:', err))
      .finally(() => setLoading(false))
  }

  const handleCreateProjectSuccess = () => {
    setCreateProjectOpen(false)
    loadData()
  }

  const completedTasks = tasks.filter(t => t.done).length
  const totalTasks = tasks.length
  const completionRate = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0
  const activeTasks = totalTasks - completedTasks

  if (loading) {
    return (
      <div className="grid grid-cols-1 gap-4 px-4 lg:px-6 @xl/main:grid-cols-2 @5xl/main:grid-cols-4">
        {[1, 2, 3, 4].map(i => (
          <Card key={i} className="animate-pulse">
            <CardHeader>
              <div className="h-4 bg-gray-200 rounded w-24 mb-2"></div>
              <div className="h-8 bg-gray-200 rounded w-16"></div>
            </CardHeader>
          </Card>
        ))}
      </div>
    )
  }

  return (
    <>
      <div className="flex items-center justify-between px-4 lg:px-6 mb-2">
        <div>
          <h2 className="text-2xl font-bold">Dashboard Overview</h2>
          <p className="text-sm text-muted-foreground">Quick stats and metrics</p>
        </div>
        <Button onClick={() => setCreateProjectOpen(true)} size="sm">
          <PlusIcon className="mr-2 h-4 w-4" />
          New Project
        </Button>
      </div>

      <div className="*:data-[slot=card]:shadow-xs @xl/main:grid-cols-2 @5xl/main:grid-cols-4 grid grid-cols-1 gap-4 px-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card lg:px-6">
      <Card className="@container/card">
        <CardHeader className="relative">
          <CardDescription className="flex items-center gap-2">
            <FolderIcon className="size-4" />
            Total Projects
          </CardDescription>
          <CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
            {projects.length}
          </CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium text-muted-foreground">
            Active projects in system
          </div>
        </CardFooter>
      </Card>
      
      <Card className="@container/card">
        <CardHeader className="relative">
          <CardDescription className="flex items-center gap-2">
            <ListTodoIcon className="size-4" />
            Total Tasks
          </CardDescription>
          <CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
            {totalTasks}
          </CardTitle>
          <div className="absolute right-4 top-4">
            <Badge variant="outline" className="flex gap-1 rounded-lg text-xs">
              {completionRate}% done
            </Badge>
          </div>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium text-muted-foreground">
            {completedTasks} completed, {activeTasks} active
          </div>
        </CardFooter>
      </Card>
      
      <Card className="@container/card">
        <CardHeader className="relative">
          <CardDescription className="flex items-center gap-2">
            <CheckSquareIcon className="size-4" />
            Completed Tasks
          </CardDescription>
          <CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
            {completedTasks}
          </CardTitle>
          <div className="absolute right-4 top-4">
            <Badge variant="outline" className="flex gap-1 rounded-lg text-xs">
              <TrendingUpIcon className="size-3" />
              {completionRate}%
            </Badge>
          </div>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium">
            Good progress <TrendingUpIcon className="size-4" />
          </div>
          <div className="text-muted-foreground">Tasks marked as done</div>
        </CardFooter>
      </Card>
      
      <Card className="@container/card">
        <CardHeader className="relative">
          <CardDescription className="flex items-center gap-2">
            <ClockIcon className="size-4" />
            Active Tasks
          </CardDescription>
          <CardTitle className="@[250px]/card:text-3xl text-2xl font-semibold tabular-nums">
            {activeTasks}
          </CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1 text-sm">
          <div className="line-clamp-1 flex gap-2 font-medium text-muted-foreground">
            Pending completion
          </div>
        </CardFooter>
      </Card>
    </div>

    {/* Remote Project Dialog */}
    <Suspense fallback={null}>
      {createProjectOpen && (
        <CreateProjectDialog
          open={createProjectOpen}
          onOpenChange={setCreateProjectOpen}
          onSuccess={handleCreateProjectSuccess}
        />
      )}
    </Suspense>
  </>
  )
}

