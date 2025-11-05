import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { GalleryVerticalEnd } from "lucide-react";
import { LoginForm } from "@/components/login-form";
export default function LoginPage() {
    return (_jsxs("div", { className: "grid min-h-svh lg:grid-cols-2", children: [_jsxs("div", { className: "flex flex-col gap-4 p-6 md:p-10", children: [_jsx("div", { className: "flex justify-center gap-2 md:justify-start", children: _jsxs("a", { href: "#", className: "flex items-center gap-2 font-medium", children: [_jsx("div", { className: "bg-primary text-primary-foreground flex size-6 items-center justify-center rounded-md", children: _jsx(GalleryVerticalEnd, { className: "size-4" }) }), "Acme Inc."] }) }), _jsx("div", { className: "flex flex-1 items-center justify-center", children: _jsx("div", { className: "w-full max-w-xs", children: _jsx(LoginForm, {}) }) })] }), _jsx("div", { className: "bg-muted relative hidden lg:block", children: _jsx("img", { src: "/placeholder.svg", alt: "Image", className: "absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale" }) })] }));
}
