<template>
  <div id="channel_script">
    <TaskPanelHead />
    <el-divider />
    <ChannelScriptTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import ChannelScriptTable from '@/components/tables/script/ChannelScriptTable'
import { getChannelScriptList } from '@/api/script'

export default {
  name: 'ChannelScript',
  components: { TaskPanelHead, ChannelScriptTable },
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
      getChannelScriptList().then(response => {
        console.log(response)
        this.list = response.data.channel_script
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
