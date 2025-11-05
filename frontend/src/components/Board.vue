<template>
  <div>
    <div v-if="loading" class="text-center py-8 text-gray-500">
      Loading tasks...
    </div>
    
    <div v-else-if="!projectId" class="text-center py-8 text-gray-500 italic">
      Select a project to view its board
    </div>
    
    <div v-else class="grid grid-cols-4 gap-4">
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
    
    <!-- Edit Modal -->
    <div
      v-if="editingTask"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click.self="closeEditModal"
    >
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
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
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">
              Description
            </label>
            <textarea
              v-model="editForm.description"
              rows="4"
              class="w-full border border-gray-300 rounded-lg px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            ></textarea>
          </div>
          
          <div class="flex justify-end space-x-2">
            <button
              @click="closeEditModal"
              class="px-4 py-2 text-gray-700 hover:bg-gray-100 rounded-lg"
            >
              Cancel
            </button>
            <button
              @click="saveEdit"
              class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
            >
              Save
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
  try {
    const response = await api.getProjectTasks(props.projectId);
    tasks.value = response.data;
  } catch (error) {
    console.error('Failed to load tasks:', error);
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
    // Optimistic update - task is already moved in UI by draggable
  } catch (error) {
    console.error('Failed to update task:', error);
    // Reload tasks on error
    await loadTasks();
  }
};

const openEditModal = (task) => {
  editingTask.value = task;
  editForm.value = {
    title: task.title,
    description: task.description,
  };
};

const closeEditModal = () => {
  editingTask.value = null;
  editForm.value = { title: '', description: '' };
};

const saveEdit = async () => {
  try {
    await api.updateTask(editingTask.value.id, editForm.value);
    
    // Update local task
    const task = tasks.value.find(t => t.id === editingTask.value.id);
    if (task) {
      task.title = editForm.value.title;
      task.description = editForm.value.description;
    }
    
    closeEditModal();
  } catch (error) {
    alert('Failed to update task');
  }
};

const handleDeleteTask = async (taskId) => {
  if (!confirm('Are you sure you want to delete this task?')) return;
  
  try {
    await api.deleteTask(taskId);
    tasks.value = tasks.value.filter(t => t.id !== taskId);
  } catch (error) {
    alert('Failed to delete task');
  }
};

// Watch for project changes
watch(() => props.projectId, () => {
  loadTasks();
}, { immediate: true });
</script>