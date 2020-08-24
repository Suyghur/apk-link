<template>
  <div id="plugin_sdk">
    <TaskSearchDrawer @search="search" />
    <el-divider />
    <PluginSdkTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import PluginSdkTable from '@/components/tables/sdk/PluginSdkTable'
import { getPluginSdkList } from '@/api/sdk'

export default {
  name: 'PluginSdk',
  components: { TaskSearchDrawer, PluginSdkTable },
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
      getPluginSdkList().then(response => {
        console.log(response)
        this.list = response.data.plugin_sdk
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
