import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { Routes, Route } from "react-router-dom";
import { CookiesProvider } from "react-cookie";
import LoginPage from "./app/login/page";
// Component này được export cho Module Federation
// Không có BrowserRouter vì shell đã có rồi
function UserApp() {
    return (_jsx(CookiesProvider, { children: _jsxs(Routes, { children: [_jsx(Route, { path: "/", element: _jsx(LoginPage, {}) }), _jsx(Route, { path: "/login", element: _jsx(LoginPage, {}) })] }) }));
}
export default UserApp;
