import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { useEffect, useState, lazy, Suspense } from "react";
import { FolderIcon, HelpCircleIcon, LayoutDashboardIcon, PlusIcon, SearchIcon, SettingsIcon, UsersIcon, CheckSquareIcon, FolderKanbanIcon, } from "lucide-react";
import { useLocation } from "react-router";
import { NavMain } from "@/components/nav-main";
import { NavSecondary } from "@/components/nav-secondary";
import { NavUser } from "@/components/nav-user";
import { Sidebar, SidebarContent, SidebarFooter, SidebarHeader, SidebarMenu, SidebarMenuButton, SidebarMenuItem, SidebarGroup, SidebarGroupLabel, SidebarGroupContent, } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import { projectsApi } from "@/lib/api";
// Import remote dialog
const CreateProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.CreateProjectDialog })));
export function AppSidebar({ ...props }) {
    const [projects, setProjects] = useState([]);
    const [createProjectOpen, setCreateProjectOpen] = useState(false);
    const location = useLocation();
    useEffect(() => {
        loadProjects();
    }, []);
    const loadProjects = () => {
        projectsApi.list()
            .then(setProjects)
            .catch(err => console.error('Failed to load projects:', err));
    };
    const handleCreateProjectSuccess = () => {
        setCreateProjectOpen(false);
        loadProjects();
    };
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
    ];
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
    ];
    const user = {
        name: "Admin User",
        email: "admin@taskmanager.com",
        avatar: "/avatars/default.jpg",
    };
    return (_jsxs(Sidebar, { collapsible: "offcanvas", ...props, children: [_jsx(SidebarHeader, { children: _jsx(SidebarMenu, { children: _jsx(SidebarMenuItem, { children: _jsx(SidebarMenuButton, { asChild: true, className: "data-[slot=sidebar-menu-button]:!p-1.5", children: _jsxs("a", { href: "/", children: [_jsx(LayoutDashboardIcon, { className: "h-5 w-5" }), _jsx("span", { className: "text-base font-semibold", children: "TaskManager" })] }) }) }) }) }), _jsxs(SidebarContent, { children: [_jsx(NavMain, { items: navMain }), _jsxs(SidebarGroup, { children: [_jsxs(SidebarGroupLabel, { className: "flex items-center justify-between", children: [_jsx("span", { children: "My Projects" }), _jsx(Button, { variant: "ghost", size: "icon", className: "h-5 w-5", onClick: () => setCreateProjectOpen(true), children: _jsx(PlusIcon, { className: "h-4 w-4" }) })] }), _jsx(SidebarGroupContent, { children: _jsx(SidebarMenu, { children: projects.length === 0 ? (_jsx("div", { className: "px-2 py-4 text-xs text-muted-foreground", children: "No projects yet" })) : (projects.map((project) => (_jsx(SidebarMenuItem, { children: _jsx(SidebarMenuButton, { asChild: true, isActive: location.pathname === `/projects/${project.id}`, children: _jsxs("a", { href: `/projects/${project.id}`, children: [_jsx(FolderIcon, { className: "size-4" }), _jsx("span", { children: project.name })] }) }) }, project.id)))) }) })] }), _jsx(NavSecondary, { items: navSecondary, className: "mt-auto" })] }), _jsx(SidebarFooter, { children: _jsx(NavUser, { user: user }) }), _jsx(Suspense, { fallback: null, children: createProjectOpen && (_jsx(CreateProjectDialog, { open: createProjectOpen, onOpenChange: setCreateProjectOpen, onSuccess: handleCreateProjectSuccess })) })] }));
}
