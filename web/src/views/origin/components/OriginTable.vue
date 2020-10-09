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
      <el-table-column align="center" label="游戏组">
        <template slot-scope="scope">
          {{ scope.row.game_group }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="Gid">
        <template slot-scope="scope">
          {{ scope.row.gid }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="是否接入聚合SDK">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.is_fuse_sdk===1" effect="plain" size="medium">是
          </el-tag>
          <el-tag v-if="scope.row.is_fuse_sdk===0" effect="plain" size="medium">否
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="聚合SDK版本">
        <template slot-scope="scope">
          {{ scope.row.sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="母包文件名">
        <template slot-scope="scope">
          {{ scope.row.game_file_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="母包版本号">
        <template slot-scope="scope">
          {{ scope.row.game_version_code }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="母包版本名">
        <template slot-scope="scope">
          {{ scope.row.game_version_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="URL">
        <template slot-scope="scope">
          <el-link :href="'https://'+scope.row.apk_url" type="primary">下载</el-link>
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
  name: 'OriginTable',
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
