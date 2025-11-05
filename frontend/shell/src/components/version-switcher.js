import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import * as React from "react";
import { Check, ChevronsUpDown, GalleryVerticalEnd } from "lucide-react";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger, } from "@/components/ui/dropdown-menu";
import { SidebarMenu, SidebarMenuButton, SidebarMenuItem, } from "@/components/ui/sidebar";
export function VersionSwitcher({ versions, defaultVersion, }) {
    const [selectedVersion, setSelectedVersion] = React.useState(defaultVersion);
    return (_jsx(SidebarMenu, { children: _jsx(SidebarMenuItem, { children: _jsxs(DropdownMenu, { children: [_jsx(DropdownMenuTrigger, { asChild: true, children: _jsxs(SidebarMenuButton, { size: "lg", className: "data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground", children: [_jsx("div", { className: "flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground", children: _jsx(GalleryVerticalEnd, { className: "size-4" }) }), _jsxs("div", { className: "flex flex-col gap-0.5 leading-none", children: [_jsx("span", { className: "font-semibold", children: "Documentation" }), _jsxs("span", { className: "", children: ["v", selectedVersion] })] }), _jsx(ChevronsUpDown, { className: "ml-auto" })] }) }), _jsx(DropdownMenuContent, { className: "w-[--radix-dropdown-menu-trigger-width]", align: "start", children: versions.map((version) => (_jsxs(DropdownMenuItem, { onSelect: () => setSelectedVersion(version), children: ["v", version, " ", version === selectedVersion && _jsx(Check, { className: "ml-auto" })] }, version))) })] }) }) }));
}
