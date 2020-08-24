<template>
  <div id="channel_script">
    <TaskSearchDrawer />
    <el-divider />
    <PluginSdkTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import { getPluginScriptList } from '@/api/script'
import PluginSdkTable from '@/components/tables/sdk/PluginSdkTable'

export default {
  name: 'PluginScript',
  components: { PluginSdkTable, TaskSearchDrawer },
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
      getPluginScriptList().then(response => {
        console.log(response)
        this.list = response.data.plugin_script
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
