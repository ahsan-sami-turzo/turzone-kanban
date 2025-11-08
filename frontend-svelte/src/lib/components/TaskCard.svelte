<script>
  import { createEventDispatcher } from 'svelte'
  import StatusIndicator from './StatusIndicator.svelte'
  
  const dispatch = createEventDispatcher()
  
  export let task
  export let editable = false
  
  let isEditing = false
  let editTitle = ''
  let editDescription = ''
  
  function handleEdit() {
    isEditing = true
    editTitle = task.title
    editDescription = task.description
  }
  
  function handleSave() {
    dispatch('update', {
      id: task.id,
      title: editTitle,
      description: editDescription
    })
    isEditing = false
  }
  
  function handleCancel() {
    isEditing = false
    editTitle = task.title
    editDescription = task.description
  }
  
  function handleDelete() {
    if (confirm('Are you sure you want to delete this task?')) {
      dispatch('delete', task.id)
    }
  }
  
  function handleDragStart(event) {
    event.dataTransfer.setData('text/plain', task.id.toString())
    dispatch('dragstart', task)
  }
</script>

<div 
  class="task-card card status-{task.status} fade-in"
  draggable="true"
  on:dragstart={handleDragStart}
>
  {#if isEditing}
    <div class="edit-form">
      <input 
        type="text" 
        bind:value={editTitle}
        class="form-input"
        placeholder="Task title"
      />
      <textarea
        bind:value={editDescription}
        class="form-textarea"
        placeholder="Task description"
        rows="3"
      ></textarea>
      <div class="edit-actions">
        <button class="btn btn-primary" on:click={handleSave}>Save</button>
        <button class="btn btn-secondary" on:click={handleCancel}>Cancel</button>
      </div>
    </div>
  {:else}
    <div class="task-header">
      <h4 class="task-title">{task.title}</h4>
      <StatusIndicator {status} />
    </div>
    
    <div class="task-description">
      {#each task.description.split('\n') as line}
        {#if line.trim()}
          <p>{line}</p>
        {/if}
      {/each}
    </div>
    
    {#if editable}
      <div class="task-actions">
        <button class="btn btn-secondary" on:click={handleEdit}>Edit</button>
        <button class="btn btn-danger" on:click={handleDelete}>Delete</button>
      </div>
    {/if}
  {/if}
</div>

<style>
  .task-card {
    padding: 1rem;
    margin-bottom: 1rem;
    cursor: grab;
    transition: transform 0.2s, box-shadow 0.2s;
  }
  
  .task-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  }
  
  .task-card:active {
    cursor: grabbing;
  }
  
  .task-header {
    display: flex;
    justify-content: between;
    align-items: start;
    gap: 1rem;
    margin-bottom: 0.5rem;
  }
  
  .task-title {
    flex: 1;
    margin: 0;
    font-size: 1rem;
    font-weight: 600;
    color: var(--text-color);
  }
  
  .task-description {
    color: var(--text-muted);
    font-size: 0.875rem;
    line-height: 1.5;
  }
  
  .task-description p {
    margin: 0.25rem 0;
  }
  
  .task-actions {
    display: flex;
    gap: 0.5rem;
    margin-top: 1rem;
  }
  
  .edit-form {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }
  
  .form-input,
  .form-textarea {
    padding: 0.5rem;
    border: 1px solid var(--border-color);
    border-radius: 0.375rem;
    font-family: inherit;
    font-size: 0.875rem;
  }
  
  .form-input:focus,
  .form-textarea:focus {
    outline: none;
    border-color: var(--primary-color);
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }
  
  .edit-actions {
    display: flex;
    gap: 0.5rem;
  }
  
  .btn {
    font-size: 0.75rem;
    padding: 0.375rem 0.75rem;
  }
</style>