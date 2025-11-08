<script>
  import { onMount } from 'svelte'
  import { projects } from './lib/stores/projects.js'
  import { tasks } from './lib/stores/tasks.js'
  import ProjectList from './lib/components/ProjectList.svelte'
  import KanbanBoard from './lib/components/KanbanBoard.svelte'
  import FileUpload from './lib/components/FileUpload.svelte'
  import { testFileUpload } from './lib/utils/testUpload.js'
  import './styles/global.css'
  
  let currentView = 'projects' // 'projects' or 'kanban'
  let selectedProject = null
  let uploadSuccess = null
  
  onMount(() => {
    projects.fetchAll()
  })
  
  async function handleFileUpload(file) {
    console.log('App: Received file for upload:', file)
    
    if (!file || !(file instanceof File)) {
      console.error('App: Invalid file received:', file)
      return
    }
    
    try {
      const result = await projects.uploadProject(file)
      uploadSuccess = `Successfully uploaded "${result.project.name}" with ${result.tasks_created} tasks`
      
      // Clear success message after 3 seconds
      setTimeout(() => {
        uploadSuccess = null
      }, 3000)
      
      // Refresh projects list
      projects.fetchAll()
    } catch (error) {
      console.error('App: Upload failed:', error)
    }
  }
  
  function handleFileSelected(event) {
    const file = event.detail
    console.log('App: File selected event:', file)
    if (file instanceof File) {
      handleFileUpload(file)
    } else {
      console.error('App: Invalid file in event detail:', file)
    }
  }
  
  function handleProjectSelect(project) {
    selectedProject = project
    currentView = 'kanban'
    tasks.fetchByProject(project.id)
  }
  
  function handleBackToProjects() {
    currentView = 'projects'
    selectedProject = null
  }
  
  async function handleProjectDelete(projectId) {
    try {
      await projects.deleteProject(projectId)
      if (selectedProject && selectedProject.id === projectId) {
        handleBackToProjects()
      }
    } catch (error) {
      // Error handled by store
    }
  }
  
  async function handleTaskUpdate(updateData) {
    try {
      await tasks.updateTask(updateData.id, updateData)
    } catch (error) {
      // Error handled by store
    }
  }
  
  async function handleTaskDelete({ taskId, projectId }) {
    try {
      await tasks.deleteTask(taskId, projectId)
    } catch (error) {
      // Error handled by store
    }
  }
  
  async function handleStatusChange({ taskId, newStatus, newPosition }) {
    try {
      await tasks.updateTask(taskId, {
        status: newStatus,
        position: newPosition
      })
    } catch (error) {
      // Error handled by store
    }
  }
  
  async function handleTestUpload() {
    try {
      const testFile = await testFileUpload()
      await handleFileUpload(testFile)
    } catch (error) {
      console.error('Test upload failed:', error)
    }
  }
</script>

<svelte:head>
  <title>Turzone Kanban</title>
</svelte:head>

