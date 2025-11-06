<template>
  <div
    @click="$emit('select', project)"
    class="bg-white rounded-lg shadow-sm hover:shadow-md transition cursor-pointer border-l-4 p-6"
    :class="borderColorClass"
  >
    <div class="flex items-start justify-between">
      <div class="flex-1">
        <h3 class="text-xl font-bold text-gray-800 mb-2">{{ project.name }}</h3>
        <p class="text-sm text-gray-500 mb-4">
          Created {{ formatDate(project.created_at) }}
        </p>
        
        <div class="flex items-center space-x-4 text-sm">
          <div class="flex items-center">
            <span class="font-semibold text-gray-700 mr-1">{{ stats.total }}</span>
            <span class="text-gray-500">Total Tasks</span>
          </div>
        </div>
        
        <div class="mt-4 flex items-center space-x-3">
          <div class="flex items-center text-sm">
            <div class="w-3 h-3 rounded-full bg-gray-400 mr-2"></div>
            <span class="text-gray-600">{{ stats.waiting }} waiting</span>
          </div>
          <div class="flex items-center text-sm">
            <div class="w-3 h-3 rounded-full bg-blue-500 mr-2"></div>
            <span class="text-gray-600">{{ stats.in_progress }} in progress</span>
          </div>
          <div class="flex items-center text-sm">
            <div class="w-3 h-3 rounded-full bg-yellow-500 mr-2"></div>
            <span class="text-gray-600">{{ stats.testing }} testing</span>
          </div>
          <div class="flex items-center text-sm">
            <div class="w-3 h-3 rounded-full bg-green-500 mr-2"></div>
            <span class="text-gray-600">{{ stats.completed }} completed</span>
          </div>
        </div>
      </div>
      
      <button
        @click.stop="$emit('delete', project.id)"
        class="text-gray-400 hover:text-red-500 transition p-2"
        title="Delete project"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  project: {
    type: Object,
    required: true,
  },
  stats: {
    type: Object,
    required: true,
  },
});

defineEmits(['select', 'delete']);

const borderColorClass = computed(() => {
  const { waiting, in_progress, testing, completed, total } = props.stats;
  
  // All completed - green
  if (completed === total && total > 0) {
    return 'border-green-500';
  }
  
  // Some in progress or testing - blue
  if (in_progress > 0 || testing > 0) {
    return 'border-blue-500';
  }
  
  // All waiting - red
  if (waiting === total && total > 0) {
    return 'border-red-500';
  }
  
  // Default - gray
  return 'border-gray-300';
});

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric'
  });
};
</script>