<template>
  <div id="keystore">
    <TaskSearchDrawer />
    <el-divider />
    <KeystoreTable :list="list" />
  </div>
</template>

<script>
import { getKeystoreList } from '@/api/game'
import TaskSearchDrawer from '@/components/drawer/TaskSearchDrawer'
import KeystoreTable from '@/components/tables/game/KeystoreTable'

export default {
  name: 'KeyStore',
  components: { KeystoreTable, TaskSearchDrawer },
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
      getKeystoreList().then(response => {
        console.log(response)
        this.list = response.data.keystore
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
