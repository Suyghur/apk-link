<template>
  <div id="keystore">
    <KeystorePanelHead
      :query-map="queryMap"
      @searchKeystore="searchKeystore"
      @createKeystore="createKeystore"
    />
    <el-divider />
    <KeystoreTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchKeystore"
    />
  </div>
</template>

<script>
import KeystoreTable from '@/views/keystore/components/KeystoreTable'
import { searchKeystore } from '@/api/keystore'
import KeystorePanelHead from '@/views/keystore/components/KeystorePanelHead'
import Pagination from '@/components/Pagination/index'

export default {
  name: 'KeyStore',
  components: { KeystorePanelHead, KeystoreTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        game_site: undefined
      }
    }
  },
  created() {
    this.searchKeystore()
  },
  methods: {
    searchKeystore() {
      this.listLoading = true
      searchKeystore(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
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
