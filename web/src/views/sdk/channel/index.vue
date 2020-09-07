<template>
  <div id="channel_id">
    <TaskPanelHead @search="search" />
    <el-divider />
    <ChannelSdkTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import ChannelSdkTable from '@/components/tables/sdk/ChannelSdkTable'
import { getChannelSdkList } from '@/api/sdk'

export default {
  name: 'ChannelSdk',
  components: { TaskPanelHead, ChannelSdkTable },
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
      getChannelSdkList().then(response => {
        console.log(response)
        this.list = response.data.channel_sdk
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
