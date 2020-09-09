<template>
  <div id="linking_bag">
    <LinkPanelHead
      :query-map="queryMap"
      @searchLink="searchLink"
    />
    <el-divider />
    <LinkTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchLink"
    />
  </div>
</template>

<script>
import Pagination from '@/components/Pagination'
import LinkPanelHead from '@/views/link/components/LinkPanelHead'
import LinkTable from '@/views/link/components/LinkTable'
import { searchLink } from '@/api/link'

export default {
  name: 'Link',
  components: { LinkPanelHead, LinkTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        task_id: undefined,
        game_group: undefined,
        aid: undefined,
        fuse_sdk_version: undefined,
        channel_name: undefined,
        plugin_name: undefined
      }
    }
  },
  created() {
    this.searchLink()
  },
  methods: {
    searchLink() {
      this.listLoading = true
      searchLink(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    }
  }
}
</script>

<style scoped>

</style>
