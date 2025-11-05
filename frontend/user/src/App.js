import { jsx as _jsx } from "react/jsx-runtime";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import LoginPage from "./app/login/page";
function App() {
    return (_jsx(BrowserRouter, { children: _jsx(Routes, { children: _jsx(Route, { path: "/", element: _jsx(LoginPage, {}) }) }) }));
}
export default App;
