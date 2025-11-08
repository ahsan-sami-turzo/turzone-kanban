<template>
  <div class="min-h-full flex flex-col pt-4">
    <div class="flex items-start justify-between pb-6 border-b border-gray-100 mb-8">
      <div>
        <h2 class="text-3xl font-bold text-gray-800 tracking-tight mb-1">{{ project?.title }}</h2>
        <p class="text-base text-gray-500 max-w-2xl">{{ project?.description }}</p>
      </div>

      <button
        @click="goBack"
        class="inline-flex items-center text-sm font-medium px-4 py-2 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 text-gray-700 transition-all duration-200 shadow-sm hover:shadow-md"
      >
        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path></svg>
        Back to Dashboard
      </button>
    </div>

    <div class="flex-1">
      <Board v-if="project" :columns="project.columns" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Board from '../components/Board.vue';
import api from '../services/api';

const route = useRoute();
const router = useRouter();
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

const goBack = () => {
  router.push('/');
};
</script>