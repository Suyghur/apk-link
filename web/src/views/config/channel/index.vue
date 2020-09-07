<template>
  <div id="keystore">
    <KeystorePanelHead
      :query-map="queryMap"
      @searchKeystore="searchKeystore"
      @createKeystore="createKeystore"
    />
    <el-divider />
    <KeystoreTable :list="list" />
  </div>
</template>

<script>
import KeystoreTable from '@/components/tables/game/KeystoreTable'
import { searchKeystore } from '@/api/keystore'
import KeystorePanelHead from '@/views/keystore/components/KeystorePanelHead'

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
    createKeystore() {

    },
    search() {
      this.list = []
      this.searchKeystore()
    }
  }
}
</script>

<style scoped>

</style>
