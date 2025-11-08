<template>
  <div class="flex flex-col h-full">
    <div v-if="!loading && projectId" 
         class="flex-1 overflow-x-auto pb-4 -mx-4 sm:-mx-6 lg:-mx-8">
      <div class="inline-flex space-x-6 h-full px-4 sm:px-6 lg:px-8">
        <Column
          v-for="col in columnsWithTasks"
          :key="col.status"
          :title="col.title"
          :status="col.status"
          :tasks="col.tasks"
          @task-moved="handleTaskMoved"
          @edit-task="openEditModal"
          @delete-task="handleDeleteTask"
        />
      </div>
    </div>
    
    <div 
      v-if="editingTask" 
      class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-30 backdrop-blur-sm transition-opacity" 
      @click.self="closeEditModal"
    >
      <div class="bg-white rounded-xl p-8 shadow-2xl w-full max-w-lg transform transition-all duration-300 scale-100">
        <div class="flex justify-between items-center mb-6">
          <h3 class="text-2xl font-bold text-gray-800">Edit Task: {{ editingTask.title }}</h3>
          <button @click="closeEditModal" class="text-gray-400 hover:text-gray-700 p-1.5 rounded-full hover:bg-gray-100 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
          </button>
        </div>

        <form @submit.prevent="saveEdit" class="space-y-4">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700 mb-1">Title</label>
            <input
              id="title"
              v-model="editForm.title"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-indigo-500 focus:border-indigo-500 transition duration-150"
              required
            />
          </div>
          <div>
            <label for="description" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
            <textarea
              id="description"
              v-model="editForm.description"
              rows="4"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-indigo-500 focus:border-indigo-500 transition duration-150"
            ></textarea>
          </div>
          
          <div class="pt-4 flex justify-end space-x-3">
            <button
              type="button"
              @click="closeEditModal"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-lg hover:bg-indigo-700 transition-colors shadow-md"
            >
              Save Changes
            </button>
          </div>
        </form>
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