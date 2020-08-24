<template>
  <div id="origin_bag">
    <TaskSearchDrawer />
    <el-divider />
    <OriginBagTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import OriginBagTable from '@/components/tables/game/OriginBagTable'
import { getOriginBagList } from '@/api/game'

export default {
  name: 'OriginBag',
  components: { OriginBagTable, TaskSearchDrawer },
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
