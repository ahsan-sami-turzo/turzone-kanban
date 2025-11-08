import { writable } from 'svelte/store'
import { projectsAPI } from '../api/client.js'

function createProjectsStore() {
  const { subscribe, set, update } = writable([])
  const loading = writable(false)
  const error = writable(null)

  return {
    subscribe,
    loading: { subscribe: loading.subscribe },
    error: { subscribe: error.subscribe },

    fetchAll: async () => {
      loading.set(true)
      error.set(null)
      try {
        const response = await projectsAPI.getAll()
        set(response.data)
      } catch (err) {
        error.set(err.response?.data?.error || 'Failed to fetch projects')
      } finally {
        loading.set(false)
      }
    },

    uploadProject: async (file) => {
      loading.set(true)
      error.set(null)
      try {
        console.log('Store: Starting upload for file:', file.name)
        
        const response = await projectsAPI.upload(file)
        console.log('Store: Upload successful', response.data)
        
        update(projects => {
          const existingIndex = projects.findIndex(p => p.id === response.data.project.id)
          if (existingIndex >= 0) {
            // Update existing project
            projects[existingIndex] = response.data.project
            return [...projects]
          } else {
            // Add new project
            return [...projects, response.data.project]
          }
        })
        
        return response.data
      } catch (err) {
        const errorMessage = err.response?.data?.error || `Upload failed: ${err.message}`
        console.error('Store: Upload error', errorMessage)
        error.set(errorMessage)
        throw err
      } finally {
        loading.set(false)
      }
    },

    deleteProject: async (id) => {
      loading.set(true)
      error.set(null)
      try {
        await projectsAPI.delete(id)
        update(projects => projects.filter(p => p.id !== id))
      } catch (err) {
        error.set(err.response?.data?.error || 'Failed to delete project')
        throw err
      } finally {
        loading.set(false)
      }
    },

    clearError: () => error.set(null)
  }
}

export const projects = createProjectsStore()