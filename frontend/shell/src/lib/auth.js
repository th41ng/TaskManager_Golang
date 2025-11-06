// Auth utilities
export const auth = {
    // Get current token
    getToken() {
        return localStorage.getItem('token') || undefined;
    },
    // Get current user
    getUser() {
        const userStr = localStorage.getItem('user');
        if (!userStr)
            return undefined;
        try {
            return JSON.parse(userStr);
        }
        catch {
            return undefined;
        }
    },
    // Check if user is logged in
    isAuthenticated() {
        return !!this.getToken();
    },
    // Logout - clear localStorage
    logout() {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
    },
};
