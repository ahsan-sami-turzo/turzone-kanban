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
        <h2 class="text-2xl font-bold mb-4">{{ selectedProject.name }}</h2>
        <Board :project-id="selectedProject.id" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import FileUpload from './components/FileUpload.vue';
import ProjectList from './components/ProjectList.vue';
import Board from './components/Board.vue';

const projectListRef = ref(null);
const selectedProject = ref(null);

const handleUploadSuccess = () => {
  projectListRef.value?.loadProjects();
};

const handleProjectSelected = (project) => {
  selectedProject.value = project;
};
</script>