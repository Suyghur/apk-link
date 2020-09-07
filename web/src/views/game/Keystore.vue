<template>
  <div id="keystore">
    <KeystorePanelHead />
    <el-divider />
    <KeystoreTable :list="list" />
  </div>
</template>

<script>
import KeystoreTable from '@/components/tables/game/KeystoreTable'
import KeystorePanelHead from '@/views/game/components/KeystorePanelHead'
import { searchKeystore } from '@/api/keystore'

export default {
  name: 'KeyStore',
  components: { KeystorePanelHead, KeystoreTable },
  data() {
    return {
      list: null,
      listLoading: true,
      queryMap: {
        page: 1,
        page_size: 20,
        game_group: undefined
      }
    }
  },
  created() {
    this.searchKeystore()
  },
  methods: {
    searchKeystore() {
      searchKeystore().then(response => {
        console.log(response)
        this.list = response.data.keystore
        this.listLoading = false
      })
    },
    search() {
      this.list = []
      this.fetchKeystoreList()
    }
  }
}
</script>

<style scoped>

</style>
