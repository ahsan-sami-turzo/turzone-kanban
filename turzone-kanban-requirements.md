# Workflow Visualization Tool - Requirements & Development Plan

## Project Overview
A local web application for visualizing and managing workflow tasks through markdown file uploads and kanban-style board management.

## Technical Stack
- **Frontend**: Vue 3 (Composition API) + Tailwind CSS
- **Backend**: Go (single binary)
- **Database**: SQLite with GORM
- **Drag-and-Drop**: vuedraggable

## Core Features
1. Upload markdown files to create/update projects
2. Kanban board with 4 columns: Waiting → In Progress → Testing → Completed
3. Drag-and-drop tasks between columns
4. Edit tasks after upload
5. Multiple projects/boards support
6. Auto-save board state

## Markdown File Format
```markdown
# Project Name

## Task 1 Title
Description of the task
- Additional details
- Subtask items

## Task 2 Title
Another task description
```

- One markdown file per project
- Project name extracted from `# Header`
- Each `## Header` is a task
- Content below header is task description

## Database Schema

### Projects Table
```
id          INTEGER PRIMARY KEY
name        TEXT UNIQUE NOT NULL
created_at  DATETIME
updated_at  DATETIME
```

### Tasks Table
```
id          INTEGER PRIMARY KEY
project_id  INTEGER FOREIGN KEY
title       TEXT NOT NULL
description TEXT
status      TEXT (waiting/in_progress/testing/completed)
position    INTEGER (for ordering within column)
created_at  DATETIME
updated_at  DATETIME
```

## API Endpoints

### Projects
- `GET /api/projects` - List all projects
- `GET /api/projects/:id` - Get project details
- `POST /api/projects/upload` - Upload markdown file
- `DELETE /api/projects/:id` - Delete project

### Tasks
- `GET /api/projects/:id/tasks` - Get all tasks for a project
- `PUT /api/tasks/:id` - Update task (status, position, title, description)
- `DELETE /api/tasks/:id` - Delete task

## Development Steps

### Phase 1: Backend Foundation (Steps 1-5)

#### Step 1: Project Setup & Structure
- [ ] Create Go module (`go mod init workflow-viz`)
- [ ] Setup directory structure
- [ ] Install dependencies (gin, gorm, sqlite driver)
- [ ] Create main.go with basic server

#### Step 2: Database Setup
- [ ] Define GORM models (Project, Task)
- [ ] Create database connection
- [ ] Implement auto-migration
- [ ] Test database creation

#### Step 3: Markdown Parser Service
- [ ] Create markdown parser to extract project name
- [ ] Extract tasks (title from ##, description from content)
- [ ] Handle existing project detection
- [ ] Return structured data (project + tasks list)

#### Step 4: Upload API Handler
- [ ] Create POST /api/projects/upload endpoint
- [ ] Accept multipart/form-data file upload
- [ ] Parse markdown using service
- [ ] Check if project exists (by name)
- [ ] If exists: add new tasks, update existing
- [ ] If new: create project and all tasks
- [ ] Return project with tasks

#### Step 5: CRUD API Handlers
- [ ] GET /api/projects - list projects
- [ ] GET /api/projects/:id/tasks - get tasks by project
- [ ] PUT /api/tasks/:id - update task (title, description, status, position)
- [ ] DELETE /api/tasks/:id - delete task
- [ ] DELETE /api/projects/:id - delete project

### Phase 2: Frontend Foundation (Steps 6-8)

#### Step 6: Vue Project Setup
- [ ] Create Vue 3 project with Vite
- [ ] Install Tailwind CSS
- [ ] Install vuedraggable, axios
- [ ] Configure Tailwind
- [ ] Setup basic routing (if needed)

#### Step 7: Project List & Upload Component
- [ ] Create ProjectList component (shows all projects)
- [ ] Create FileUpload component (drag-drop or button)
- [ ] Implement upload functionality
- [ ] Show upload success/error messages
- [ ] Refresh project list after upload

#### Step 8: Kanban Board Component
- [ ] Create Board component (main container)
- [ ] Create Column component (4 columns: waiting, in_progress, testing, completed)
- [ ] Create TaskCard component (display title, description)
- [ ] Fetch tasks for selected project
- [ ] Display tasks in appropriate columns

### Phase 3: Core Functionality (Steps 9-11)

#### Step 9: Drag-and-Drop Implementation
- [ ] Integrate vuedraggable in Column component
- [ ] Enable drag between columns
- [ ] Update task status on drop
- [ ] Update task position within column
- [ ] Call API to save changes
- [ ] Show loading/success states

#### Step 10: Task Editing
- [ ] Add edit button to TaskCard
- [ ] Create inline edit mode or modal
- [ ] Edit title and description
- [ ] Save changes to API
- [ ] Update UI optimistically

#### Step 11: Build & Packaging
- [ ] Build Vue frontend (npm run build)
- [ ] Embed frontend dist in Go binary (using embed)
- [ ] Serve frontend from Go
- [ ] Test single binary execution
- [ ] Create simple run script

### Phase 4: Polish & Testing (Steps 12-13)

#### Step 12: Error Handling & Validation
- [ ] Frontend: handle API errors gracefully
- [ ] Backend: validate markdown format
- [ ] Backend: validate task updates
- [ ] Add loading states
- [ ] Add empty states (no projects, no tasks)

#### Step 13: UI/UX Improvements
- [ ] Add task count badges on columns
- [ ] Add delete confirmation dialogs
- [ ] Improve visual design (colors, spacing)
- [ ] Add keyboard shortcuts (optional)
- [ ] Test full workflow end-to-end

## Success Criteria
- [ ] Upload markdown file creates/updates project
- [ ] Tasks appear in "Waiting" column by default
- [ ] Drag-and-drop works smoothly between all columns
- [ ] Changes persist after browser refresh
- [ ] Can edit task details
- [ ] Can manage multiple projects
- [ ] Single binary runs entire application

## Future Enhancements (Out of Scope)
- Task filtering and search
- Task tags/labels
- Due dates and reminders
- Export project to markdown
- Dark mode
- Task comments/notes
- Task dependencies

## Getting Started Command
```bash
# Backend
cd backend
go run main.go

# Frontend (development)
cd frontend
npm run dev

# Production build
./build.sh
./workflow-viz
```

## Notes
- Default port: 8080
- Database location: ./data/app.db
- All tasks start in "waiting" status
- Position auto-calculated when moving tasks