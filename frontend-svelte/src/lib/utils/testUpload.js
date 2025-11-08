export async function testFileUpload() {
  // Create a test markdown file programmatically
  const markdownContent = `# Test Project

## First Task
This is the first task description
- Subtask 1
- Subtask 2

## Second Task  
This is the second task description
- Another subtask
`

  const testFile = new File([markdownContent], 'test-project.md', {
    type: 'text/markdown',
    lastModified: new Date()
  })

  console.log('Test file created:', {
    name: testFile.name,
    type: testFile.type,
    size: testFile.size
  })

  return testFile
}