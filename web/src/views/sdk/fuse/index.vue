<template>
  <div id="fuse_sdk">
    <FuseSdkPanelHead
      :query-map="queryMap"
      @searchFuseSdk="searchFuseSdk"
      @createFuseSdk="createFuseSdk"
    />
    <el-divider />
    <FuseSdkTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchFuseSdk"
    />
  </div>

</template>
<script>
import { searchFuseSdk } from '@/api/sdk'
import FuseSdkTable from '@/views/sdk/fuse/components/FuseSdkTable'
import FuseSdkPanelHead from '@/views/sdk/fuse/components/FuseSdkPanelHead'
import Pagination from '@/components/Pagination/index'

export default {
  name: 'FuseSdk',
  components: { FuseSdkTable, FuseSdkPanelHead, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        sdk_version: undefined
      }
    }
  },
  created() {
    this.searchFuseSdk()
  },
  methods: {
    searchFuseSdk() {
      this.listLoading = true
      searchFuseSdk(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createFuseSdk() {

    }
  }
}
</script>

<style scoped>

</style>
