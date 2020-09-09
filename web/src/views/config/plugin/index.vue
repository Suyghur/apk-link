<template>
  <div id="game">
    <PluginPanelHead
      :query-map="queryMap"
      @searchPlugin="searchPlugin"
      @createPlugin="createPlugin"
    />
    <el-divider />
    <PluginTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="createPlugin"
    />
  </div>
</template>

<script>

import Pagination from '@/components/Pagination/index'
import { searchPlugin } from '@/api/plugin'
import PluginPanelHead from '@/views/config/plugin/components/PluginPanelHead'
import PluginTable from '@/views/config/plugin/components/PluginTable'

export default {
  name: 'Plugin',
  components: { PluginPanelHead, PluginTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        plugin_alias: undefined
      }
    }
  },
  created() {
    this.searchPlugin()
  },
  methods: {
    searchPlugin() {
      this.listLoading = true
      searchPlugin(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createPlugin() {

    }
  }
}
</script>

<style scoped>

</style>
