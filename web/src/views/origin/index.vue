<template>
  <div id="origin_bag">
    <OriginPanelHead
      :query-map="queryMap"
      @searchOrigin="searchOrigin"
      @createOrigin="createOrigin"
    />
    <el-divider />
    <OriginTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchOrigin"
    />
  </div>
</template>

<script>

import OriginTable from '@/views/origin/components/OriginTable'
import OriginPanelHead from '@/views/origin/components/OriginPanelHead'
import { searchOrigin } from '@/api/origin'
import Pagination from '@/components/Pagination/index'

export default {
  name: 'Origin',
  components: { OriginTable, OriginPanelHead, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        game_group: undefined
      }
    }
  },
  created() {
    this.searchOrigin()
  },
  methods: {
    searchOrigin() {
      this.listLoading = true
      searchOrigin(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createOrigin() {
    }
  }
}
</script>

<style scoped>

</style>
