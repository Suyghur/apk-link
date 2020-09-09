<template>
  <div id="game">
    <ChannelPanelHead
      :query-map="queryMap"
      @searchChannel="searchChannel"
      @createChannel="createChannel"
    />
    <el-divider />
    <ChannelTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="createChannel"
    />
  </div>
</template>

<script>

import Pagination from '@/components/Pagination/index'
import ChannelPanelHead from '@/views/config/channel/components/ChannelPanelHead'
import ChannelTable from '@/views/config/channel/components/ChannelTable'
import { searchChannel } from '@/api/channel'

export default {
  name: 'Channel',
  components: { ChannelPanelHead, ChannelTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        channel_alias: undefined
      }
    }
  },
  created() {
    this.searchChannel()
  },
  methods: {
    searchChannel() {
      this.listLoading = true
      searchChannel(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createChannel() {

    }
  }
}
</script>

<style scoped>

</style>
