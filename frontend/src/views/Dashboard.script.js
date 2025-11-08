import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import ProjectCard from '../components/ProjectCard.vue'

export default {
  components: { ProjectCard },
  setup() {
    const router = useRouter()
    const projects = ref([])
    const projectStats = ref({})

    const goToUpload = () => router.push('/upload')
    const viewProject = (id) => router.push(`/project/${id}`)

    const loadProjects = async () => {
      try {
        const res = await axios.get('/api/projects')
        projects.value = res.data

        for (const project of projects.value) {
          await loadStats(project.id)
        }
      } catch (err) {
        console.error('Failed to load projects:', err)
      }
    }

    const loadStats = async (projectId) => {
      try {
        const res = await axios.get(`/api/projects/${projectId}/tasks`)
        const tasks = res.data
        projectStats.value[projectId] = {
          total: tasks.length,
          completed: tasks.filter(t => t.status === 'completed').length
        }
      } catch (err) {
        console.error(`Failed to load stats for project ${projectId}:`, err)
      }
    }

    onMounted(loadProjects)

    return { projects, projectStats, goToUpload, viewProject }
  }
}
