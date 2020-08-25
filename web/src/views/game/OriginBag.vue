<template>
  <div id="origin_bag">
    <TaskPanelHead />
    <el-divider />
    <OriginBagTable :list="list" />
  </div>
</template>

<script>
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import OriginBagTable from '@/components/tables/game/OriginBagTable'
import { getOriginBagList } from '@/api/game'

export default {
  name: 'OriginBag',
  components: { OriginBagTable, TaskPanelHead },
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
      getOriginBagList().then(response => {
        console.log(response)
        this.list = response.data.origin_bag
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
