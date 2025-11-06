<template>
  <div class="relative">
    <div 
      class="relative border-2 border-dashed rounded-2xl p-10 text-center transition-all duration-300 ease-in-out"
      :class="isDragging 
        ? 'border-blue-500 bg-blue-50 scale-[1.02]' 
        : 'border-gray-300 bg-white hover:border-blue-400 hover:bg-gray-50'"
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
        <!-- Icon -->
        <div class="flex justify-center">
          <div class="p-4 bg-blue-50 rounded-full">
            <svg class="h-10 w-10 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
            </svg>
          </div>
        </div>
        
        <!-- Text -->
        <div>
          <button
            @click="$refs.fileInput.click()"
            class="inline-flex items-center px-6 py-3 bg-gradient-to-r from-blue-500 to-blue-600 text-white font-medium rounded-xl hover:from-blue-600 hover:to-blue-700 transition-all duration-200 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5"
          >
            <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Upload Markdown
          </button>
          <p class="text-sm text-gray-500 mt-3">
            or drag and drop your <span class="font-semibold text-gray-700">.md</span> file here
          </p>
        </div>
        
        <p v-if="selectedFile" class="text-blue-600 font-medium flex items-center justify-center">
          <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          {{ selectedFile.name }}
        </p>
      </div>
    </div>
    
    <!-- Loading State -->
    <div v-if="uploading" class="mt-4 flex items-center justify-center p-4 bg-blue-50 rounded-xl border border-blue-200">
      <svg class="animate-spin h-5 w-5 text-blue-500 mr-3" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span class="text-blue-700 font-medium">Uploading...</span>
    </div>
    
    <!-- Success Message -->
    <transition name="fade">
      <div v-if="uploadSuccess" class="mt-4 p-4 bg-green-50 border border-green-200 text-green-700 rounded-xl flex items-start">
        <svg class="w-5 h-5 mr-3 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
        </svg>
        <span class="font-medium">{{ uploadSuccess }}</span>
      </div>
    </transition>
    
    <!-- Error Message -->
    <transition name="fade">
      <div v-if="uploadError" class="mt-4 p-4 bg-red-50 border border-red-200 text-red-700 rounded-xl flex items-start">
        <svg class="w-5 h-5 mr-3 mt-0.5 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
        </svg>
        <span class="font-medium">{{ uploadError }}</span>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import api from '../services/api';

const emit = defineEmits(['upload-success']);

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