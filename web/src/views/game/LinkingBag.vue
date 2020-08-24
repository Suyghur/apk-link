<template>
  <div id="linking_bag">
    <TaskSearchDrawer />
    <el-divider />
    <LinkingBagTable :list="list" />
  </div>
</template>

<script>
import { getLinkingBag } from '@/api/game'
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import LinkingBagTable from '@/components/tables/game/LinkingBagTable'

export default {
  name: 'LinkingBag',
  components: { LinkingBagTable, TaskSearchDrawer },
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
      getLinkingBag().then(response => {
        console.log(response)
        this.list = response.data.linking_bag
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
