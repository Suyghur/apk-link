<template>
  <div id="tasks">
    <TaskPanelHead :list="list" />
    <el-divider />
    <TaskTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/components/head/TaskPanelHead'
import TaskTable from '@/components/tables/TaskTable'
import { getTaskList } from '@/api/task'

export default {
  name: 'Tasks',
  components: { TaskTable, TaskPanelHead },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getTaskList().then(response => {
        console.log(response)
        this.list = response.data.tasks
        this.listLoading = false
      })
    },
    search() {
      this.list = []
      this.fetchData()
    }
  }
}
</script>

<style scoped>

</style>
