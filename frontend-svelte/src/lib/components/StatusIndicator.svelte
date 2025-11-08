<script>
  export let status;
  
  const statusConfig = {
    waiting: { color: '#ef4444', label: 'Waiting' },
    in_progress: { color: '#3b82f6', label: 'In Progress' },
    testing: { color: '#f59e0b', label: 'Testing' },
    completed: { color: '#10b981', label: 'Completed' }
  };
  
  $: config = statusConfig[status] || statusConfig.waiting;
</script>

<div class="status-indicator">
  <span 
    class="status-dot" 
    style="background-color: {config.color}"
    title="{config.label}"
  ></span>
  <span class="status-label">{config.label}</span>
</div>

<style>
  .status-indicator {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    font-size: 0.875rem;
  }
  
  .status-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    display: inline-block;
  }
  
  .status-label {
    color: var(--text-muted);
    font-weight: 500;
  }
</style>