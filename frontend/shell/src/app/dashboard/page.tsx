import { AppSidebar } from "@/components/app-sidebar"
import { SectionCards } from "@/components/section-cards"
import { TaskTable } from "@/components/task-table"
import { SiteHeader } from "@/components/site-header"
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar"
import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { auth } from "@/lib/auth"

export default function Page() {
  const navigate = useNavigate()

  useEffect(() => {
    // Check if user is authenticated
    if (!auth.isAuthenticated()) {
      // window.location.href = 'http://172.21.223.107:4001/login'
       // Navigate to local login route so dev uses the shell's /login route (remote UserApp)
      navigate('/login')
      
    }
  }, [navigate])

  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "calc(var(--spacing) * 72)",
          "--header-height": "calc(var(--spacing) * 12)",
        } as React.CSSProperties
      }
    >
      <AppSidebar variant="inset" />
      <SidebarInset>
        <SiteHeader />
        <div className="flex flex-1 flex-col">
          <div className="@container/main flex flex-1 flex-col gap-2">
            <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
              <SectionCards />
              <TaskTable />
            </div>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
