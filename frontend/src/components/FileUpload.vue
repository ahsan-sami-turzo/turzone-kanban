<template>
  <div class="mb-6">
    <div 
      class="border-2 border-dashed rounded-lg p-8 text-center transition"
      :class="isDragging ? 'border-blue-500 bg-blue-50' : 'border-gray-300 hover:border-blue-400'"
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
      
      <div class="space-y-3">
        <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
          <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
        </svg>
        
        <div>
          <button
            @click="$refs.fileInput.click()"
            class="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600 transition font-medium"
          >
            Choose Markdown File
          </button>
          <p class="text-gray-500 text-sm mt-2">or drag and drop your .md file here</p>
        </div>
        
        <p v-if="selectedFile" class="text-blue-600 font-medium">
          📄 {{ selectedFile.name }}
        </p>
      </div>
    </div>
    
    <div v-if="uploading" class="mt-4 flex items-center justify-center text-blue-600">
      <svg class="animate-spin h-5 w-5 mr-2" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" fill="none"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      Uploading...
    </div>
    
    <div v-if="uploadSuccess" class="mt-4 bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded flex items-center">
      <span class="mr-2">✓</span>
      {{ uploadSuccess }}
    </div>
    
    <div v-if="uploadError" class="mt-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded flex items-center">
      <span class="mr-2">✗</span>
      {{ uploadError }}
    </div>
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