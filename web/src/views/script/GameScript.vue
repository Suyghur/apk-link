<template>
  <div id="channel_script">
    <TaskSearchDrawer />
    <el-divider />
    <GameScriptTable :list="list" />
  </div>
</template>

<script>
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import { getGameScriptList } from '@/api/script'
import GameScriptTable from '@/components/tables/script/GameScriptTable'

export default {
  name: 'GameScript',
  components: { GameScriptTable, TaskSearchDrawer },
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
      getGameScriptList().then(response => {
        console.log(response)
        this.list = response.data.game_script
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
