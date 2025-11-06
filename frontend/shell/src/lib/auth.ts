// Auth utilities

export const auth = {
  // Get current token
  getToken(): string | undefined {
    return localStorage.getItem('token') || undefined
  },

  // Get current user
  getUser(): any | undefined {
    const userStr = localStorage.getItem('user')
    if (!userStr) return undefined
    try {
      return JSON.parse(userStr)
    } catch {
      return undefined
    }
  },

  // Check if user is logged in
  isAuthenticated(): boolean {
    return !!this.getToken()
  },

  // Logout - clear localStorage
  logout(): void {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  },
}
