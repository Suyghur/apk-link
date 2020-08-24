<template>
  <div id="infos">
    <TaskSearchDrawer />
    <el-divider />
    <TaskTable />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSerachDrawer'
import TaskTable from '@/components/tables/TaskTable'

export default {
  name: 'Infos',
  components: { TaskTable, TaskSearchDrawer }
}
</script>

<style scoped>

</style>
