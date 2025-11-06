<template>
  <div
    @click="$emit('select', project)"
    class="group relative bg-white rounded-2xl shadow-sm hover:shadow-xl transition-all duration-300 cursor-pointer border-l-4 overflow-hidden transform hover:-translate-y-1"
    :class="borderColorClass"
  >
    <!-- Colored Header Bar -->
    <div class="absolute top-0 right-0 w-32 h-32 -mr-8 -mt-8 rounded-full opacity-10 transition-all duration-300 group-hover:scale-110"
         :class="bgColorClass">
    </div>
    
    <div class="relative p-6">
      <div class="flex items-start justify-between mb-4">
        <div class="flex-1">
          <h3 class="text-xl font-bold text-gray-800 mb-1 group-hover:text-blue-600 transition-colors">
            {{ project.name }}
          </h3>
          <p class="text-xs text-gray-500 flex items-center">
            <svg class="w-3.5 h-3.5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
            {{ formatDate(project.created_at) }}
          </p>
        </div>
        
        <button
          @click.stop="$emit('delete', project.id)"
          class="opacity-0 group-hover:opacity-100 transition-opacity p-2 hover:bg-red-50 rounded-lg text-gray-400 hover:text-red-500"
          title="Delete project"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </div>
      
      <!-- Total Tasks -->
      <div class="mb-4 inline-flex items-center px-4 py-2 bg-gray-50 rounded-lg">
        <svg class="w-5 h-5 text-gray-400 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <span class="text-2xl font-bold text-gray-800 mr-2">{{ stats.total }}</span>
        <span class="text-sm text-gray-500">{{ stats.total === 1 ? 'task' : 'tasks' }}</span>
      </div>
      
      <!-- Status Grid -->
      <div class="grid grid-cols-2 gap-3">
        <div class="flex items-center px-3 py-2 bg-gray-50 rounded-lg">
          <div class="w-2 h-2 rounded-full bg-gray-400 mr-2"></div>
          <span class="text-sm text-gray-600">
            <span class="font-semibold text-gray-800">{{ stats.waiting }}</span> waiting
          </span>
        </div>
        
        <div class="flex items-center px-3 py-2 bg-blue-50 rounded-lg">
          <div class="w-2 h-2 rounded-full bg-blue-500 mr-2"></div>
          <span class="text-sm text-blue-700">
            <span class="font-semibold">{{ stats.in_progress }}</span> active
          </span>
        </div>
        
        <div class="flex items-center px-3 py-2 bg-yellow-50 rounded-lg">
          <div class="w-2 h-2 rounded-full bg-yellow-500 mr-2"></div>
          <span class="text-sm text-yellow-700">
            <span class="font-semibold">{{ stats.testing }}</span> testing
          </span>
        </div>
        
        <div class="flex items-center px-3 py-2 bg-green-50 rounded-lg">
          <div class="w-2 h-2 rounded-full bg-green-500 mr-2"></div>
          <span class="text-sm text-green-700">
            <span class="font-semibold">{{ stats.completed }}</span> done
          </span>
        </div>
      </div>
      
      <!-- Progress Bar -->
      <div class="mt-4 pt-4 border-t border-gray-100">
        <div class="flex items-center justify-between text-xs text-gray-500 mb-2">
          <span>Progress</span>
          <span class="font-semibold">{{ progressPercentage }}%</span>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-2 overflow-hidden">
          <div 
            class="h-2 rounded-full transition-all duration-500 ease-out"
            :class="progressColorClass"
            :style="{ width: progressPercentage + '%' }"
          ></div>
        </div>
      </div>
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

const progressPercentage = computed(() => {
  if (props.stats.total === 0) return 0;
  return Math.round((props.stats.completed / props.stats.total) * 100);
});

const borderColorClass = computed(() => {
  const { waiting, in_progress, testing, completed, total } = props.stats;
  
  if (completed === total && total > 0) return 'border-green-500';
  if (in_progress > 0 || testing > 0) return 'border-blue-500';
  if (waiting === total && total > 0) return 'border-red-500';
  return 'border-gray-300';
});

const bgColorClass = computed(() => {
  const { waiting, in_progress, testing, completed, total } = props.stats;
  
  if (completed === total && total > 0) return 'bg-green-500';
  if (in_progress > 0 || testing > 0) return 'bg-blue-500';
  if (waiting === total && total > 0) return 'bg-red-500';
  return 'bg-gray-300';
});

const progressColorClass = computed(() => {
  const percentage = progressPercentage.value;
  if (percentage === 100) return 'bg-green-500';
  if (percentage >= 50) return 'bg-blue-500';
  if (percentage > 0) return 'bg-yellow-500';
  return 'bg-gray-400';
});

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric'
  });
};
</script>