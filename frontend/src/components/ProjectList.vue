<template>
  <div class="mb-6">
    <h2 class="text-xl font-bold text-gray-800 mb-4 tracking-tight">Available Projects</h2>
    
    <div v-if="loading" class="text-gray-500">Loading projects...</div>
    
    <div v-else-if="projects.length === 0" class="text-gray-500 italic py-4">
      No projects yet. Upload a markdown file to create one.
    </div>
    
    <div v-else class="space-y-3">
      <div
        v-for="project in projects"
        :key="project.id"
        @click="selectProject(project)"
        class="flex items-center justify-between p-4 bg-white rounded-xl shadow-sm cursor-pointer transition-all duration-200 border border-gray-100"
        :class="selectedProjectId === project.id 
          ? 'ring-2 ring-indigo-500 border-indigo-500 shadow-md' 
          : 'hover:shadow-lg hover:border-gray-300 hover:scale-[1.01]'"
      >
        <div class="flex-1 pr-4">
          <h3 class="font-semibold text-gray-800">{{ project.name }}</h3>
          <p class="text-xs text-gray-500 mt-0.5">
            Created: {{ formatDate(project.created_at) }}
          </p>
        </div>
        
        <button
          @click.stop="deleteProject(project.id)"
          class="flex-shrink-0 p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
          title="Delete Project"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import api from '../services/api';

const emit = defineEmits(['project-selected']);

const projects = ref([]);
const loading = ref(false);
const selectedProjectId = ref(null);

const loadProjects = async () => {
  loading.value = true;
  try {
    const response = await api.getProjects();
    projects.value = response.data;
    
    // Auto-select first project if exists
    if (projects.value.length > 0 && !selectedProjectId.value) {
      selectProject(projects.value[0]);
    }
  } catch (error) {
    console.error('Failed to load projects:', error);
  } finally {
    loading.value = false;
  }
};

const selectProject = (project) => {
  selectedProjectId.value = project.id;
  emit('project-selected', project);
};

const deleteProject = async (projectId) => {
  if (!confirm('Are you sure you want to delete this project?')) return;
  
  try {
    await api.deleteProject(projectId);
    await loadProjects();
    
    // Clear selection if deleted project was selected
    if (selectedProjectId.value === projectId) {
      selectedProjectId.value = null;
      emit('project-selected', null);
    }
  } catch (error) {
    alert('Failed to delete project');
  }
};

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString();
};

onMounted(() => {
  loadProjects();
});

defineExpose({ loadProjects });
</script>