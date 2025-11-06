import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../views/Dashboard.vue';
import ProjectView from '../views/ProjectView.vue';

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
  },
  {
    path: '/project/:id',
    name: 'ProjectView',
    component: ProjectView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;