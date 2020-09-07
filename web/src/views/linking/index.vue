<template>
  <div id="linking_bag">
    <TaskPanelHead />
    <el-divider />
    <LinkingBagTable :list="list" />
  </div>
</template>

<script>
import { getLinkingBag } from '@/api/game'
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import LinkingBagTable from '@/components/tables/game/LinkingBagTable'

export default {
  name: 'Linking',
  components: { LinkingBagTable, TaskPanelHead },
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
