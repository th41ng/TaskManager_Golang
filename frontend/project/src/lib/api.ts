// API Base URL configuration
export const API_BASE_URL = (import.meta.env.VITE_API_BASE as string) || 'http://172.21.223.107:8080/api'

// Helper to get auth token from localStorage
export const getAuthHeaders = (): HeadersInit => {
  const token = localStorage.getItem('token')
  return {
    'Content-Type': 'application/json',
    ...(token && { 'Authorization': `Bearer ${token}` })
  }
}
