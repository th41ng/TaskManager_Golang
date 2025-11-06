import { jsx as _jsx } from "react/jsx-runtime";
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { PublicClientApplication } from '@azure/msal-browser';
import { MsalProvider } from '@azure/msal-react';
import { msalConfig } from './authConfig';
import './index.css';
import UserApp from './UserApp';
// Khởi tạo MSAL instance
const msalInstance = new PublicClientApplication(msalConfig);
// Standalone mode - wrap UserApp với BrowserRouter
createRoot(document.getElementById('root')).render(_jsx(StrictMode, { children: _jsx(MsalProvider, { instance: msalInstance, children: _jsx(BrowserRouter, { children: _jsx(UserApp, {}) }) }) }));
