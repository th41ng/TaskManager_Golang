import { jsx as _jsx, jsxs as _jsxs, Fragment as _Fragment } from "react/jsx-runtime";
import { TrendingUpIcon, FolderIcon, CheckSquareIcon, ListTodoIcon, ClockIcon, PlusIcon } from "lucide-react";
import { useEffect, useState, lazy, Suspense } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardDescription, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card";
import { projectsApi, tasksApi } from "@/lib/api";
// Import remote dialogs
const CreateProjectDialog = lazy(() => import("projectApp/ProjectApp").then(m => ({ default: m.CreateProjectDialog })));
export function SectionCards() {
    const [projects, setProjects] = useState([]);
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [createProjectOpen, setCreateProjectOpen] = useState(false);
    useEffect(() => {
        loadData();
    }, []);
    const loadData = () => {
        Promise.all([
            projectsApi.list(),
            tasksApi.list()
        ])
            .then(([projectsData, tasksData]) => {
            console.log('SectionCards - Projects:', projectsData);
            console.log('SectionCards - Tasks:', tasksData);
            setProjects(projectsData);
            setTasks(tasksData);
        })
            .catch(err => console.error('Failed to load data:', err))
            .finally(() => setLoading(false));
    };
    const handleCreateProjectSuccess = () => {
        setCreateProjectOpen(false);
        loadData();
    };
    const completedTasks = tasks.filter(t => t.done).length;
    const totalTasks = tasks.length;
    const completionRate = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;
    const activeTasks = totalTasks - completedTasks;
    if (loading) {
        return (_jsx("div", { className: "grid grid-cols-1 gap-4 px-4 lg:px-6 @xl/main:grid-cols-2 @5xl/main:grid-cols-4", children: [1, 2, 3, 4].map(i => (_jsx(Card, { className: "animate-pulse", children: _jsxs(CardHeader, { children: [_jsx("div", { className: "h-4 bg-gray-200 rounded w-24 mb-2" }), _jsx("div", { className: "h-8 bg-gray-200 rounded w-16" })] }) }, i))) }));
    }
    return (_jsxs(_Fragment, { children: [_jsxs("div", { className: "flex items-center justify-between px-4 lg:px-6 mb-2", children: [_jsxs("div", { children: [_jsx("h2", { className: "text-2xl font-bold", children: "Dashboard Overview" }), _jsx("p", { className: "text-sm text-muted-foreground", children: "Quick stats and metrics" })] }), _jsxs(Button, { onClick: () => setCreateProjectOpen(true), size: "sm", children: [_jsx(PlusIcon, { className: "mr-2 h-4 w-4" }), "New Project"] })] }), _jsxs("div", { className: "*:data-[slot=card]:shadow-xs @xl/main:grid-cols-2 @5xl/main:grid-cols-4 grid grid-cols-1 gap-4 px-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card dark:*:data-[slot=card]:bg-card lg:px-6", children: [_jsxs(Card, { className: "@container/card", children: [_jsxs(CardHeader, { className: "relative", children: [_jsxs(CardDescription, { className: "flex items-center gap-2", children: [_jsx(FolderIcon, { className: "size-4" }), "Total Projects"] }), _jsx(CardTitle, { className: "@[250px]/card:text-3xl text-2xl font-semibold tabular-nums", children: projects.length })] }), _jsx(CardFooter, { className: "flex-col items-start gap-1 text-sm", children: _jsx("div", { className: "line-clamp-1 flex gap-2 font-medium text-muted-foreground", children: "Active projects in system" }) })] }), _jsxs(Card, { className: "@container/card", children: [_jsxs(CardHeader, { className: "relative", children: [_jsxs(CardDescription, { className: "flex items-center gap-2", children: [_jsx(ListTodoIcon, { className: "size-4" }), "Total Tasks"] }), _jsx(CardTitle, { className: "@[250px]/card:text-3xl text-2xl font-semibold tabular-nums", children: totalTasks }), _jsx("div", { className: "absolute right-4 top-4", children: _jsxs(Badge, { variant: "outline", className: "flex gap-1 rounded-lg text-xs", children: [completionRate, "% done"] }) })] }), _jsx(CardFooter, { className: "flex-col items-start gap-1 text-sm", children: _jsxs("div", { className: "line-clamp-1 flex gap-2 font-medium text-muted-foreground", children: [completedTasks, " completed, ", activeTasks, " active"] }) })] }), _jsxs(Card, { className: "@container/card", children: [_jsxs(CardHeader, { className: "relative", children: [_jsxs(CardDescription, { className: "flex items-center gap-2", children: [_jsx(CheckSquareIcon, { className: "size-4" }), "Completed Tasks"] }), _jsx(CardTitle, { className: "@[250px]/card:text-3xl text-2xl font-semibold tabular-nums", children: completedTasks }), _jsx("div", { className: "absolute right-4 top-4", children: _jsxs(Badge, { variant: "outline", className: "flex gap-1 rounded-lg text-xs", children: [_jsx(TrendingUpIcon, { className: "size-3" }), completionRate, "%"] }) })] }), _jsxs(CardFooter, { className: "flex-col items-start gap-1 text-sm", children: [_jsxs("div", { className: "line-clamp-1 flex gap-2 font-medium", children: ["Good progress ", _jsx(TrendingUpIcon, { className: "size-4" })] }), _jsx("div", { className: "text-muted-foreground", children: "Tasks marked as done" })] })] }), _jsxs(Card, { className: "@container/card", children: [_jsxs(CardHeader, { className: "relative", children: [_jsxs(CardDescription, { className: "flex items-center gap-2", children: [_jsx(ClockIcon, { className: "size-4" }), "Active Tasks"] }), _jsx(CardTitle, { className: "@[250px]/card:text-3xl text-2xl font-semibold tabular-nums", children: activeTasks })] }), _jsx(CardFooter, { className: "flex-col items-start gap-1 text-sm", children: _jsx("div", { className: "line-clamp-1 flex gap-2 font-medium text-muted-foreground", children: "Pending completion" }) })] })] }), _jsx(Suspense, { fallback: null, children: createProjectOpen && (_jsx(CreateProjectDialog, { open: createProjectOpen, onOpenChange: setCreateProjectOpen, onSuccess: handleCreateProjectSuccess })) })] }));
}
