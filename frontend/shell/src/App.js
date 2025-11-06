import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { BrowserRouter, Routes, Route, useNavigate } from "react-router-dom";
import { lazy, Suspense, useState, useEffect } from "react";
import { ThemeProvider } from "@/components/theme-provider";
import DashboardPage from "@/app/dashboard/page";
import ProjectDetailPage from "@/app/projects/[id]/page";
import Layout from "@/components/layout";
// ============================================================================
// MICRO FRONTEND IMPORTS - Remote Applications
// ============================================================================
// Kiến trúc: Shell owns Layout, Remote owns Content Only
// - Shell: Quản lý routing, layout (sidebar + header), truyền data xuống remote
// - Remote: Chỉ render nội dung, nhận props từ shell, gọi callbacks về shell
//
// Type declarations: src/types/remote-modules.d.ts
const UserApp = lazy(() => import("userApp/UserApp"));
const ProjectApp = lazy(() => import("projectApp/ProjectApp"));
const TaskApp = lazy(() => import("taskApp/TaskApp"));
export default function App() {
    // ============================================================================
    // SHELL STATE - Shared data để truyền xuống remote apps
    // ============================================================================
    const [shellData, setShellData] = useState({
        user: null,
        theme: 'light',
        token: localStorage.getItem('token') || '',
    });
    // Load user info khi app khởi động
    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            // TODO: Fetch user info from API
            setShellData(prev => ({
                ...prev,
                token,
                user: { name: 'Admin User', email: 'admin@taskmanager.com' }
            }));
        }
    }, []);
    // target project id when Shell receives openTaskList event
    const [targetProjectId, setTargetProjectId] = useState(null);
    // ============================================================================
    // CALLBACK HANDLERS - Remote apps gọi về Shell
    // ============================================================================
    const handleTaskEvent = (event, data) => {
        console.log('📨 Task Event:', event, data);
        // TODO: Handle task events (refresh sidebar, show notification, etc.)
    };
    const handleProjectEvent = (event, data) => {
        console.log('📨 Project Event:', event, data);
        // TODO: Handle project events (refresh sidebar, show notification, etc.)
    };
    // Inner component so we can use useNavigate() inside BrowserRouter
    function AppRoutes() {
        const navigate = useNavigate();
        useEffect(() => {
            const handler = (e) => {
                const projectId = e?.detail;
                setTargetProjectId(projectId || null);
                // Navigate to tasks and include projectId as query param
                navigate(`/tasks${projectId ? `?projectId=${projectId}` : ''}`);
            };
            window.addEventListener('openTaskList', handler);
            return () => window.removeEventListener('openTaskList', handler);
        }, [navigate]);
        return (_jsxs(Routes, { children: [_jsx(Route, { path: "/", element: _jsx(DashboardPage, {}) }), _jsx(Route, { path: "/login/*", element: _jsx(Suspense, { fallback: _jsx("div", { children: "Loading..." }), children: _jsx(UserApp, {}) }) }), _jsx(Route, { path: "/users/*", element: _jsx(Layout, { breadcrumbItems: [
                            { label: "Dashboard", href: "/" },
                            { label: "Users", isActive: true }
                        ], children: _jsx(Suspense, { fallback: _jsx("div", { children: "Loading..." }), children: _jsx(UserApp, {}) }) }) }), _jsx(Route, { path: "/projects/:id", element: _jsx(Layout, { breadcrumbItems: [
                            { label: "Dashboard", href: "/" },
                            { label: "Projects", href: "/" },
                            { label: "Detail", isActive: true }
                        ], children: _jsx(Suspense, { fallback: _jsx("div", { children: "Loading Project..." }), children: _jsx(ProjectDetailPage, {}) }) }) }), _jsx(Route, { path: "/projects-app/*", element: _jsx(Layout, { breadcrumbItems: [
                            { label: "Dashboard", href: "/" },
                            { label: "Projects Management", isActive: true }
                        ], children: _jsx(Suspense, { fallback: _jsx("div", { className: "p-6", children: "Loading Projects..." }), children: _jsx(ProjectApp, { shellData: shellData, onProjectCreated: (project) => handleProjectEvent('created', project), onProjectUpdated: (project) => handleProjectEvent('updated', project) }) }) }) }), _jsx(Route, { path: "/tasks/*", element: _jsx(Layout, { breadcrumbItems: [
                            { label: "Dashboard", href: "/" },
                            { label: "Tasks Management", isActive: true }
                        ], children: _jsx(Suspense, { fallback: _jsx("div", { className: "p-6", children: "Loading Tasks..." }), children: _jsx(TaskApp, { shellData: shellData, projectId: targetProjectId ?? undefined, onTaskCreated: (task) => handleTaskEvent('created', task), onTaskUpdated: (task) => handleTaskEvent('updated', task) }) }) }) })] }));
    }
    return (_jsx("div", { id: "shell-root", children: _jsx(ThemeProvider, { defaultTheme: "system", storageKey: "vite-ui-theme", children: _jsx(BrowserRouter, { children: _jsx(AppRoutes, {}) }) }) }));
}
