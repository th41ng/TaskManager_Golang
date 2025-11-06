import { jsx as _jsx, jsxs as _jsxs, Fragment as _Fragment } from "react/jsx-runtime";
import * as React from "react";
import { DndContext, KeyboardSensor, MouseSensor, TouchSensor, closestCenter, useSensor, useSensors, } from "@dnd-kit/core";
import { restrictToVerticalAxis } from "@dnd-kit/modifiers";
import { SortableContext, arrayMove, useSortable, verticalListSortingStrategy, } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { flexRender, getCoreRowModel, getFacetedRowModel, getFacetedUniqueValues, getFilteredRowModel, getPaginationRowModel, getSortedRowModel, useReactTable, } from "@tanstack/react-table";
import { CheckCircle2Icon, ChevronDownIcon, ChevronLeftIcon, ChevronRightIcon, ChevronsLeftIcon, ChevronsRightIcon, ColumnsIcon, GripVerticalIcon, LoaderIcon, MoreVerticalIcon, PlusIcon, TrendingUpIcon, } from "lucide-react";
import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";
import { toast } from "sonner";
import { z } from "zod";
import { useIsMobile } from "@/hooks/use-mobile";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ChartContainer, ChartTooltip, ChartTooltipContent, } from "@/components/ui/chart";
import { Checkbox } from "@/components/ui/checkbox";
import { DropdownMenu, DropdownMenuCheckboxItem, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger, } from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue, } from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { Sheet, SheetClose, SheetContent, SheetDescription, SheetFooter, SheetHeader, SheetTitle, SheetTrigger, } from "@/components/ui/sheet";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, } from "@/components/ui/table";
import { Tabs, TabsContent, TabsList, TabsTrigger, } from "@/components/ui/tabs";
export const schema = z.object({
    id: z.number(),
    header: z.string(),
    type: z.string(),
    status: z.string(),
    target: z.string(),
    limit: z.string(),
    reviewer: z.string(),
});
// Create a separate component for the drag handle
function DragHandle({ id }) {
    const { attributes, listeners } = useSortable({
        id,
    });
    return (_jsxs(Button, { ...attributes, ...listeners, variant: "ghost", size: "icon", className: "size-7 text-muted-foreground hover:bg-transparent", children: [_jsx(GripVerticalIcon, { className: "size-3 text-muted-foreground" }), _jsx("span", { className: "sr-only", children: "Drag to reorder" })] }));
}
const columns = [
    {
        id: "drag",
        header: () => null,
        cell: ({ row }) => _jsx(DragHandle, { id: row.original.id }),
    },
    {
        id: "select",
        header: ({ table }) => (_jsx("div", { className: "flex items-center justify-center", children: _jsx(Checkbox, { checked: table.getIsAllPageRowsSelected()
                    ? true
                    : table.getIsSomePageRowsSelected()
                        ? "indeterminate"
                        : false, onCheckedChange: (value) => table.toggleAllPageRowsSelected(!!value), "aria-label": "Select all" }) })),
        cell: ({ row }) => (_jsx("div", { className: "flex items-center justify-center", children: _jsx(Checkbox, { checked: row.getIsSelected(), onCheckedChange: (value) => row.toggleSelected(!!value), "aria-label": "Select row" }) })),
        enableSorting: false,
        enableHiding: false,
    },
    {
        accessorKey: "header",
        header: "Header",
        cell: ({ row }) => {
            return _jsx(TableCellViewer, { item: row.original });
        },
        enableHiding: false,
    },
    {
        accessorKey: "type",
        header: "Section Type",
        cell: ({ row }) => (_jsx("div", { className: "w-32", children: _jsx(Badge, { variant: "outline", className: "px-1.5 text-muted-foreground", children: row.original.type }) })),
    },
    {
        accessorKey: "status",
        header: "Status",
        cell: ({ row }) => (_jsxs(Badge, { variant: "outline", className: "flex gap-1 px-1.5 text-muted-foreground [&_svg]:size-3", children: [row.original.status === "Done" ? (_jsx(CheckCircle2Icon, { className: "text-green-500 dark:text-green-400" })) : (_jsx(LoaderIcon, {})), row.original.status] })),
    },
    {
        accessorKey: "target",
        header: () => _jsx("div", { className: "w-full text-right", children: "Target" }),
        cell: ({ row }) => (_jsxs("form", { onSubmit: (e) => {
                e.preventDefault();
                toast.promise(new Promise((resolve) => setTimeout(resolve, 1000)), {
                    loading: `Saving ${row.original.header}`,
                    success: "Done",
                    error: "Error",
                });
            }, children: [_jsx(Label, { htmlFor: `${row.original.id}-target`, className: "sr-only", children: "Target" }), _jsx(Input, { className: "h-8 w-16 border-transparent bg-transparent text-right shadow-none hover:bg-input/30 focus-visible:border focus-visible:bg-background", defaultValue: row.original.target, id: `${row.original.id}-target` })] })),
    },
    {
        accessorKey: "limit",
        header: () => _jsx("div", { className: "w-full text-right", children: "Limit" }),
        cell: ({ row }) => (_jsxs("form", { onSubmit: (e) => {
                e.preventDefault();
                toast.promise(new Promise((resolve) => setTimeout(resolve, 1000)), {
                    loading: `Saving ${row.original.header}`,
                    success: "Done",
                    error: "Error",
                });
            }, children: [_jsx(Label, { htmlFor: `${row.original.id}-limit`, className: "sr-only", children: "Limit" }), _jsx(Input, { className: "h-8 w-16 border-transparent bg-transparent text-right shadow-none hover:bg-input/30 focus-visible:border focus-visible:bg-background", defaultValue: row.original.limit, id: `${row.original.id}-limit` })] })),
    },
    {
        accessorKey: "reviewer",
        header: "Reviewer",
        cell: ({ row }) => {
            const isAssigned = row.original.reviewer !== "Assign reviewer";
            if (isAssigned) {
                return row.original.reviewer;
            }
            return (_jsxs(_Fragment, { children: [_jsx(Label, { htmlFor: `${row.original.id}-reviewer`, className: "sr-only", children: "Reviewer" }), _jsxs(Select, { children: [_jsx(SelectTrigger, { className: "h-8 w-40", id: `${row.original.id}-reviewer`, children: _jsx(SelectValue, { placeholder: "Assign reviewer" }) }), _jsxs(SelectContent, { align: "end", children: [_jsx(SelectItem, { value: "Eddie Lake", children: "Eddie Lake" }), _jsx(SelectItem, { value: "Jamik Tashpulatov", children: "Jamik Tashpulatov" })] })] })] }));
        },
    },
    {
        id: "actions",
        cell: () => (_jsxs(DropdownMenu, { children: [_jsx(DropdownMenuTrigger, { asChild: true, children: _jsxs(Button, { variant: "ghost", className: "flex size-8 text-muted-foreground data-[state=open]:bg-muted", size: "icon", children: [_jsx(MoreVerticalIcon, {}), _jsx("span", { className: "sr-only", children: "Open menu" })] }) }), _jsxs(DropdownMenuContent, { align: "end", className: "w-32", children: [_jsx(DropdownMenuItem, { children: "Edit" }), _jsx(DropdownMenuItem, { children: "Make a copy" }), _jsx(DropdownMenuItem, { children: "Favorite" }), _jsx(DropdownMenuSeparator, {}), _jsx(DropdownMenuItem, { children: "Delete" })] })] })),
    },
];
function DraggableRow({ row }) {
    const { transform, transition, setNodeRef, isDragging } = useSortable({
        id: row.original.id,
    });
    return (_jsx(TableRow, { "data-state": row.getIsSelected() && "selected", "data-dragging": isDragging, ref: setNodeRef, className: "relative z-0 data-[dragging=true]:z-10 data-[dragging=true]:opacity-80", style: {
            transform: CSS.Transform.toString(transform),
            transition: transition,
        }, children: row.getVisibleCells().map((cell) => (_jsx(TableCell, { children: flexRender(cell.column.columnDef.cell, cell.getContext()) }, cell.id))) }));
}
export function DataTable({ data: initialData, }) {
    const [data, setData] = React.useState(() => initialData);
    const [rowSelection, setRowSelection] = React.useState({});
    const [columnVisibility, setColumnVisibility] = React.useState({});
    const [columnFilters, setColumnFilters] = React.useState([]);
    const [sorting, setSorting] = React.useState([]);
    const [pagination, setPagination] = React.useState({
        pageIndex: 0,
        pageSize: 10,
    });
    const sortableId = React.useId();
    const sensors = useSensors(useSensor(MouseSensor, {}), useSensor(TouchSensor, {}), useSensor(KeyboardSensor, {}));
    const dataIds = React.useMemo(() => data?.map(({ id }) => id) || [], [data]);
    const table = useReactTable({
        data,
        columns,
        state: {
            sorting,
            columnVisibility,
            rowSelection,
            columnFilters,
            pagination,
        },
        getRowId: (row) => row.id.toString(),
        enableRowSelection: true,
        onRowSelectionChange: setRowSelection,
        onSortingChange: setSorting,
        onColumnFiltersChange: setColumnFilters,
        onColumnVisibilityChange: setColumnVisibility,
        onPaginationChange: setPagination,
        getCoreRowModel: getCoreRowModel(),
        getFilteredRowModel: getFilteredRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        getSortedRowModel: getSortedRowModel(),
        getFacetedRowModel: getFacetedRowModel(),
        getFacetedUniqueValues: getFacetedUniqueValues(),
    });
    function handleDragEnd(event) {
        const { active, over } = event;
        if (active && over && active.id !== over.id) {
            setData((data) => {
                const oldIndex = dataIds.indexOf(active.id);
                const newIndex = dataIds.indexOf(over.id);
                return arrayMove(data, oldIndex, newIndex);
            });
        }
    }
    return (_jsxs(Tabs, { defaultValue: "outline", className: "flex w-full flex-col justify-start gap-6", children: [_jsxs("div", { className: "flex items-center justify-between px-4 lg:px-6", children: [_jsx(Label, { htmlFor: "view-selector", className: "sr-only", children: "View" }), _jsxs(Select, { defaultValue: "outline", children: [_jsx(SelectTrigger, { className: "@4xl/main:hidden flex w-fit", id: "view-selector", children: _jsx(SelectValue, { placeholder: "Select a view" }) }), _jsxs(SelectContent, { children: [_jsx(SelectItem, { value: "outline", children: "Outline" }), _jsx(SelectItem, { value: "past-performance", children: "Past Performance" }), _jsx(SelectItem, { value: "key-personnel", children: "Key Personnel" }), _jsx(SelectItem, { value: "focus-documents", children: "Focus Documents" })] })] }), _jsxs(TabsList, { className: "@4xl/main:flex hidden", children: [_jsx(TabsTrigger, { value: "outline", children: "Outline" }), _jsxs(TabsTrigger, { value: "past-performance", className: "gap-1", children: ["Past Performance", " ", _jsx(Badge, { variant: "secondary", className: "flex h-5 w-5 items-center justify-center rounded-full bg-muted-foreground/30", children: "3" })] }), _jsxs(TabsTrigger, { value: "key-personnel", className: "gap-1", children: ["Key Personnel", " ", _jsx(Badge, { variant: "secondary", className: "flex h-5 w-5 items-center justify-center rounded-full bg-muted-foreground/30", children: "2" })] }), _jsx(TabsTrigger, { value: "focus-documents", children: "Focus Documents" })] }), _jsxs("div", { className: "flex items-center gap-2", children: [_jsxs(DropdownMenu, { children: [_jsx(DropdownMenuTrigger, { asChild: true, children: _jsxs(Button, { variant: "outline", size: "sm", children: [_jsx(ColumnsIcon, {}), _jsx("span", { className: "hidden lg:inline", children: "Customize Columns" }), _jsx("span", { className: "lg:hidden", children: "Columns" }), _jsx(ChevronDownIcon, {})] }) }), _jsx(DropdownMenuContent, { align: "end", className: "w-56", children: table
                                            .getAllColumns()
                                            .filter((column) => typeof column.accessorFn !== "undefined" &&
                                            column.getCanHide())
                                            .map((column) => {
                                            return (_jsx(DropdownMenuCheckboxItem, { className: "capitalize", checked: column.getIsVisible(), onCheckedChange: (value) => column.toggleVisibility(!!value), children: column.id }, column.id));
                                        }) })] }), _jsxs(Button, { variant: "outline", size: "sm", children: [_jsx(PlusIcon, {}), _jsx("span", { className: "hidden lg:inline", children: "Add Section" })] })] })] }), _jsxs(TabsContent, { value: "outline", className: "relative flex flex-col gap-4 overflow-auto px-4 lg:px-6", children: [_jsx("div", { className: "overflow-hidden rounded-lg border", children: _jsx(DndContext, { collisionDetection: closestCenter, modifiers: [restrictToVerticalAxis], onDragEnd: handleDragEnd, sensors: sensors, id: sortableId, children: _jsxs(Table, { children: [_jsx(TableHeader, { className: "sticky top-0 z-10 bg-muted", children: table.getHeaderGroups().map((headerGroup) => (_jsx(TableRow, { children: headerGroup.headers.map((header) => {
                                                return (_jsx(TableHead, { colSpan: header.colSpan, children: header.isPlaceholder
                                                        ? null
                                                        : flexRender(header.column.columnDef.header, header.getContext()) }, header.id));
                                            }) }, headerGroup.id))) }), _jsx(TableBody, { className: "**:data-[slot=table-cell]:first:w-8", children: table.getRowModel().rows?.length ? (_jsx(SortableContext, { items: dataIds, strategy: verticalListSortingStrategy, children: table.getRowModel().rows.map((row) => (_jsx(DraggableRow, { row: row }, row.id))) })) : (_jsx(TableRow, { children: _jsx(TableCell, { colSpan: columns.length, className: "h-24 text-center", children: "No results." }) })) })] }) }) }), _jsxs("div", { className: "flex items-center justify-between px-4", children: [_jsxs("div", { className: "hidden flex-1 text-sm text-muted-foreground lg:flex", children: [table.getFilteredSelectedRowModel().rows.length, " of", " ", table.getFilteredRowModel().rows.length, " row(s) selected."] }), _jsxs("div", { className: "flex w-full items-center gap-8 lg:w-fit", children: [_jsxs("div", { className: "hidden items-center gap-2 lg:flex", children: [_jsx(Label, { htmlFor: "rows-per-page", className: "text-sm font-medium", children: "Rows per page" }), _jsxs(Select, { value: `${table.getState().pagination.pageSize}`, onValueChange: (value) => {
                                                    table.setPageSize(Number(value));
                                                }, children: [_jsx(SelectTrigger, { className: "w-20", id: "rows-per-page", children: _jsx(SelectValue, { placeholder: table.getState().pagination.pageSize }) }), _jsx(SelectContent, { side: "top", children: [10, 20, 30, 40, 50].map((pageSize) => (_jsx(SelectItem, { value: `${pageSize}`, children: pageSize }, pageSize))) })] })] }), _jsxs("div", { className: "flex w-fit items-center justify-center text-sm font-medium", children: ["Page ", table.getState().pagination.pageIndex + 1, " of", " ", table.getPageCount()] }), _jsxs("div", { className: "ml-auto flex items-center gap-2 lg:ml-0", children: [_jsxs(Button, { variant: "outline", className: "hidden h-8 w-8 p-0 lg:flex", onClick: () => table.setPageIndex(0), disabled: !table.getCanPreviousPage(), children: [_jsx("span", { className: "sr-only", children: "Go to first page" }), _jsx(ChevronsLeftIcon, {})] }), _jsxs(Button, { variant: "outline", className: "size-8", size: "icon", onClick: () => table.previousPage(), disabled: !table.getCanPreviousPage(), children: [_jsx("span", { className: "sr-only", children: "Go to previous page" }), _jsx(ChevronLeftIcon, {})] }), _jsxs(Button, { variant: "outline", className: "size-8", size: "icon", onClick: () => table.nextPage(), disabled: !table.getCanNextPage(), children: [_jsx("span", { className: "sr-only", children: "Go to next page" }), _jsx(ChevronRightIcon, {})] }), _jsxs(Button, { variant: "outline", className: "hidden size-8 lg:flex", size: "icon", onClick: () => table.setPageIndex(table.getPageCount() - 1), disabled: !table.getCanNextPage(), children: [_jsx("span", { className: "sr-only", children: "Go to last page" }), _jsx(ChevronsRightIcon, {})] })] })] })] })] }), _jsx(TabsContent, { value: "past-performance", className: "flex flex-col px-4 lg:px-6", children: _jsx("div", { className: "aspect-video w-full flex-1 rounded-lg border border-dashed" }) }), _jsx(TabsContent, { value: "key-personnel", className: "flex flex-col px-4 lg:px-6", children: _jsx("div", { className: "aspect-video w-full flex-1 rounded-lg border border-dashed" }) }), _jsx(TabsContent, { value: "focus-documents", className: "flex flex-col px-4 lg:px-6", children: _jsx("div", { className: "aspect-video w-full flex-1 rounded-lg border border-dashed" }) })] }));
}
const chartData = [
    { month: "January", desktop: 186, mobile: 80 },
    { month: "February", desktop: 305, mobile: 200 },
    { month: "March", desktop: 237, mobile: 120 },
    { month: "April", desktop: 73, mobile: 190 },
    { month: "May", desktop: 209, mobile: 130 },
    { month: "June", desktop: 214, mobile: 140 },
];
const chartConfig = {
    desktop: {
        label: "Desktop",
        color: "var(--primary)",
    },
    mobile: {
        label: "Mobile",
        color: "var(--primary)",
    },
};
function TableCellViewer({ item }) {
    const isMobile = useIsMobile();
    return (_jsxs(Sheet, { children: [_jsx(SheetTrigger, { asChild: true, children: _jsx(Button, { variant: "link", className: "w-fit px-0 text-left text-foreground", children: item.header }) }), _jsxs(SheetContent, { side: "right", className: "flex flex-col", children: [_jsxs(SheetHeader, { className: "gap-1", children: [_jsx(SheetTitle, { children: item.header }), _jsx(SheetDescription, { children: "Showing total visitors for the last 6 months" })] }), _jsxs("div", { className: "flex flex-1 flex-col gap-4 overflow-y-auto py-4 text-sm", children: [!isMobile && (_jsxs(_Fragment, { children: [_jsx(ChartContainer, { config: chartConfig, children: _jsxs(AreaChart, { accessibilityLayer: true, data: chartData, margin: {
                                                left: 0,
                                                right: 10,
                                            }, children: [_jsx(CartesianGrid, { vertical: false }), _jsx(XAxis, { dataKey: "month", tickLine: false, axisLine: false, tickMargin: 8, tickFormatter: (value) => value.slice(0, 3), hide: true }), _jsx(ChartTooltip, { cursor: false, content: _jsx(ChartTooltipContent, { indicator: "dot" }) }), _jsx(Area, { dataKey: "mobile", type: "natural", fill: "var(--color-mobile)", fillOpacity: 0.6, stroke: "var(--color-mobile)", stackId: "a" }), _jsx(Area, { dataKey: "desktop", type: "natural", fill: "var(--color-desktop)", fillOpacity: 0.4, stroke: "var(--color-desktop)", stackId: "a" })] }) }), _jsx(Separator, {}), _jsxs("div", { className: "grid gap-2", children: [_jsxs("div", { className: "flex gap-2 font-medium leading-none", children: ["Trending up by 5.2% this month", " ", _jsx(TrendingUpIcon, { className: "size-4" })] }), _jsx("div", { className: "text-muted-foreground", children: "Showing total visitors for the last 6 months. This is just some random text to test the layout. It spans multiple lines and should wrap around." })] }), _jsx(Separator, {})] })), _jsxs("form", { className: "flex flex-col gap-4", children: [_jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "header", children: "Header" }), _jsx(Input, { id: "header", defaultValue: item.header })] }), _jsxs("div", { className: "grid grid-cols-2 gap-4", children: [_jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "type", children: "Type" }), _jsxs(Select, { defaultValue: item.type, children: [_jsx(SelectTrigger, { id: "type", className: "w-full", children: _jsx(SelectValue, { placeholder: "Select a type" }) }), _jsxs(SelectContent, { children: [_jsx(SelectItem, { value: "Table of Contents", children: "Table of Contents" }), _jsx(SelectItem, { value: "Executive Summary", children: "Executive Summary" }), _jsx(SelectItem, { value: "Technical Approach", children: "Technical Approach" }), _jsx(SelectItem, { value: "Design", children: "Design" }), _jsx(SelectItem, { value: "Capabilities", children: "Capabilities" }), _jsx(SelectItem, { value: "Focus Documents", children: "Focus Documents" }), _jsx(SelectItem, { value: "Narrative", children: "Narrative" }), _jsx(SelectItem, { value: "Cover Page", children: "Cover Page" })] })] })] }), _jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "status", children: "Status" }), _jsxs(Select, { defaultValue: item.status, children: [_jsx(SelectTrigger, { id: "status", className: "w-full", children: _jsx(SelectValue, { placeholder: "Select a status" }) }), _jsxs(SelectContent, { children: [_jsx(SelectItem, { value: "Done", children: "Done" }), _jsx(SelectItem, { value: "In Progress", children: "In Progress" }), _jsx(SelectItem, { value: "Not Started", children: "Not Started" })] })] })] })] }), _jsxs("div", { className: "grid grid-cols-2 gap-4", children: [_jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "target", children: "Target" }), _jsx(Input, { id: "target", defaultValue: item.target })] }), _jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "limit", children: "Limit" }), _jsx(Input, { id: "limit", defaultValue: item.limit })] })] }), _jsxs("div", { className: "flex flex-col gap-3", children: [_jsx(Label, { htmlFor: "reviewer", children: "Reviewer" }), _jsxs(Select, { defaultValue: item.reviewer, children: [_jsx(SelectTrigger, { id: "reviewer", className: "w-full", children: _jsx(SelectValue, { placeholder: "Select a reviewer" }) }), _jsxs(SelectContent, { children: [_jsx(SelectItem, { value: "Eddie Lake", children: "Eddie Lake" }), _jsx(SelectItem, { value: "Jamik Tashpulatov", children: "Jamik Tashpulatov" }), _jsx(SelectItem, { value: "Emily Whalen", children: "Emily Whalen" })] })] })] })] })] }), _jsxs(SheetFooter, { className: "mt-auto flex gap-2 sm:flex-col sm:space-x-0", children: [_jsx(Button, { className: "w-full", children: "Submit" }), _jsx(SheetClose, { asChild: true, children: _jsx(Button, { variant: "outline", className: "w-full", children: "Done" }) })] })] })] }));
}
