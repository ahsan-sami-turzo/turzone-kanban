<template>
  <div class="bg-gray-50 rounded-lg p-4 min-h-[500px]">
    <div class="flex items-center justify-between mb-4">
      <h3 class="font-bold text-gray-700 uppercase text-sm">
        {{ title }}
      </h3>
      <span class="bg-gray-200 text-gray-700 rounded-full px-2 py-1 text-xs">
        {{ tasks.length }}
      </span>
    </div>
    
    <draggable
      :list="tasks"
      group="tasks"
      item-key="id"
      class="space-y-3 min-h-[400px]"
      @change="handleDragChange"
    >
      <template #item="{ element }">
        <TaskCard
          :task="element"
          @edit="$emit('edit-task', element)"
          @delete="$emit('delete-task', $event)"
        />
      </template>
    </draggable>
  </div>
</template>

<script setup>
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