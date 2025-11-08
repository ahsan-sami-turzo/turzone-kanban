import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import Upload from '../views/Upload.vue'
import Project from '../views/Project.vue'

const routes = [
  { path: '/', component: Dashboard },
  { path: '/upload', component: Upload },
  { path: '/project/:id', component: Project },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
