<template>
  <div id="tasks">
    <TaskPanelHead
      :query-list="queryList"
      @searchTask="fetchTasks"
      @createTask="createTask"
    />
    <el-divider />
    <TaskTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryList.page"
      :limit.sync="queryList.page_size"
      @pagination="fetchTasks"
    />
  </div>
</template>

<script>
import TaskPanelHead from '@/views/task/componets/TaskPanelHead'
import TaskTable from '@/views/task/componets/TaskTable'
import Pagination from '@/components/Pagination'
import { getTaskList } from '@/api/task'

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
        page_size: 20,
        task_id: undefined,
        game_group: undefined,
        channel_name: undefined,
        plugin_name: undefined,
        status_code: -1
      }
    }
  },
  created() {
    this.fetchTasks()
    this.$store.dispatch('options/fetchOptions')
  },
  methods: {
    fetchTasks() {
      this.listLoading = true
      getTaskList(this.queryList).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createTask() {
      this.$router.push({ path: '/task/create' })
    }
  }
}
</script>

<style scoped>

</style>
