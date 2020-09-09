<template>
  <div id="channel_sdk">
    <ChannelSdkPanelHead
      :query-map="queryMap"
      @searchChannelSdk="searchChannelSdk"
      @createChannelSdk="createChannelSdk"
    />
    <el-divider />
    <ChannelSdkTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchChannelSdk"
    />
  </div>
</template>

<script>
import { searchChannelSdk } from '@/api/sdk'
import Pagination from '@/components/Pagination/index'
import ChannelSdkTable from '@/views/sdk/channel/components/ChannelSdkTable'
import ChannelSdkPanelHead from '@/views/sdk/channel/components/ChannelSdkPanelHead'

export default {
  name: 'ChannelSdk',
  components: { ChannelSdkTable, ChannelSdkPanelHead, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        channel_name: undefined
      }
    }
  },
  created() {
    this.searchChannelSdk()
  },
  methods: {
    searchChannelSdk() {
      this.listLoading = true
      searchChannelSdk(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createChannelSdk() {

    }
  }
}
</script>

<style scoped>

</style>
