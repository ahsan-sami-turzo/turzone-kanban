<template>
  <div>
    <div v-if="loading" class="flex items-center justify-center py-16">
      <svg class="animate-spin h-8 w-8 text-blue-500" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span class="ml-3 text-gray-600">Loading tasks...</span>
    </div>
    
    <div v-else-if="!projectId" class="text-center py-16 text-gray-500">
      <svg class="mx-auto h-16 w-16 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
      </svg>
      <p class="text-lg">Select a project to view its board</p>
    </div>
    
    <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
      Error loading tasks: {{ error }}
    </div>
    
    <div v-else>
      <div v-if="tasks.length === 0" class="text-center py-16 text-gray-500">
        <p class="text-lg">No tasks in this project yet</p>
        <p class="text-sm mt-2">Upload a markdown file to add tasks</p>
      </div>
      
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <Column
          title="Waiting"
          status="waiting"
          :tasks="tasksByStatus.waiting"
          @task-moved="handleTaskMoved"
          @edit-task="openEditModal"
          @delete-task="handleDeleteTask"
        />
        
        <Column
          title="In Progress"
          status="in_progress"
          :tasks="tasksByStatus.in_progress"
          @task-moved="handleTaskMoved"
          @edit-task="openEditModal"
          @delete-task="handleDeleteTask"
        />
        
        <Column
          title="Testing"
          status="testing"
          :tasks="tasksByStatus.testing"
          @task-moved="handleTaskMoved"
          @edit-task="openEditModal"
          @delete-task="handleDeleteTask"
        />
        
        <Column
          title="Completed"
          status="completed"
          :tasks="tasksByStatus.completed"
          @task-moved="handleTaskMoved"
          @edit-task="openEditModal"
          @delete-task="handleDeleteTask"
        />
      </div>
    </div>
    
    <!-- Edit Modal (same as before) -->
    <div
      v-if="editingTask"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      @click.self="closeEditModal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md shadow-xl">
        <h3 class="text-xl font-bold mb-4">Edit Task</h3>
        
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              Title
            </label>
            <input
              v-model="editForm.title"
              type="text"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Task title"
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              Description
            </label>
            <textarea
              v-model="editForm.description"
              rows="6"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Task description..."
            ></textarea>
          </div>
          
          <div class="flex justify-end space-x-2 pt-4">
            <button
              @click="closeEditModal"
              class="px-4 py-2 text-gray-700 hover:bg-gray-100 rounded-lg transition"
            >
              Cancel
            </button>
            <button
              @click="saveEdit"
              class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
              :disabled="!editForm.title.trim()"
            >
              Save Changes
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue';
import Column from './Column.vue';
import api from '../services/api';

const props = defineProps({
  projectId: {
    type: Number,
    default: null,
  },
});

const tasks = ref([]);
const loading = ref(false);
const error = ref('');
const editingTask = ref(null);
const editForm = ref({ title: '', description: '' });

const tasksByStatus = computed(() => {
  return {
    waiting: tasks.value.filter(t => t.status === 'waiting'),
    in_progress: tasks.value.filter(t => t.status === 'in_progress'),
    testing: tasks.value.filter(t => t.status === 'testing'),
    completed: tasks.value.filter(t => t.status === 'completed'),
  };
});

const loadTasks = async () => {
  if (!props.projectId) return;
  
  loading.value = true;
  error.value = '';
  try {
    const response = await api.getProjectTasks(props.projectId);
    tasks.value = response.data;
  } catch (err) {
    error.value = err.response?.data?.error || 'Failed to load tasks';
    console.error('Failed to load tasks:', err);
  } finally {
    loading.value = false;
  }
};

const handleTaskMoved = async ({ taskId, newStatus, newPosition }) => {
  try {
    await api.updateTask(taskId, {
      status: newStatus,
      position: newPosition,
    });
  } catch (err) {
    console.error('Failed to update task:', err);
    await loadTasks();
  }
};

const openEditModal = (task) => {
  editingTask.value = task;
  editForm.value = {
    title: task.title,
    description: task.description || '',
  };
};

const closeEditModal = () => {
  editingTask.value = null;
  editForm.value = { title: '', description: '' };
};

const saveEdit = async () => {
  if (!editForm.value.title.trim()) return;
  
  try {
    await api.updateTask(editingTask.value.id, editForm.value);
    
    const task = tasks.value.find(t => t.id === editingTask.value.id);
    if (task) {
      task.title = editForm.value.title;
      task.description = editForm.value.description;
    }
    
    closeEditModal();
  } catch (err) {
    alert('Failed to update task: ' + (err.response?.data?.error || 'Unknown error'));
  }
};

const handleDeleteTask = async (taskId) => {
  if (!confirm('Are you sure you want to delete this task?')) return;
  
  try {
    await api.deleteTask(taskId);
    tasks.value = tasks.value.filter(t => t.id !== taskId);
  } catch (err) {
    alert('Failed to delete task: ' + (err.response?.data?.error || 'Unknown error'));
  }
};

watch(() => props.projectId, () => {
  loadTasks();
}, { immediate: true });
</script>