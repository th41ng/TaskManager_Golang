import * as React from "react"
import { useEffect, useState, lazy, Suspense } from "react"
import {
  ArrowUpCircleIcon,
  BarChartIcon,
  FolderIcon,
  HelpCircleIcon,
  LayoutDashboardIcon,
  PlusIcon,
  SearchIcon,
  SettingsIcon,
  UsersIcon,
  CheckSquareIcon,
  FolderKanbanIcon,
} from "lucide-react"
import { useLocation } from "react-router"

import { NavMain } from "@/components/nav-main"
import { NavSecondary } from "@/components/nav-secondary"
import { NavUser } from "@/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarGroupContent,
} from "@/components/ui/sidebar"
import { Button } from "@/components/ui/button"
import { projectsApi } from "@/lib/api"
import type { Project } from "@/types"

// Import remote dialog
const CreateProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.CreateProjectDialog })))

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const [projects, setProjects] = useState<Project[]>([])
  const [createProjectOpen, setCreateProjectOpen] = useState(false)
  const location = useLocation()

  useEffect(() => {
    loadProjects()
  }, [])

  const loadProjects = () => {
    projectsApi.list()
      .then(setProjects)
      .catch(err => console.error('Failed to load projects:', err))
  }

  const handleCreateProjectSuccess = () => {
    setCreateProjectOpen(false)
    loadProjects()
  }

  const navMain = [
    {
      title: "Dashboard",
      url: "/",
      icon: LayoutDashboardIcon,
    },
    {
      title: "Tasks",
      url: "/tasks",
      icon: CheckSquareIcon,
    },
    {
      title: "Projects",
      url: "/projects-app",
      icon: FolderKanbanIcon,
    },
    {
      title: "Users",
      url: "/users",
      icon: UsersIcon,
    },
  ]

  const navSecondary = [
    {
      title: "Settings",
      url: "#",
      icon: SettingsIcon,
    },
    {
      title: "Get Help",
      url: "#",
      icon: HelpCircleIcon,
    },
    {
      title: "Search",
      url: "#",
      icon: SearchIcon,
    },
  ]

  const user = {
    name: "Admin User",
    email: "admin@taskmanager.com",
    avatar: "/avatars/default.jpg",
  }

  return (
    <Sidebar collapsible="offcanvas" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              asChild
              className="data-[slot=sidebar-menu-button]:!p-1.5"
            >
              <a href="/">
                <LayoutDashboardIcon className="h-5 w-5" />
                <span className="text-base font-semibold">TaskManager</span>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={navMain} />
        
        {/* ================================================================
            PROJECT LIST SECTION
            ================================================================
            Hiển thị danh sách projects hiện có
            Khi click vào project → navigate đến /projects/:id (detail page)
            Button "+" → Mở CreateProjectDialog
            
            Note: Khác với menu "Projects" ở trên:
            - Menu "Projects" → /projects-app (full page ProjectApp)
            - Section này → /projects/:id (project detail)
        */}
        <SidebarGroup>
          <SidebarGroupLabel className="flex items-center justify-between">
            <span>My Projects</span>
            <Button 
              variant="ghost" 
              size="icon" 
              className="h-5 w-5"
              onClick={() => setCreateProjectOpen(true)}
            >
              <PlusIcon className="h-4 w-4" />
            </Button>
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {projects.length === 0 ? (
                <div className="px-2 py-4 text-xs text-muted-foreground">
                  No projects yet
                </div>
              ) : (
                projects.map((project) => (
                  <SidebarMenuItem key={project.id}>
                    <SidebarMenuButton
                      asChild
                      isActive={location.pathname === `/projects/${project.id}`}
                    >
                      <a href={`/projects/${project.id}`}>
                        <FolderIcon className="size-4" />
                        <span>{project.name}</span>
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))
              )}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
        
        <NavSecondary items={navSecondary} className="mt-auto" />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user} />
      </SidebarFooter>

      {/* Remote Create Project Dialog */}
      <Suspense fallback={null}>
        {createProjectOpen && (
          <CreateProjectDialog
            open={createProjectOpen}
            onOpenChange={setCreateProjectOpen}
            onSuccess={handleCreateProjectSuccess}
          />
        )}
      </Suspense>
    </Sidebar>
  )
}
