export default {
  props: {
    project: Object,
    stats: Object
  },
  emits: ['click'],
  computed: {
    completionRate() {
      const { completed, total } = this.stats
      return total > 0 ? Math.round((completed / total) * 100) : 0
    }
  },
  methods: {
    handleClick() {
      this.$emit('click')
    }
  }
}
