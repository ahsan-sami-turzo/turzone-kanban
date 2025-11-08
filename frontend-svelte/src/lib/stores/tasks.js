import { writable } from 'svelte/store'
import { tasksAPI } from '../api/client.js'

function createTasksStore() {
  const { subscribe, set, update } = writable({})
  const loading = writable(false)
  const error = writable(null)

  return {
    subscribe,
    loading: { subscribe: loading.subscribe },
    error: { subscribe: error.subscribe },

    fetchByProject: async (projectId) => {
      loading.set(true)
      error.set(null)
      try {
        const response = await tasksAPI.getByProject(projectId)
        update(tasks => ({
          ...tasks,
          [projectId]: response.data
        }))
        return response.data
      } catch (err) {
        const errorMessage = err.response?.data?.error || 'Failed to fetch tasks'
        error.set(errorMessage)
        throw new Error(errorMessage)
      } finally {
        loading.set(false)
      }
    },

    updateTask: async (taskId, updates) => {
      loading.set(true)
      error.set(null)
      try {
        const response = await tasksAPI.update(taskId, updates)
        
        update(tasks => {
          const newTasks = { ...tasks }
          Object.keys(newTasks).forEach(projectId => {
            const taskIndex = newTasks[projectId].findIndex(t => t.id === taskId)
            if (taskIndex >= 0) {
              newTasks[projectId][taskIndex] = { ...newTasks[projectId][taskIndex], ...response.data }
            }
          })
          return newTasks
        })
        
        return response.data
      } catch (err) {
        const errorMessage = err.response?.data?.error || 'Failed to update task'
        error.set(errorMessage)
        throw new Error(errorMessage)
      } finally {
        loading.set(false)
      }
    },

    deleteTask: async (taskId, projectId) => {
      loading.set(true)
      error.set(null)
      try {
        await tasksAPI.delete(taskId)
        update(tasks => ({
          ...tasks,
          [projectId]: (tasks[projectId] || []).filter(t => t.id !== taskId)
        }))
      } catch (err) {
        const errorMessage = err.response?.data?.error || 'Failed to delete task'
        error.set(errorMessage)
        throw new Error(errorMessage)
      } finally {
        loading.set(false)
      }
    },

    clearError: () => error.set(null),
    clearTasks: () => set({})
  }
}

export const tasks = createTasksStore()