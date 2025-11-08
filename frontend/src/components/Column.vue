<template>
  <div class="flex flex-col h-full w-full sm:w-[300px] flex-shrink-0">
    <div class="mb-4 flex items-center justify-between px-1">
      <div class="flex items-center">
        <div class="w-2.5 h-2.5 rounded-full mr-2" :class="dotColorClass"></div>
        <h3 class="font-bold text-gray-700 uppercase text-xs tracking-wider">
          {{ title }}
        </h3>
      </div>
      <span class="text-xs font-semibold text-gray-500 py-1 px-2 bg-gray-100 rounded-full">
        {{ tasks.length }}
      </span>
    </div>
    
    <div class="flex-1 bg-gray-50/70 rounded-xl p-3 min-h-[500px] border-2 border-gray-200 transition-all duration-300">
      <draggable
        :list="tasks"
        group="tasks"
        item-key="id"
        class="space-y-4 min-h-full p-1 transition-all duration-200"
        @change="handleDragChange"
        :animation="200"
        
        :class="{'bg-indigo-100/50 border-indigo-300': isDraggingOn}"
        @dragover.prevent="isDraggingOn = true"
        @dragleave.prevent="isDraggingOn = false"
        @drop="isDraggingOn = false"
      >
        <template #item="{ element }">
          <TaskCard
            :task="element"
            @edit="$emit('edit-task', element)"
            @delete="$emit('delete-task', $event)"
          />
        </template>

        <template #footer>
          <div v-if="tasks.length === 0" class="flex flex-col items-center justify-center text-center py-12 text-gray-400 h-full">
            <svg class="w-8 h-8 mb-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
            </svg>
            <p class="text-sm font-medium">No tasks here</p>
          </div>
        </template>
      </draggable>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
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

// State to track if an element is being dragged over this column
const isDraggingOn = ref(false); 

const dotColorClass = computed(() => {
  const colors = {
    waiting: 'bg-gray-400',
    in_progress: 'bg-indigo-500',
    testing: 'bg-yellow-500',
    completed: 'bg-green-500',
    default: 'bg-gray-300',
  };
  return colors[props.status] || colors.default;
});

// Removed badgeColorClass as it's replaced by a simpler Tailwind style
const handleDragChange = (event) => {
  if (event.added || event.moved) {
    const task = event.added ? event.added.element : event.moved.element;
    const newIndex = event.added ? event.added.newIndex : event.moved.newIndex;
    
    emit('task-moved', {
      taskId: task.id,
      newStatus: props.status,
      newPosition: newIndex,
    });
  }
};
</script>