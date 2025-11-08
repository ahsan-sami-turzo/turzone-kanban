<template>
  <div class="min-h-screen flex flex-col bg-gray-50 text-gray-900 antialiased">
    
    <header class="flex items-center justify-between px-6 sm:px-10 py-4 bg-white shadow-md z-10">
      <h1 class="text-2xl font-bold tracking-tight text-gray-800">
        <span class="text-indigo-600">Project</span> Kanban
      </h1>
      <nav class="flex gap-6 text-sm font-medium text-gray-500">
        <button class="hover:text-indigo-600 transition duration-200 ease-in-out">Dashboard</button>
        <button class="hover:text-indigo-600 transition duration-200 ease-in-out">Projects</button>
        <button class="hover:text-indigo-600 transition duration-200 ease-in-out">Settings</button>
      </nav>
    </header>

    <main class="flex-1 overflow-y-auto px-4 sm:px-8 lg:px-12 py-8">
      <RouterView />
    </main>

    <footer class="mt-auto py-4 text-center text-xs text-gray-400 border-t border-gray-100 bg-white">
      © {{ new Date().getFullYear() }} Project Kanban — All rights reserved.
    </footer>
  </div>
</template>

<script setup>
  import { RouterView } from 'vue-router'
</script>

<style>
  html {
    font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  }
</style>
