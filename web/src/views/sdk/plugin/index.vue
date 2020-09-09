<template>
  <div id="plugin_sdk">
    <PluginSdkPanelHead
      :query-map="queryMap"
      @searchPluginSdk="searchPluginSdk"
      @createPluginSdk="createPluginSdk"
    />
    <el-divider />
    <PluginSdkTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchPluginSdk"
    />
  </div>
</template>

<script>
import { searchPluginSdk } from '@/api/sdk'
import PluginSdkPanelHead from '@/views/sdk/plugin/components/PluginSdkPanelHead'
import PluginSdkTable from '@/views/sdk/plugin/components/PluginSdkTable'
import Pagination from '@/components/Pagination/index'

export default {
  name: 'PluginSdk',
  components: { PluginSdkPanelHead, PluginSdkTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        plugin_name: undefined
      }
    }
  },
  created() {
    this.searchPluginSdk()
  },
  methods: {
    searchPluginSdk() {
      this.listLoading = true
      searchPluginSdk(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createPluginSdk() {
    }
  }
}
</script>

<style scoped>

</style>
