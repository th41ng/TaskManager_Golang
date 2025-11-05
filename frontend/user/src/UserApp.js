import { jsx as _jsx, jsxs as _jsxs } from "react/jsx-runtime";
import { Routes, Route } from "react-router-dom";
import { PublicClientApplication } from '@azure/msal-browser';
import { MsalProvider } from '@azure/msal-react';
import { msalConfig } from './authConfig';
import LoginPage from "./app/login/page";
// Khởi tạo MSAL instance
const msalInstance = new PublicClientApplication(msalConfig);
// Component này được export cho Module Federation
// Không có BrowserRouter vì shell đã có rồi
function UserApp() {
    return (_jsx(MsalProvider, { instance: msalInstance, children: _jsxs(Routes, { children: [_jsx(Route, { path: "/", element: _jsx(LoginPage, {}) }), _jsx(Route, { path: "/login", element: _jsx(LoginPage, {}) })] }) }));
}
export default UserApp;
