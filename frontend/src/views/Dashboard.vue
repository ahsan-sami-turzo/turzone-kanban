<template>
  <n-layout>
    <!-- Navbar -->
    <n-layout-header bordered class="navbar">
      <n-space justify="space-between" align="center">
        <n-h2>📁 Project Dashboard</n-h2>
        <n-button type="primary" @click="goToUpload">+ New Project</n-button>
      </n-space>
    </n-layout-header>

    <!-- Main Content -->
    <n-layout-content class="dashboard-content">
      <n-space vertical size="large">
        <n-grid :cols="4" :x-gap="24" :y-gap="24">
          <n-grid-item v-for="project in projects" :key="project.id">
            <ProjectCard
              :project="project"
              :stats="projectStats[project.id] || { total: 0, completed: 0 }"
              @click="viewProject(project.id)"
            />
          </n-grid-item>
        </n-grid>
      </n-space>
    </n-layout-content>
  </n-layout>
</template>

<script src="./Dashboard.script.js"></script>

<style scoped>
.navbar {
  padding: 16px 24px;
  background-color: #f9f9f9;
}
.dashboard-content {
  padding: 32px;
}
</style>
