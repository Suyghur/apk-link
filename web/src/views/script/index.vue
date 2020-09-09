<template>
  <div id="channel_script">
    <GameScriptPanelHead
      :query-map="queryMap"
      @searchScript="searchScript"
      @createScript="createScript"
    />
    <el-divider />
    <GameScriptTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="searchScript"
    />
  </div>
</template>

<script>
import { searchScript } from '@/api/script'
import GameScriptTable from '@/views/script/components/GameScriptTable'
import GameScriptPanelHead from '@/views/script/components/GameScriptPanelHead'
import Pagination from '@/components/Pagination/index'

export default {
  name: 'GameScript',
  components: { GameScriptTable, GameScriptPanelHead, Pagination },
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
    this.searchScript()
  },
  methods: {
    searchScript() {
      this.listLoading = true
      searchScript(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createScript() {

    }
  }
}
</script>

<style scoped>

</style>
