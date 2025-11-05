import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { lazy, Suspense } from "react";
import Page from "@/app/dashboard/page";
// Import UserApp từ remote
const UserApp = lazy(() => import("userApp/UserApp"));
export default function App() {
    return (_jsx(BrowserRouter, { children: _jsxs(Routes, { children: [_jsx(Route, { path: "/", element: _jsx(Page, {}) }), _jsx(Route, { path: "/login", element: _jsx(Suspense, { fallback: _jsx("div", { children: "Loading..." }), children: _jsx(UserApp, {}) }) }), _jsx(Route, { path: "/users/*", element: _jsx(Suspense, { fallback: _jsx("div", { children: "Loading..." }), children: _jsx(UserApp, {}) }) })] }) }));
}
