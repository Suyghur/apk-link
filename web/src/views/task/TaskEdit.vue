<template>
  <div id="task-edit">
    <TaskInfo :disabled="false" />
  </div>
</template>

<script>

import TaskInfo from '@/views/task/componets/TaskInfo'

export default {
  name: 'TaskEdit',
  components: { TaskInfo }
}
</script>

