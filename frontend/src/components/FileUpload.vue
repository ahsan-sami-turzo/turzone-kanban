<template>
  <div class="relative w-full max-w-lg mx-auto bg-white rounded-3xl shadow-2xl p-8">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold text-gray-800">Upload New Project</h2>
      <button @click="$emit('close')" class="text-gray-400 hover:text-gray-600 p-2 transition-colors rounded-full hover:bg-gray-50">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>
      </button>
    </div>
    
    <div 
      class="relative border-3 border-dashed rounded-2xl p-12 text-center transition-all duration-300 ease-in-out"
      :class="isDragging 
        ? 'border-indigo-500 bg-indigo-50/70 scale-[1.01]' 
        : 'border-gray-300 bg-gray-50 hover:border-indigo-400 hover:bg-indigo-50/50'"
      @dragover.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @drop.prevent="handleDrop"
    >
      <input
        type="file"
        ref="fileInput"
        accept=".md"
        @change="handleFileSelect"
        class="hidden"
      />
      
      <div class="space-y-4">
        <div class="flex justify-center">
          <div class="p-4 bg-indigo-50 rounded-full">
            <svg class="h-10 w-10 text-indigo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 014 4v2a4 4 0 01-4 4h-2m-2-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
          </div>
        </div>
        
        <p v-if="uploading" class="text-lg font-semibold text-indigo-600">
          Uploading...
        </p>
        <p v-else class="text-lg font-semibold text-gray-700">
          Drag & Drop a Markdown (.md) file here
        </p>
        
        <p class="text-sm text-gray-500">or</p>
        
        <button
          @click="fileInput.click()"
          :disabled="uploading"
          class="px-6 py-2 bg-indigo-600 text-white text-sm font-medium rounded-xl hover:bg-indigo-700 transition-colors shadow-md disabled:opacity-50"
        >
          Browse File
        </button>
        
        <p v-if="selectedFile" class="text-sm text-gray-600 pt-2">
          File Selected: <span class="font-medium">{{ selectedFile.name }}</span>
        </p>
      </div>
    </div>
    
    <div v-if="uploadError" class="mt-4 text-center text-red-600 font-medium">
      <p>{{ uploadError }}</p>
    </div>
    <div v-if="uploadSuccess" class="mt-4 text-center text-green-600 font-medium">
      <p>{{ uploadSuccess }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import api from '../services/api';

const emit = defineEmits(['upload-success', 'close']);

const fileInput = ref(null);
const selectedFile = ref(null);
const uploading = ref(false);
const uploadSuccess = ref('');
const uploadError = ref('');
const isDragging = ref(false);

const handleFileSelect = async (event) => {
  const file = event.target.files[0];
  if (!file) return;
  
  if (!file.name.endsWith('.md')) {
    uploadError.value = 'Please select a .md file';
    setTimeout(() => uploadError.value = '', 3000);
    return;
  }
  
  selectedFile.value = file;
  await uploadFile(file);
};

const handleDrop = async (event) => {
  isDragging.value = false;
  const file = event.dataTransfer.files[0];
  
  if (!file) return;
  
  if (!file.name.endsWith('.md')) {
    uploadError.value = 'Please drop a .md file';
    setTimeout(() => uploadError.value = '', 3000);
    return;
  }
  
  selectedFile.value = file;
  await uploadFile(file);
};

const uploadFile = async (file) => {
  uploading.value = true;
  uploadSuccess.value = '';
  uploadError.value = '';
  
  try {
    const response = await api.uploadMarkdown(file);
    uploadSuccess.value = response.data.message;
    selectedFile.value = null;
    fileInput.value.value = '';
    
    emit('upload-success');
    
    setTimeout(() => {
      uploadSuccess.value = '';
    }, 3000);
  } catch (error) {
    uploadError.value = error.response?.data?.error || 'Upload failed';
  } finally {
    uploading.value = false;
  }
};
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>