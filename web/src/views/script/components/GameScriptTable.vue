<template>
  <div>
    <el-table
      v-loading="listLoading"
      :data="list"
      height="680"
      border
      :fit="true"
      style="width: 100%"
      element-loading-text="Loading"
      empty-text="暂无数据"
    >
      <el-table-column align="center" label="游戏组">
        <template slot-scope="scope">
          {{ scope.row.game_group }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="脚本文件名">
        <template slot-scope="scope">
          {{ scope.row.script_file_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="URL">
        <template slot-scope="scope">
          <el-link :href="'https://'+scope.row.script_file_url" type="primary">下载</el-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改日期">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ utc2local(scope.row.UpdatedAt) }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { utc2local } from '@/utils/time'

export default {
  name: 'GameScriptTable',
  props: {
    list: {
      type: Array,
      default: () => []
    },
    listLoading: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    utc2local(dataStr) {
      return utc2local(dataStr)
    }
  }
}
</script>

<style scoped>

</style>
