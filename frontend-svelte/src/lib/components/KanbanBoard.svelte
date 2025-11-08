<script>
  import { createEventDispatcher } from 'svelte'
  import TaskCard from './TaskCard.svelte'
  
  const dispatch = createEventDispatcher()
  
  export let tasks = []
  export let projectId
  
  const columns = [
    { id: 'waiting', title: 'Waiting', color: '#ef4444' },
    { id: 'in_progress', title: 'In Progress', color: '#3b82f6' },
    { id: 'testing', title: 'Testing', color: '#f59e0b' },
    { id: 'completed', title: 'Completed', color: '#10b981' }
  ]
  
  let dragOverColumn = null
  
  function getTasksByStatus(status) {
    return tasks
      .filter(task => task.status === status)
      .sort((a, b) => a.position - b.position)
  }
  
  function handleDragOver(event, columnId) {
    event.preventDefault()
    dragOverColumn = columnId
  }
  
  function handleDragLeave(event) {
    event.preventDefault()
    dragOverColumn = null
  }
  
  function handleDrop(event, columnId) {
    event.preventDefault()
    dragOverColumn = null
    
    const taskId = event.dataTransfer.getData('text/plain')
    if (taskId) {
      const task = tasks.find(t => t.id == taskId)
      if (task && task.status !== columnId) {
        dispatch('statusChange', {
          taskId: parseInt(taskId),
          newStatus: columnId,
          newPosition: getTasksByStatus(columnId).length
        })
      }
    }
  }
  
  function handleTaskUpdate(event) {
    dispatch('taskUpdate', event.detail)
  }
  
  function handleTaskDelete(event) {
    dispatch('taskDelete', { taskId: event.detail, projectId })
  }
</script>

<div class="kanban-board">
  {#each columns as column}
    <div 
      class="column {dragOverColumn === column.id ? 'drag-over' : ''}"
      on:dragover={(e) => handleDragOver(e, column.id)}
      on:dragleave={handleDragLeave}
      on:drop={(e) => handleDrop(e, column.id)}
    >
      <div class="column-header">
        <h3 class="column-title" style="color: {column.color}">
          {column.title}
        </h3>
        <span class="task-count">
          {getTasksByStatus(column.id).length}
        </span>
      </div>
      
      <div class="tasks-container">
        {#each getTasksByStatus(column.id) as task (task.id)}
          <TaskCard 
            {task} 
            editable={true}
            on:update={handleTaskUpdate}
            on:delete={handleTaskDelete}
          />
        {/each}
      </div>
    </div>
  {/each}
</div>

<style>
  .kanban-board {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1rem;
    margin-top: 2rem;
  }
  
  .column-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid var(--border-color);
  }
  
  .column-title {
    font-size: 1.125rem;
    font-weight: 600;
    margin: 0;
  }
  
  .task-count {
    background: var(--border-color);
    color: var(--text-muted);
    padding: 0.25rem 0.5rem;
    border-radius: 1rem;
    font-size: 0.75rem;
    font-weight: 600;
  }
  
  .tasks-container {
    min-height: 500px;
  }
  
  @media (max-width: 1024px) {
    .kanban-board {
      grid-template-columns: repeat(2, 1fr);
    }
  }
  
  @media (max-width: 768px) {
    .kanban-board {
      grid-template-columns: 1fr;
    }
  }
</style>