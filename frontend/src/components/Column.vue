<template>
  <div class="flex flex-col h-full">
    <div class="mb-4 flex items-center justify-between px-1">
      <div class="flex items-center">
        <div class="w-3 h-3 rounded-full mr-3" :class="dotColorClass"></div>
        <h3 class="font-bold text-gray-700 uppercase text-sm tracking-wide">
          {{ title }}
        </h3>
      </div>
      <span class="flex items-center justify-center w-7 h-7 rounded-full text-xs font-bold" 
            :class="badgeColorClass">
        {{ tasks.length }}
      </span>
    </div>
    
    <div class="flex-1 bg-gray-50 rounded-2xl p-4 min-h-[600px] border-2 border-gray-100">
      <draggable
        :list="tasks"
        group="tasks"
        item-key="id"
        class="space-y-3 min-h-[550px]"
        @change="handleDragChange"
        :animation="200"
      >
        <template #item="{ element }">
          <TaskCard
            :task="element"
            @edit="$emit('edit-task', element)"
            @delete="$emit('delete-task', $event)"
          />
        </template>
      </draggable>
      
      <div v-if="tasks.length === 0" class="flex flex-col items-center justify-center h-full text-gray-400">
        <svg class="w-16 h-16 mb-3 opacity-50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <p class="text-sm font-medium">No tasks</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import draggable from 'vuedraggable';
import TaskCard from './TaskCard.vue';

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  status: {
    type: String,
    required: true,
  },
  tasks: {
    type: Array,
    required: true,
  },
});

const emit = defineEmits(['task-moved', 'edit-task', 'delete-task']);

const dotColorClass = computed(() => {
  const colors = {
    waiting: 'bg-gray-400',
    in_progress: 'bg-blue-500',
    testing: 'bg-yellow-500',
    completed: 'bg-green-500',
  };
  return colors[props.status] || 'bg-gray-400';
});

const badgeColorClass = computed(() => {
  const colors = {
    waiting: 'bg-gray-200 text-gray-700',
    in_progress: 'bg-blue-100 text-blue-700',
    testing: 'bg-yellow-100 text-yellow-700',
    completed: 'bg-green-100 text-green-700',
  };
  return colors[props.status] || 'bg-gray-200 text-gray-700';
});

const handleDragChange = (event) => {
  if (event.added) {
    const task = event.added.element;
    emit('task-moved', {
      taskId: task.id,
      newStatus: props.status,
      newPosition: event.added.newIndex,
    });
  }
};
</script>