import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { useEffect, useState, lazy, Suspense } from "react";
import { useParams } from "react-router";
import { FolderIcon, CalendarIcon, UserIcon, TrendingUpIcon, MoreVerticalIcon, PencilIcon, TrashIcon, } from "lucide-react";
import { projectsApi, tasksApi } from "@/lib/api";
import { TaskTable } from "@/components/task-table";
import { Card, CardContent, CardDescription, CardHeader, CardTitle, } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger, } from "@/components/ui/dropdown-menu";
// Import remote dialog
const EditProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.EditProjectDialog })));
export default function ProjectDetailPage() {
    const { id } = useParams();
    const [project, setProject] = useState(null);
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [editDialogOpen, setEditDialogOpen] = useState(false);
    const loadProjectData = async () => {
        if (!id)
            return;
        setLoading(true);
        try {
            const [projectData, tasksData] = await Promise.all([
                projectsApi.get(Number(id)),
                tasksApi.list()
            ]);
            setProject(projectData);
            setTasks(tasksData.filter(t => t.project_id === Number(id)));
        }
        catch (err) {
            console.error('Failed to load project:', err);
        }
        finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        loadProjectData();
    }, [id]);
    const handleDelete = async () => {
        if (!project || !confirm(`Delete project "${project.name}"?`))
            return;
        try {
            await projectsApi.delete(project.id);
            window.location.href = '/';
        }
        catch (err) {
            console.error('Failed to delete project:', err);
            alert('Failed to delete project');
        }
    };
    const handleEditSuccess = () => {
        setEditDialogOpen(false);
        loadProjectData();
    };
    if (loading) {
        return (_jsx("div", { className: "p-6", children: _jsxs("div", { className: "animate-pulse space-y-4", children: [_jsx("div", { className: "h-8 bg-gray-200 rounded w-1/3" }), _jsx("div", { className: "h-48 bg-gray-200 rounded" })] }) }));
    }
    if (!project) {
        return (_jsx("div", { className: "p-6", children: _jsxs("div", { className: "text-center", children: [_jsx("h2", { className: "text-xl font-semibold mb-2", children: "Project not found" }), _jsx("p", { className: "text-muted-foreground", children: "The project you're looking for doesn't exist." }), _jsx(Button, { asChild: true, className: "mt-4", children: _jsx("a", { href: "/", children: "Back to Dashboard" }) })] }) }));
    }
    const completedTasks = tasks.filter(t => t.done).length;
    const totalTasks = tasks.length;
    const progressPercent = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;
    return (_jsxs("div", { className: "flex flex-col gap-6 p-6", children: [_jsxs("div", { className: "flex items-center justify-between", children: [_jsxs("div", { children: [_jsxs("h1", { className: "text-3xl font-bold flex items-center gap-3", children: [_jsx(FolderIcon, { className: "size-8" }), project.name] }), _jsxs("p", { className: "text-muted-foreground mt-1", children: ["Project #", project.id] })] }), _jsxs(DropdownMenu, { children: [_jsx(DropdownMenuTrigger, { asChild: true, children: _jsx(Button, { variant: "outline", size: "icon", children: _jsx(MoreVerticalIcon, { className: "size-4" }) }) }), _jsxs(DropdownMenuContent, { align: "end", children: [_jsxs(DropdownMenuItem, { onClick: () => setEditDialogOpen(true), children: [_jsx(PencilIcon, { className: "mr-2 size-4" }), "Edit Project"] }), _jsx(DropdownMenuSeparator, {}), _jsxs(DropdownMenuItem, { onClick: handleDelete, className: "text-red-600", children: [_jsx(TrashIcon, { className: "mr-2 size-4" }), "Delete Project"] })] })] })] }), _jsxs(Card, { children: [_jsxs(CardHeader, { children: [_jsxs(CardTitle, { className: "flex items-center gap-2", children: [_jsx(TrendingUpIcon, { className: "size-5" }), "Project Overview"] }), _jsx(CardDescription, { children: "Progress and statistics" })] }), _jsxs(CardContent, { className: "space-y-6", children: [_jsxs("div", { className: "grid grid-cols-1 md:grid-cols-3 gap-4", children: [_jsxs("div", { className: "flex items-center gap-3", children: [_jsx("div", { className: "p-3 rounded-lg bg-blue-50", children: _jsx(UserIcon, { className: "size-5 text-blue-600" }) }), _jsxs("div", { children: [_jsx("p", { className: "text-sm text-muted-foreground", children: "Owner ID" }), _jsx("p", { className: "font-semibold", children: project.owner_id })] })] }), project.created_at && (_jsxs("div", { className: "flex items-center gap-3", children: [_jsx("div", { className: "p-3 rounded-lg bg-green-50", children: _jsx(CalendarIcon, { className: "size-5 text-green-600" }) }), _jsxs("div", { children: [_jsx("p", { className: "text-sm text-muted-foreground", children: "Created" }), _jsx("p", { className: "font-semibold", children: new Date(project.created_at).toLocaleDateString() })] })] })), _jsxs("div", { className: "flex items-center gap-3", children: [_jsx("div", { className: "p-3 rounded-lg bg-purple-50", children: _jsx(FolderIcon, { className: "size-5 text-purple-600" }) }), _jsxs("div", { children: [_jsx("p", { className: "text-sm text-muted-foreground", children: "Total Tasks" }), _jsx("p", { className: "font-semibold", children: totalTasks })] })] })] }), _jsxs("div", { className: "space-y-2", children: [_jsxs("div", { className: "flex items-center justify-between text-sm", children: [_jsx("span", { className: "text-muted-foreground", children: "Completion Progress" }), _jsxs("span", { className: "font-semibold", children: [progressPercent, "%"] })] }), _jsx("div", { className: "relative w-full h-2 bg-secondary rounded-full overflow-hidden", children: _jsx("div", { className: "absolute h-full bg-primary transition-all duration-300", style: { width: `${progressPercent}%` } }) }), _jsxs("div", { className: "flex items-center justify-between text-xs text-muted-foreground", children: [_jsxs("span", { children: [completedTasks, " completed"] }), _jsxs("span", { children: [totalTasks - completedTasks, " remaining"] })] })] })] })] }), _jsxs("div", { className: "space-y-4", children: [_jsx("div", { className: "flex items-center justify-between", children: _jsx("h2", { className: "text-xl font-semibold", children: "Project Tasks" }) }), _jsx(TaskTable, { projectId: Number(id), onTaskUpdate: loadProjectData })] }), _jsx(Suspense, { fallback: null, children: editDialogOpen && project && (_jsx(EditProjectDialog, { open: editDialogOpen, onOpenChange: setEditDialogOpen, onSuccess: handleEditSuccess, project: project })) })] }));
}
