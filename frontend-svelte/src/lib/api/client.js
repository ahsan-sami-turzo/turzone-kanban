import axios from 'axios'

const API_BASE_URL = '/api'

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Add request interceptor for debugging
apiClient.interceptors.request.use(
  (config) => {
    console.log('Making API request:', config.method?.toUpperCase(), config.url)
    if (config.data instanceof FormData) {
      console.log('Request contains FormData')
      // Log FormData entries for debugging
      for (let [key, value] of config.data.entries()) {
        if (value instanceof File) {
          console.log(`FormData ${key}:`, {
            name: value.name,
            type: value.type,
            size: value.size
          })
        }
      }
    }
    return config
  },
  (error) => {
    console.error('Request error:', error)
    return Promise.reject(error)
  }
)

// Add response interceptor for debugging
apiClient.interceptors.response.use(
  (response) => {
    console.log('API response success:', response.status, response.data)
    return response
  },
  (error) => {
    console.error('API error response:', {
      status: error.response?.status,
      data: error.response?.data,
      message: error.message
    })
    return Promise.reject(error)
  }
)

// Projects API
export const projectsAPI = {
  getAll: () => apiClient.get('/projects'),
  getById: (id) => apiClient.get(`/projects/${id}`),
  
  upload: async (file) => {
  console.log('API: Starting upload for file:', file)
  
  if (!file || !(file instanceof File)) {
    throw new Error('Invalid file provided')
  }
  
  const formData = new FormData()
  formData.append('file', file)
  
  // Debug FormData contents
  console.log('API: FormData contents:')
    for (let [key, value] of formData.entries()) {
      if (value instanceof File) {
        console.log(`  ${key}:`, {
          name: value.name,
          type: value.type,
          size: value.size
        })
      } else {
        console.log(`  ${key}:`, value)
      }
    }
    
    try {
      const response = await apiClient.post('/projects/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
        timeout: 30000
      })
      console.log('API: Upload successful')
      return response
    } catch (error) {
      console.error('API: Upload failed:', error)
      throw error
    }
  },

  delete: (id) => apiClient.delete(`/projects/${id}`)
}

// Tasks API
export const tasksAPI = {
  getByProject: (projectId) => apiClient.get(`/projects/${projectId}/tasks`),
  update: (id, data) => apiClient.put(`/tasks/${id}`, data),
  delete: (id) => apiClient.delete(`/tasks/${id}`)
}

export default apiClient