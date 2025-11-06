<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-gray-900 mb-2">Turzone Kanban</h1>
        <p class="text-gray-600">Manage your workflow projects</p>
      </div>
      
      <!-- Upload Section -->
      <div class="mb-8">
        <FileUpload @upload-success="loadProjects" />
      </div>
      
      <!-- Projects Grid -->
      <div v-if="loading" class="flex items-center justify-center py-16">
        <svg class="animate-spin h-8 w-8 text-blue-500" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span class="ml-3 text-gray-600">Loading projects...</span>
      </div>
      
      <div v-else-if="projects.length === 0" class="text-center py-16">
        <svg class="mx-auto h-16 w-16 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m-3-3v6m5 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
        </svg>
        <h3 class="text-xl font-medium text-gray-700 mb-2">No projects yet</h3>
        <p class="text-gray-500">Upload a markdown file to create your first project</p>
      </div>
      
      <div v-else>
        <h2 class="text-2xl font-bold text-gray-800 mb-4">Projects</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <ProjectCard
            v-for="project in projects"
            :key="project.id"
            :project="project"
            :stats="projectStats[project.id] || defaultStats"
            @select="selectProject"
            @delete="deleteProject"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import FileUpload from '../components/FileUpload.vue';
import ProjectCard from '../components/ProjectCard.vue';
import api from '../services/api';

const router = useRouter();

const projects = ref([]);
const projectStats = ref({});
const loading = ref(false);

const defaultStats = {
  total: 0,
  waiting: 0,
  in_progress: 0,
  testing: 0,
  completed: 0,
};

const loadProjects = async () => {
  loading.value = true;
  try {
    const response = await api.getProjects();
    projects.value = response.data;
    
    // Load stats for each project
    for (const project of projects.value) {
      await loadProjectStats(project.id);
    }
  } catch (error) {
    console.error('Failed to load projects:', error);
  } finally {
    loading.value = false;
  }
};

const loadProjectStats = async (projectId) => {
  try {
    const response = await api.getProjectTasks(projectId);
    const tasks = response.data;
    
    projectStats.value[projectId] = {
      total: tasks.length,
      waiting: tasks.filter(t => t.status === 'waiting').length,
      in_progress: tasks.filter(t => t.status === 'in_progress').length,
      testing: tasks.filter(t => t.status === 'testing').length,
      completed: tasks.filter(t => t.status === 'completed').length,
    };
  } catch (error) {
    console.error(`Failed to load stats for project ${projectId}:`, error);
  }
};

const selectProject = (project) => {
  router.push(`/project/${project.id}`);
};

const deleteProject = async (projectId) => {
  if (!confirm('Are you sure you want to delete this project and all its tasks?')) return;
  
  try {
    await api.deleteProject(projectId);
    await loadProjects();
  } catch (error) {
    alert('Failed to delete project');
  }
};

onMounted(() => {
  loadProjects();
});
</script>
