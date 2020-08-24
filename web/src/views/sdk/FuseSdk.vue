<template>
  <div id="fuse_sdk">
    <TaskSearchDrawer @search="search" />
    <el-divider />
    <FuseSdkTable :list="list" />
  </div>

</template>
<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import FuseSdkTable from '@/components/tables/sdk/FuseSdkTable'
import { getFuseSdkList } from '@/api/sdk'

export default {
  name: 'FuseSdk',
  components: { TaskSearchDrawer, FuseSdkTable },

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
      getFuseSdkList().then(response => {
        console.log(response)
        this.list = response.data.fuse_sdk
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
