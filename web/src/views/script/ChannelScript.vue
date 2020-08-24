<template>
  <div id="channel_script">
    <TaskSearchDrawer />
    <el-divider />
    <ChannelScriptTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import ChannelScriptTable from '@/components/tables/script/ChannelScriptTable'
import { getChannelScriptList } from '@/api/script'

export default {
  name: 'ChannelScript',
  components: { TaskSearchDrawer, ChannelScriptTable },
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
