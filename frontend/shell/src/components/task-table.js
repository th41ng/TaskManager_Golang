import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { useState, useEffect, lazy, Suspense } from "react";
import { CheckIcon, PlusIcon } from "lucide-react";
import { tasksApi, projectsApi } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, } from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
// Import remote dialogs
const CreateTaskDialog = lazy(() => import("taskApp/TaskApp").then(m => ({ default: m.CreateTaskDialog })));
const EditTaskDialog = lazy(() => import("taskApp/TaskApp").then(m => ({ default: m.EditTaskDialog })));
export function TaskTable({ projectId, onTaskUpdate }) {
    const [tasks, setTasks] = useState([]);
    const [projects, setProjects] = useState([]);
    const [loading, setLoading] = useState(true);
    const [createDialogOpen, setCreateDialogOpen] = useState(false);
    const [editDialogOpen, setEditDialogOpen] = useState(false);
    const [selectedTask, setSelectedTask] = useState(null);
    const loadData = async () => {
        setLoading(true);
        try {
            const [tasksData, projectsData] = await Promise.all([
                tasksApi.list(),
                projectsApi.list()
            ]);
            console.log('Tasks loaded:', tasksData);
            console.log('Projects loaded:', projectsData);
            let filteredTasks = tasksData;
            if (projectId) {
                filteredTasks = tasksData.filter(t => t.project_id === projectId);
            }
            setTasks(filteredTasks);
            setProjects(projectsData);
        }
        catch (err) {
            console.error('Failed to load data:', err);
        }
        finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        loadData();
    }, [projectId]);
    const handleToggleComplete = async (task) => {
        try {
            const updatedTask = await tasksApi.toggleComplete(task);
            setTasks(prev => prev.map(t => t.id === task.id ? updatedTask : t));
            onTaskUpdate?.();
        }
        catch (err) {
            console.error('Failed to toggle task:', err);
        }
    };
    const handleCreateSuccess = () => {
        setCreateDialogOpen(false);
        loadData();
        onTaskUpdate?.();
    };
    const handleEditSuccess = () => {
        setEditDialogOpen(false);
        setSelectedTask(null);
        loadData();
        onTaskUpdate?.();
    };
    const handleEditTask = (task) => {
        setSelectedTask(task);
        setEditDialogOpen(true);
    };
    const getProjectName = (projectId) => {
        const project = projects.find(p => p.id === projectId);
        return project?.name || `Project ${projectId}`;
    };
    const getPriorityBadge = (priority) => {
        const variants = {
            1: { label: 'Low', className: 'bg-blue-50 text-blue-700 border-blue-200' },
            2: { label: 'Medium', className: 'bg-yellow-50 text-yellow-700 border-yellow-200' },
            3: { label: 'High', className: 'bg-red-50 text-red-700 border-red-200' },
        };
        const variant = variants[priority] || variants[2];
        return (_jsx(Badge, { variant: "outline", className: variant.className, children: variant.label }));
    };
    // Group tasks by project
    const groupedTasks = tasks.reduce((acc, task) => {
        const key = task.project_id;
        if (!acc[key])
            acc[key] = [];
        acc[key].push(task);
        return acc;
    }, {});
    if (loading) {
        return (_jsx("div", { className: "px-4 lg:px-6", children: _jsx("div", { className: "rounded-lg border bg-card", children: _jsx("div", { className: "p-6 text-center text-muted-foreground", children: "Loading tasks..." }) }) }));
    }
    return (_jsxs("div", { className: "px-4 lg:px-6", children: [_jsxs("div", { className: "rounded-lg border bg-card", children: [_jsxs("div", { className: "flex items-center justify-between p-4 border-b", children: [_jsxs("div", { children: [_jsx("h3", { className: "text-lg font-semibold", children: "Tasks" }), _jsxs("p", { className: "text-sm text-muted-foreground", children: [tasks.filter(t => !t.done).length, " active, ", tasks.filter(t => t.done).length, " completed"] })] }), _jsxs(Button, { onClick: () => setCreateDialogOpen(true), size: "sm", children: [_jsx(PlusIcon, { className: "mr-2 h-4 w-4" }), "New Task"] })] }), tasks.length === 0 ? (_jsx("div", { className: "p-8 text-center text-muted-foreground", children: "No tasks yet. Create one to get started!" })) : (_jsx("div", { children: Object.entries(groupedTasks).map(([projId, projectTasks]) => (_jsxs("div", { children: [!projectId && (_jsx("div", { className: "px-4 py-3 bg-muted/30 border-b", children: _jsx("h4", { className: "font-medium text-sm", children: getProjectName(Number(projId)) }) })), _jsxs(Table, { children: [_jsx(TableHeader, { children: _jsxs(TableRow, { children: [_jsx(TableHead, { className: "w-12" }), _jsx(TableHead, { children: "Task" }), !projectId && _jsx(TableHead, { children: "Project" }), _jsx(TableHead, { className: "w-32", children: "Priority" }), _jsx(TableHead, { className: "w-32", children: "Status" })] }) }), _jsx(TableBody, { children: projectTasks.map(task => (_jsxs(TableRow, { className: "cursor-pointer hover:bg-muted/50", onClick: () => handleEditTask(task), children: [_jsx(TableCell, { onClick: (e) => e.stopPropagation(), children: _jsx(Checkbox, { checked: task.done, onCheckedChange: () => handleToggleComplete(task) }) }), _jsx(TableCell, { children: _jsx("span", { className: task.done ? 'line-through text-muted-foreground' : '', children: task.title }) }), !projectId && (_jsx(TableCell, { className: "text-sm text-muted-foreground", children: getProjectName(task.project_id) })), _jsx(TableCell, { children: getPriorityBadge(task.priority) }), _jsx(TableCell, { children: task.done ? (_jsxs(Badge, { variant: "outline", className: "bg-green-50 text-green-700 border-green-200", children: [_jsx(CheckIcon, { className: "mr-1 size-3" }), "Done"] })) : (_jsx(Badge, { variant: "outline", children: "Active" })) })] }, task.id))) })] })] }, projId))) }))] }), _jsx(Suspense, { fallback: null, children: createDialogOpen && (_jsx(CreateTaskDialog, { open: createDialogOpen, onOpenChange: setCreateDialogOpen, onSuccess: handleCreateSuccess, projectId: projectId })) }), _jsx(Suspense, { fallback: null, children: editDialogOpen && selectedTask && (_jsx(EditTaskDialog, { open: editDialogOpen, onOpenChange: setEditDialogOpen, onSuccess: handleEditSuccess, task: selectedTask })) })] }));
}
