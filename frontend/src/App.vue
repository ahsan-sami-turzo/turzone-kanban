<template>
  <div class="min-h-screen bg-gray-100">
    <div class="container mx-auto px-4 py-8">
      <h1 class="text-4xl font-bold mb-8 text-gray-800">Turzone Kanban</h1>
      
      <FileUpload @upload-success="handleUploadSuccess" />
      
      <ProjectList 
        ref="projectListRef"
        @project-selected="handleProjectSelected" 
      />
      
      <div v-if="selectedProject" class="mt-8">
        <h2 class="text-2xl font-bold mb-4">{{ selectedProject.name }} - Board</h2>
        <p class="text-gray-600 mb-4">Kanban board will appear here in the next step</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import FileUpload from './components/FileUpload.vue';
import ProjectList from './components/ProjectList.vue';

const projectListRef = ref(null);
const selectedProject = ref(null);

const handleUploadSuccess = () => {
  // Reload projects after successful upload
  projectListRef.value?.loadProjects();
};

const handleProjectSelected = (project) => {
  selectedProject.value = project;
};
</script>