import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default {
  // Projects
  getProjects() {
    return api.get('/projects');
  },
  
  getProject(id) {
    return api.get(`/projects/${id}`);
  },
  
  uploadMarkdown(file) {
    const formData = new FormData();
    formData.append('file', file);
    return api.post('/projects/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
  },
  
  deleteProject(id) {
    return api.delete(`/projects/${id}`);
  },
  
  // Tasks
  getProjectTasks(projectId) {
    return api.get(`/projects/${projectId}/tasks`);
  },
  
  updateTask(taskId, data) {
    return api.put(`/tasks/${taskId}`, data);
  },
  
  deleteTask(taskId) {
    return api.delete(`/tasks/${taskId}`);
  },
};