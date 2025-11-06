<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Header with back button -->
      <div class="mb-8">
        <button
          @click="$router.push('/')"
          class="flex items-center text-blue-600 hover:text-blue-700 mb-4"
        >
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          Back to Projects
        </button>
        
        <h1 class="text-3xl font-bold text-gray-900">{{ project?.name || 'Loading...' }}</h1>
      </div>
      
      <!-- Board -->
      <Board :project-id="projectId" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import Board from '../components/Board.vue';
import api from '../services/api';

const route = useRoute();
const projectId = parseInt(route.params.id);
const project = ref(null);

onMounted(async () => {
  try {
    const response = await api.getProject(projectId);
    project.value = response.data;
  } catch (error) {
    console.error('Failed to load project:', error);
  }
});
</script>