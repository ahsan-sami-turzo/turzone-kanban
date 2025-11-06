<template>
  <div class="group bg-white rounded-xl shadow-sm hover:shadow-lg transition-all duration-300 cursor-move border border-gray-100 hover:border-blue-200 transform hover:scale-[1.02]">
    <div class="p-4">
      <div class="flex items-start justify-between mb-3">
        <h3 class="font-semibold text-gray-800 flex-1 pr-2 leading-snug">
          {{ task.title }}
        </h3>
        
        <div class="flex-shrink-0 flex space-x-1 opacity-0 group-hover:opacity-100 transition-opacity">
          <button
            @click.stop="$emit('edit', task)"
            class="p-1.5 hover:bg-blue-50 rounded-lg text-gray-400 hover:text-blue-600 transition-colors"
            title="Edit task"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
          </button>
          <button
            @click.stop="$emit('delete', task.id)"
            class="p-1.5 hover:bg-red-50 rounded-lg text-gray-400 hover:text-red-600 transition-colors"
            title="Delete task"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
          </button>
        </div>
      </div>
      
      <p v-if="task.description" class="text-sm text-gray-600 mb-3 line-clamp-3 leading-relaxed">
        {{ task.description }}
      </p>
      
      <div class="flex items-center justify-between pt-3 border-t border-gray-100">
        <span class="inline-flex items-center px-2 py-1 bg-gray-100 rounded-md text-xs font-medium text-gray-600">
          #{{ task.id }}
        </span>
        
        <div class="flex items-center text-xs text-gray-400">
          <svg class="w-3.5 h-3.5 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
          </svg>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  task: {
    type: Object,
    required: true,
  },
});

defineEmits(['edit', 'delete']);
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>