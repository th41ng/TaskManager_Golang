import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { Search } from "lucide-react";
import { Label } from "@/components/ui/label";
import { SidebarGroup, SidebarGroupContent, SidebarInput, } from "@/components/ui/sidebar";
export function SearchForm({ ...props }) {
    return (_jsx("form", { ...props, children: _jsx(SidebarGroup, { className: "py-0", children: _jsxs(SidebarGroupContent, { className: "relative", children: [_jsx(Label, { htmlFor: "search", className: "sr-only", children: "Search" }), _jsx(SidebarInput, { id: "search", placeholder: "Search the docs...", className: "pl-8" }), _jsx(Search, { className: "pointer-events-none absolute left-2 top-1/2 size-4 -translate-y-1/2 select-none opacity-50" })] }) }) }));
}
