<template>
  <div class="mb-6">
    <h2 class="text-2xl font-bold mb-4">Projects</h2>
    
    <div v-if="loading" class="text-gray-500">Loading projects...</div>
    
    <div v-else-if="projects.length === 0" class="text-gray-500 italic">
      No projects yet. Upload a markdown file to create one.
    </div>
    
    <div v-else class="space-y-2">
      <div
        v-for="project in projects"
        :key="project.id"
        @click="selectProject(project)"
        class="flex items-center justify-between p-4 border rounded-lg cursor-pointer transition"
        :class="selectedProjectId === project.id ? 'bg-blue-50 border-blue-500' : 'hover:bg-gray-50'"
      >
        <div>
          <h3 class="font-semibold">{{ project.name }}</h3>
          <p class="text-sm text-gray-500">
            Created: {{ formatDate(project.created_at) }}
          </p>
        </div>
        
        <button
          @click.stop="deleteProject(project.id)"
          class="text-red-500 hover:text-red-700 px-3 py-1 rounded hover:bg-red-50"
        >
          Delete
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