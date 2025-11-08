<script>
  import { createEventDispatcher } from 'svelte'
  
  const dispatch = createEventDispatcher()
  
  let fileInput
  let isDragging = false
  let uploadError = null
  
  function handleFileSelect(event) {
    const files = event.target.files
    if (files.length > 0) {
      validateAndUpload(files[0])
    }
  }
  
  function handleDragOver(event) {
    event.preventDefault()
    isDragging = true
  }
  
  function handleDragLeave(event) {
    event.preventDefault()
    isDragging = false
  }
  
  function handleDrop(event) {
    event.preventDefault()
    isDragging = false
    
    const files = event.dataTransfer.files
    if (files.length > 0) {
      validateAndUpload(files[0])
    }
  }
  
  async function validateAndUpload(file) {
    uploadError = null
    
    // Validate file extension
    if (!file.name.toLowerCase().endsWith('.md')) {
      uploadError = 'Please select a markdown file (.md)'
      return
    }
    
    // Validate file size (max 5MB)
    if (file.size > 5 * 1024 * 1024) {
      uploadError = 'File size must be less than 5MB'
      return
    }
    
    // Validate file is not empty
    if (file.size === 0) {
      uploadError = 'File is empty'
      return
    }
    
    console.log('Original file:', {
      name: file.name,
      size: file.size,
      type: file.type,
      lastModified: file.lastModified
    })
    
    try {
      // Read the file content to ensure it's valid
      const fileContent = await readFileContent(file)
      
      // Validate basic markdown structure
      if (!fileContent.includes('# ')) {
        uploadError = 'Markdown file must contain a project name starting with #'
        return
      }
      
      // Create a new File object with the correct type and actual content
      const correctedFile = new File([fileContent], file.name, { 
        type: 'text/markdown',
        lastModified: file.lastModified
      })
      
      console.log('Corrected file ready for upload:', {
        name: correctedFile.name,
        size: correctedFile.size,
        type: correctedFile.type
      })
      
      // Dispatch the file directly, not wrapped in a CustomEvent detail
      dispatch('fileSelected', correctedFile)
      
    } catch (error) {
      uploadError = 'Failed to read file content'
      console.error('File reading error:', error)
    }
  }
  
  function readFileContent(file) {
    return new Promise((resolve, reject) => {
      const reader = new FileReader()
      reader.onload = (e) => resolve(e.target.result)
      reader.onerror = (e) => reject(e)
      reader.readAsText(file)
    })
  }
  
  function triggerFileInput() {
    fileInput.click()
  }
</script>

<div 
  class="file-upload-area {isDragging ? 'drag-over' : ''}"
  on:dragover={handleDragOver}
  on:dragleave={handleDragLeave}
  on:drop={handleDrop}
>
  <input
    bind:this={fileInput}
    type="file"
    accept=".md"
    on:change={handleFileSelect}
    style="display: none;"
  />
  
  <div class="upload-content">
    <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor">
      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
      <polyline points="14,2 14,8 20,8"/>
      <line x1="16" y1="13" x2="8" y2="13"/>
      <line x1="16" y1="17" x2="8" y2="17"/>
      <polyline points="10,9 9,9 8,9"/>
    </svg>
    <h3>Upload Markdown File</h3>
    <p>Drag and drop your .md file here, or click to browse</p>
    <button class="btn btn-primary" on:click={triggerFileInput}>
      Choose File
    </button>
    
    {#if uploadError}
      <div class="upload-error">
        {uploadError}
      </div>
    {/if}
  </div>
</div>

<style>
  .upload-error {
    color: var(--error-color);
    font-size: 0.875rem;
    margin-top: 0.5rem;
    padding: 0.5rem;
    background: #fef2f2;
    border-radius: 0.25rem;
    border: 1px solid var(--error-color);
  }
</style>