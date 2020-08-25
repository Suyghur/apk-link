<template>
  <div id="channel_script">
    <TaskPanelHead />
    <el-divider />
    <PluginSdkTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import { getPluginScriptList } from '@/api/script'
import PluginSdkTable from '@/components/tables/sdk/PluginSdkTable'

export default {
  name: 'PluginScript',
  components: { PluginSdkTable, TaskPanelHead },
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
