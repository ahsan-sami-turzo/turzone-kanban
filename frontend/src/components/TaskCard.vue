<template>
  <div class="bg-white p-4 rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition cursor-move">
    <h3 class="font-semibold text-gray-800 mb-2">{{ task.title }}</h3>
    
    <p v-if="task.description" class="text-sm text-gray-600 mb-3 whitespace-pre-line">
      {{ task.description }}
    </p>
    
    <div class="flex justify-between items-center">
      <span class="text-xs text-gray-400">
        ID: {{ task.id }}
      </span>
      
      <div class="space-x-2">
        <button
          @click="$emit('edit', task)"
          class="text-blue-500 hover:text-blue-700 text-sm"
        >
          Edit
        </button>
        <button
          @click="$emit('delete', task.id)"
          class="text-red-500 hover:text-red-700 text-sm"
        >
          Delete
        </button>
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