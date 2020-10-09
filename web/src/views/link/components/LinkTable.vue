<template>
  <div>
    <el-table
      v-loading="listLoading"
      :data="list"
      height="600"
      border
      :fit="true"
      style="width: 100%"
      element-loading-text="Loading"
      empty-text="暂无数据"
    >
      <el-table-column align="center" label="任务ID">
        <template slot-scope="scope">
          {{ scope.row.task_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="AID">
        <template slot-scope="scope">
          {{ scope.row.aid }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏组">
        <template slot-scope="scope">
          {{ scope.row.game_group }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏名">
        <template slot-scope="scope">
          {{ scope.row.game_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏版本">
        <template slot-scope="scope">
          {{ scope.row.game_version_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="聚合SDK版本">
        <template slot-scope="scope">
          {{ scope.row.fuse_sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="渠道">
        <template slot-scope="scope">
          {{ scope.row.channel_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="插件">
        <template slot-scope="scope">
          {{ scope.row.plugin_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="URL" width="80">
        <template slot-scope="scope">
          <el-link :href="'https://'+scope.row.link_url" type="primary">下载</el-link>
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
  name: 'LinkTable',
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
