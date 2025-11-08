<template>
  <n-card
    :title="project.name"
    size="medium"
    hoverable
    @click="handleClick"
    class="project-card"
  >
    <n-space vertical>
      <n-tag type="info">Tasks: {{ stats.total }}</n-tag>
      <n-progress
        type="line"
        :percentage="completionRate"
        status="success"
        :height="8"
        :show-indicator="false"
      />
    </n-space>
  </n-card>
</template>

<script src="./ProjectCard.script.js"></script>

<style scoped>
.project-card {
  cursor: pointer;
  transition: transform 0.2s ease;
}
.project-card:hover {
  transform: translateY(-2px);
}
</style>
