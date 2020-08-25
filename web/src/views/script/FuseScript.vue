<template>
  <div id="fuse_script">
    <TaskPanelHead @search="search" />
    <el-divider />
    <FuseScriptTable :list="list" />
  </div>

</template>

<script>
import FuseScriptTable from '@/components/tables/script/FuseScriptTable'
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import { getFuseScriptList } from '@/api/script'

export default {
  name: 'Fuse',
  components: { TaskPanelHead, FuseScriptTable },
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
      getFuseScriptList().then(response => {
        console.log(response)
        this.list = response.data.fuse_script
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
