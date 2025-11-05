<template>
  <div class="mb-6">
    <div class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center hover:border-blue-500 transition">
      <input
        type="file"
        ref="fileInput"
        accept=".md"
        @change="handleFileSelect"
        class="hidden"
      />
      
      <button
        @click="$refs.fileInput.click()"
        class="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600 transition"
      >
        Upload Markdown File
      </button>
      
      <p class="text-gray-500 text-sm mt-2">
        {{ selectedFile ? selectedFile.name : 'Select a .md file to upload' }}
      </p>
    </div>
    
    <div v-if="uploading" class="mt-4 text-center text-blue-600">
      Uploading...
    </div>
    
    <div v-if="uploadSuccess" class="mt-4 bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
      {{ uploadSuccess }}
    </div>
    
    <div v-if="uploadError" class="mt-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
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

const handleFileSelect = async (event) => {
  const file = event.target.files[0];
  if (!file) return;
  
  selectedFile.value = file;
  
  // Auto upload
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
    
    // Notify parent component
    emit('upload-success');
    
    // Clear success message after 3 seconds
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