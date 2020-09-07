<template>
  <div id="task-detail">
    <TaskInfo :disabled="true" :task-form="taskForm" :editable="false" />
  </div>
</template>

<script>
import TaskInfo from '@/views/task/componets/TaskInfo'
import { getTaskInfo } from '@/api/task'

export default {
  name: 'TaskDetail',
  components: { TaskInfo },
  data() {
    return {
      taskForm: {}
    }
  },
  created() {
    this.fetchTaskInfo()
  },
  methods: {
    fetchTaskInfo() {
      return new Promise((resolve, reject) => {
        getTaskInfo({ task_id: parseInt(this.$route.params.id) }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.taskForm = response.data.task_info
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}
</script>

<style scoped>

</style>
