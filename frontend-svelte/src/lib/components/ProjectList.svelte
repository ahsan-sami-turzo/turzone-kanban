<script>
  import { createEventDispatcher } from 'svelte'
  import StatusIndicator from './StatusIndicator.svelte'
  
  const dispatch = createEventDispatcher()
  
  export let projects = []
  
  function getProjectStatus(project) {
    // This would need to be calculated based on tasks
    // For now, return a mock status
    return 'in_progress'
  }
  
  function handleProjectSelect(project) {
    dispatch('select', project)
  }
  
  function handleProjectDelete(project, event) {
    event.stopPropagation()
    if (confirm(`Are you sure you want to delete "${project.name}"?`)) {
      dispatch('delete', project.id)
    }
  }
</script>

<div class="project-list">
  {#each projects as project (project.id)}
    <div 
      class="project-card card fade-in"
      on:click={() => handleProjectSelect(project)}
    >
      <div class="project-info">
        <h3 class="project-name">{project.name}</h3>
        <StatusIndicator status={getProjectStatus(project)} />
      </div>
      
      <div class="project-meta">
        <span class="project-date">
          Created: {new Date(project.created_at).toLocaleDateString()}
        </span>
      </div>
      
      <button 
        class="btn btn-danger btn-sm"
        on:click={(e) => handleProjectDelete(project, e)}
      >
        Delete
      </button>
    </div>
  {:else}
    <div class="empty-state">
      <h3>No projects yet</h3>
      <p>Upload a markdown file to get started!</p>
    </div>
  {/each}
</div>

<style>
  .project-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1rem;
    margin-top: 2rem;
  }
  
  .project-card {
    padding: 1.5rem;
    cursor: pointer;
    transition: transform 0.2s, box-shadow 0.2s;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  
  .project-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  }
  
  .project-info {
    display: flex;
    justify-content: space-between;
    align-items: start;
    gap: 1rem;
  }
  
  .project-name {
    margin: 0;
    font-size: 1.25rem;
    font-weight: 600;
    color: var(--text-color);
    flex: 1;
  }
  
  .project-meta {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .project-date {
    color: var(--text-muted);
    font-size: 0.875rem;
  }
  
  .btn-sm {
    padding: 0.375rem 0.75rem;
    font-size: 0.75rem;
    align-self: flex-end;
  }
  
  .empty-state {
    grid-column: 1 / -1;
    text-align: center;
    padding: 3rem 2rem;
    color: var(--text-muted);
  }
  
  .empty-state h3 {
    margin: 0 0 0.5rem 0;
    color: var(--text-color);
  }
</style>