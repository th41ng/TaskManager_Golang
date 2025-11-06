import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { MailIcon, PlusCircleIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import { SidebarGroup, SidebarGroupContent, SidebarMenu, SidebarMenuButton, SidebarMenuItem, } from "@/components/ui/sidebar";
export function NavMain({ items, }) {
    return (_jsx(SidebarGroup, { children: _jsxs(SidebarGroupContent, { className: "flex flex-col gap-2", children: [_jsx(SidebarMenu, { children: _jsxs(SidebarMenuItem, { className: "flex items-center gap-2", children: [_jsxs(SidebarMenuButton, { tooltip: "Quick Create", className: "min-w-8 bg-primary text-primary-foreground duration-200 ease-linear hover:bg-primary/90 hover:text-primary-foreground active:bg-primary/90 active:text-primary-foreground", children: [_jsx(PlusCircleIcon, {}), _jsx("span", { children: "Quick Create" })] }), _jsxs(Button, { size: "icon", className: "h-9 w-9 shrink-0 group-data-[collapsible=icon]:opacity-0", variant: "outline", children: [_jsx(MailIcon, {}), _jsx("span", { className: "sr-only", children: "Inbox" })] })] }) }), _jsx(SidebarMenu, { children: items.map((item) => (_jsx(SidebarMenuItem, { children: _jsxs(SidebarMenuButton, { tooltip: item.title, children: [item.icon && _jsx(item.icon, {}), _jsx("span", { children: item.title })] }) }, item.title))) })] }) }));
}
