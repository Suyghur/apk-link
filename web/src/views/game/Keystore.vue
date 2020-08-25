<template>
  <div id="keystore">
    <TaskPanelHead />
    <el-divider />
    <KeystoreTable :list="list" />
  </div>
</template>

<script>
import { getKeystoreList } from '@/api/game'
import TaskPanelHead from '@/components/drawer/TaskPanelHead'
import KeystoreTable from '@/components/tables/game/KeystoreTable'

export default {
  name: 'KeyStore',
  components: { KeystoreTable, TaskPanelHead },
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