<div class="app">
  <header class="app-header">
    <div class="container">
      <h1 class="app-title">Turzone Kanban</h1>
      <p class="app-subtitle">Visualize your workflow, one task at a time</p>
    </div>
  </header>
  
  <main class="app-main">
    <div class="container">
      {#if $projects.error}
        <div class="alert alert-error">
          {$projects.error}
          <button class="alert-close" on:click={() => projects.clearError()}>×</button>
        </div>
      {/if}
      
      {#if $tasks.error}
        <div class="alert alert-error">
          {$tasks.error}
          <button class="alert-close" on:click={() => tasks.clearError()}>×</button>
        </div>
      {/if}
      
      {#if uploadSuccess}
        <div class="alert alert-success">
          {uploadSuccess}
        </div>
      {/if}
      
      {#if currentView === 'projects'}
        <!-- Debug Section -->
        <section class="debug-section">
          <h3>Debug Tools</h3>
          <button class="btn btn-secondary" on:click={handleTestUpload}>
            Test Upload with Sample File
          </button>
          <div class="debug-info">
            <p>Backend expects: Content-Type: text/markdown</p>
            <p>File field name: "file"</p>
            <p>Markdown format: # Project Name followed by ## Task titles</p>
          </div>
        </section>
        
        <section class="upload-section">
          <h2>Start a New Project</h2>
          <FileUpload on:fileSelected={handleFileSelected} />
        </section>
        
        <section class="projects-section">
          <h2>Your Projects</h2>
          {#if $projects.loading}
            <div class="loading">Loading projects...</div>
          {:else}
            <ProjectList 
              projects={$projects}
              on:select={handleProjectSelect}
              on:delete={handleProjectDelete}
            />
          {/if}
        </section>
      
      {:else if currentView === 'kanban' && selectedProject}
        <section class="kanban-section">
          <div class="kanban-header">
            <button class="btn btn-secondary" on:click={handleBackToProjects}>
              ← Back to Projects
            </button>
            <h2>{selectedProject.name}</h2>
            <div class="project-actions">
              <span class="last-updated">
                Updated: {new Date(selectedProject.updated_at).toLocaleString()}
              </span>
            </div>
          </div>
          
          {#if $tasks.loading}
            <div class="loading">Loading tasks...</div>
          {:else}
            <KanbanBoard 
              tasks={$tasks[selectedProject.id] || []}
              projectId={selectedProject.id}
              on:statusChange={handleStatusChange}
              on:taskUpdate={handleTaskUpdate}
              on:taskDelete={handleTaskDelete}
            />
          {/if}
        </section>
      {/if}
    </div>
  </main>
</div>

<style>
  :root {
    --primary-color: #3b82f6;
    --secondary-color: #64748b;
    --success-color: #10b981;
    --warning-color: #f59e0b;
    --error-color: #ef4444;
    --bg-color: #f8fafc;
    --card-bg: #ffffff;
    --border-color: #e2e8f0;
    --text-color: #1e293b;
    --text-muted: #64748b;
  }

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
    background-color: var(--bg-color);
    color: var(--text-color);
    line-height: 1.6;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 1rem;
  }

  .app-header {
    background: var(--card-bg);
    border-bottom: 1px solid var(--border-color);
    padding: 2rem 0;
    margin-bottom: 2rem;
  }
  
  .app-title {
    font-size: 2.5rem;
    font-weight: 700;
    color: var(--text-color);
    margin: 0 0 0.5rem 0;
  }
  
  .app-subtitle {
    font-size: 1.125rem;
    color: var(--text-muted);
    margin: 0;
  }
  
  .upload-section,
  .projects-section,
  .kanban-section {
    margin-bottom: 3rem;
  }
  
  .debug-section {
    margin: 2rem 0;
    padding: 1.5rem;
    background: #f8fafc;
    border-radius: 0.5rem;
    border: 1px solid var(--border-color);
  }
  
  .debug-section h3 {
    margin: 0 0 1rem 0;
    font-size: 1.25rem;
    color: var(--text-color);
  }
  
  .debug-info {
    margin-top: 1rem;
    font-size: 0.875rem;
    color: var(--text-muted);
  }
  
  .debug-info p {
    margin: 0.25rem 0;
  }
  
  h2 {
    font-size: 1.5rem;
    font-weight: 600;
    color: var(--text-color);
    margin-bottom: 1rem;
  }
  
  .alert {
    padding: 1rem;
    border-radius: 0.375rem;
    margin-bottom: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .alert-error {
    background-color: #fef2f2;
    border: 1px solid var(--error-color);
    color: var(--error-color);
  }
  
  .alert-success {
    background-color: #f0fdf4;
    border: 1px solid var(--success-color);
    color: var(--success-color);
  }
  
  .alert-close {
    background: none;
    border: none;
    font-size: 1.25rem;
    cursor: pointer;
    color: inherit;
  }
  
  .loading {
    text-align: center;
    padding: 2rem;
    color: var(--text-muted);
  }
  
  .btn {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 0.375rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
  }
  
  .btn-primary {
    background-color: var(--primary-color);
    color: white;
  }
  
  .btn-primary:hover {
    background-color: #2563eb;
  }
  
  .btn-danger {
    background-color: var(--error-color);
    color: white;
  }
  
  .btn-danger:hover {
    background-color: #dc2626;
  }
  
  .btn-secondary {
    background-color: var(--secondary-color);
    color: white;
  }
  
  .btn-secondary:hover {
    background-color: #475569;
  }
  
  .kanban-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    gap: 1rem;
  }
  
  .kanban-header h2 {
    margin: 0;
    flex: 1;
    text-align: center;
  }
  
  .last-updated {
    color: var(--text-muted);
    font-size: 0.875rem;
  }
  
  @media (max-width: 768px) {
    .app-title {
      font-size: 2rem;
    }
    
    .kanban-header {
      flex-direction: column;
      gap: 1rem;
    }
    
    .kanban-header h2 {
      text-align: center;
    }
    
    .debug-section {
      padding: 1rem;
    }
  }
</style>