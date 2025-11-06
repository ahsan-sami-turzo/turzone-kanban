<template>
  <div class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition cursor-move group">
    <div class="flex items-start justify-between mb-2">
      <h3 class="font-semibold text-gray-800 flex-1 pr-2">{{ task.title }}</h3>
      <div class="flex-shrink-0 opacity-0 group-hover:opacity-100 transition-opacity">
        <button
          @click.stop="$emit('edit', task)"
          class="text-blue-500 hover:text-blue-700 p-1"
          title="Edit task"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
          </svg>
        </button>
        <button
          @click.stop="$emit('delete', task.id)"
          class="text-red-500 hover:text-red-700 p-1 ml-1"
          title="Delete task"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </div>
    </div>
    
    <p v-if="task.description" class="text-sm text-gray-600 whitespace-pre-line line-clamp-3">
      {{ task.description }}
    </p>
    
    <div class="mt-3 text-xs text-gray-400">
      #{{ task.id }}
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