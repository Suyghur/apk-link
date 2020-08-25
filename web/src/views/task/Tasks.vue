<template>
  <div id="tasks">
    <TaskPanelHead :options="options" />
    <el-divider />
    <TaskTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/views/task/componets/TaskPanelHead'
import TaskTable from '@/views/task/componets/TaskTable'
import { getSelectOptions, getTaskList } from '@/api/task'

export default {
  name: 'Tasks',
  components: { TaskTable, TaskPanelHead },
  data() {
    return {
      list: null,
      listLoading: true,
      options: {
        gameOptions: [],
        channelOptions: [],
        pluginOptions: [],
        statusOptions: []
      }
    }
  },
  created() {
    this.fetchData()
    this.fetchSelectOptions()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getTaskList().then(response => {
        this.list = response.data.tasks
        this.listLoading = false
      })
    },

    fetchSelectOptions() {
      getSelectOptions().then(response => {
        this.options.gameOptions = response.data.game_group
        this.options.channelOptions = response.data.channel
        this.options.pluginOptions = response.data.plugin
        this.options.statusOptions = response.data.status
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
