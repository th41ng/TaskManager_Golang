import { Configuration } from "@azure/msal-browser"

export const msalConfig: Configuration = {
  auth: {
    clientId: "YOUR_CLIENT_ID", // TODO: Thay bằng Client ID từ Azure Portal
    authority: "https://login.microsoftonline.com/common", // hoặc /organizations cho work accounts
    redirectUri: "http://localhost:5173", // URL của app khi chạy dev
  },
  cache: {
    cacheLocation: "localStorage", // Lưu token trong localStorage
    storeAuthStateInCookie: false, // Set true nếu dùng IE11 hoặc Edge
  },
}

export const loginRequest = {
  scopes: ["User.Read", "email", "openid", "profile"], // Quyền cần thiết
}

// Thêm config cho Graph API nếu cần
export const graphConfig = {
  graphMeEndpoint: "https://graph.microsoft.com/v1.0/me",
}
