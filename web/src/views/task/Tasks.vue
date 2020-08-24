<template>
  <div id="tasks">
    <TaskSearchDrawer @search="search" />
    <el-divider />
    <TaskTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import TaskTable from '@/components/tables/TaskTable'
import { getTaskList } from '@/api/task'

export default {
  name: 'Tasks',
  components: { TaskTable, TaskSearchDrawer },
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
