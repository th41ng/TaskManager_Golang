const API_BASE_URL = import.meta.env.VITE_API_BASE || 'http://172.21.223.107:8080/api';
// Helper function to get auth headers
function getAuthHeaders() {
    const token = localStorage.getItem('token');
    return {
        'Content-Type': 'application/json',
        ...(token && { 'Authorization': `Bearer ${token}` })
    };
}
// Helper function to handle API errors
async function handleResponse(response) {
    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Unknown error' }));
        throw new Error(error.message || `HTTP Error: ${response.status}`);
    }
    return response.json();
}
// Projects API
export const projectsApi = {
    async list() {
        const response = await fetch(`${API_BASE_URL}/projects`, {
            headers: getAuthHeaders()
        });
        const data = await handleResponse(response);
        return data.projects || [];
    },
    async get(id) {
        const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
            headers: getAuthHeaders()
        });
        return handleResponse(response);
    },
    async create(data) {
        const response = await fetch(`${API_BASE_URL}/projects`, {
            method: 'POST',
            headers: getAuthHeaders(),
            body: JSON.stringify(data),
        });
        return handleResponse(response);
    },
    async update(id, data) {
        const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
            method: 'PUT',
            headers: getAuthHeaders(),
            body: JSON.stringify({ id, ...data }),
        });
        return handleResponse(response);
    },
    async delete(id) {
        const response = await fetch(`${API_BASE_URL}/projects/${id}`, {
            method: 'DELETE',
            headers: getAuthHeaders(),
        });
        await handleResponse(response);
    },
};
// Tasks API
export const tasksApi = {
    async list() {
        const response = await fetch(`${API_BASE_URL}/tasks`, {
            headers: getAuthHeaders()
        });
        const data = await handleResponse(response);
        return data.tasks || [];
    },
    async get(id) {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
            headers: getAuthHeaders()
        });
        return handleResponse(response);
    },
    async create(data) {
        const response = await fetch(`${API_BASE_URL}/tasks`, {
            method: 'POST',
            headers: getAuthHeaders(),
            body: JSON.stringify(data),
        });
        return handleResponse(response);
    },
    async update(id, data) {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
            method: 'PUT',
            headers: getAuthHeaders(),
            body: JSON.stringify({ id, ...data }),
        });
        return handleResponse(response);
    },
    async toggleComplete(task) {
        const response = await fetch(`${API_BASE_URL}/tasks/${task.id}`, {
            method: 'PUT',
            headers: getAuthHeaders(),
            body: JSON.stringify({
                id: task.id,
                title: task.title,
                done: !task.done,
                priority: task.priority,
                project_id: task.project_id
            }),
        });
        return handleResponse(response);
    },
    async delete(id) {
        const response = await fetch(`${API_BASE_URL}/tasks/${id}`, {
            method: 'DELETE',
            headers: getAuthHeaders(),
        });
        await handleResponse(response);
    },
};
