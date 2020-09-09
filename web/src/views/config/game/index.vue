<template>
  <div id="game">
    <GamePanelHead
      :query-map="queryMap"
      @searchGame="searchGame"
      @createGame="createGame"
    />
    <el-divider />
    <GameTable :list="list" :list-loading="listLoading" />
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryMap.page"
      :limit.sync="queryMap.page_size"
      @pagination="createGame"
    />
  </div>
</template>

<script>

import Pagination from '@/components/Pagination/index'
import GamePanelHead from '@/views/config/game/components/GamePanelHead'
import GameTable from '@/views/config/game/components/GameTable'
import { searchGame } from '@/api/game'

export default {
  name: 'Game',
  components: { GamePanelHead, GameTable, Pagination },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      queryMap: {
        page: 1,
        page_size: 20,
        gid: undefined
      }
    }
  },
  created() {
    this.searchGame()
  },
  methods: {
    searchGame() {
      this.listLoading = true
      searchGame(this.queryMap).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        setTimeout(() => {
          this.listLoading = false
        }, 1000)
      })
    },
    createGame() {

    }
  }
}
</script>

<style scoped>

</style>
