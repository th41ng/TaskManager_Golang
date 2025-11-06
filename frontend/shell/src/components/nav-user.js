import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { BellIcon, CreditCardIcon, LogOutIcon, MoreVerticalIcon, UserCircleIcon, } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { auth } from "@/lib/auth";
import { Avatar, AvatarFallback, AvatarImage, } from "@/components/ui/avatar";
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger, } from "@/components/ui/dropdown-menu";
import { SidebarMenu, SidebarMenuButton, SidebarMenuItem, useSidebar, } from "@/components/ui/sidebar";
export function NavUser({ user, }) {
    const { isMobile } = useSidebar();
    const navigate = useNavigate();
    const handleLogout = () => {
        auth.logout();
        navigate('/login');
    };
    return (_jsx(SidebarMenu, { children: _jsx(SidebarMenuItem, { children: _jsxs(DropdownMenu, { children: [_jsx(DropdownMenuTrigger, { asChild: true, children: _jsxs(SidebarMenuButton, { size: "lg", className: "data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground", children: [_jsxs(Avatar, { className: "h-8 w-8 rounded-lg grayscale", children: [_jsx(AvatarImage, { src: user.avatar, alt: user.name }), _jsx(AvatarFallback, { className: "rounded-lg", children: "CN" })] }), _jsxs("div", { className: "grid flex-1 text-left text-sm leading-tight", children: [_jsx("span", { className: "truncate font-medium", children: user.name }), _jsx("span", { className: "truncate text-xs text-muted-foreground", children: user.email })] }), _jsx(MoreVerticalIcon, { className: "ml-auto size-4" })] }) }), _jsxs(DropdownMenuContent, { className: "w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg", side: isMobile ? "bottom" : "right", align: "end", sideOffset: 4, children: [_jsx(DropdownMenuLabel, { className: "p-0 font-normal", children: _jsxs("div", { className: "flex items-center gap-2 px-1 py-1.5 text-left text-sm", children: [_jsxs(Avatar, { className: "h-8 w-8 rounded-lg", children: [_jsx(AvatarImage, { src: user.avatar, alt: user.name }), _jsx(AvatarFallback, { className: "rounded-lg", children: "CN" })] }), _jsxs("div", { className: "grid flex-1 text-left text-sm leading-tight", children: [_jsx("span", { className: "truncate font-medium", children: user.name }), _jsx("span", { className: "truncate text-xs text-muted-foreground", children: user.email })] })] }) }), _jsx(DropdownMenuSeparator, {}), _jsxs(DropdownMenuGroup, { children: [_jsxs(DropdownMenuItem, { children: [_jsx(UserCircleIcon, {}), "Account"] }), _jsxs(DropdownMenuItem, { children: [_jsx(CreditCardIcon, {}), "Billing"] }), _jsxs(DropdownMenuItem, { children: [_jsx(BellIcon, {}), "Notifications"] })] }), _jsx(DropdownMenuSeparator, {}), _jsxs(DropdownMenuItem, { onClick: handleLogout, children: [_jsx(LogOutIcon, {}), "Log out"] })] })] }) }) }));
}
