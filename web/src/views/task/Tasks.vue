<template>
  <div id="tasks">
    <TaskPanelHead :select-options="options" :query-list="queryList" @searchTask="fetchTasks" />
    <el-divider />
    <TaskTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryList.page"
      :limit.sync="queryList.limit"
      @pagination="fetchTasks"
    />
  </div>
</template>

<script>
import TaskPanelHead from '@/views/task/componets/TaskPanelHead'
import TaskTable from '@/views/task/componets/TaskTable'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

import { getSelectOptions, getTaskList } from '@/api/task'

export default {
  name: 'Tasks',
  components: { TaskTable, TaskPanelHead, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryList: {
        page: 1,
        limit: 20,
        task_id: undefined,
        game_group: undefined,
        channel_name: undefined,
        plugin_name: undefined,
        task_status: undefined
      },
      options: {
        gameGroupOptions: [],
        channelOptions: [],
        pluginOptions: [],
        statusOptions: []
      }
    }
  },
  created() {
    this.fetchTasks()
    this.fetchSelectOptions()
  },
  methods: {
    fetchTasks() {
      this.listLoading = true
      getTaskList(this.queryList).then(response => {
        this.list = response.data.tasks
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },

    fetchSelectOptions() {
      getSelectOptions().then(response => {
        this.options.gameGroupOptions = response.data.game_group
        this.options.channelOptions = response.data.channel
        this.options.pluginOptions = response.data.plugin
        this.options.statusOptions = response.data.status
      })
    },
    search() {
      this.list = []
      this.fetchTasks()
    }
  }
}
</script>

<style scoped>

</style>
