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
      <el-table-column align="center" label="插件名">
        <template slot-scope="scope">
          {{ scope.row.sdk_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="插件别名">
        <template slot-scope="scope">
          {{ scope.row.sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="SDK文件名">
        <template slot-scope="scope">
          {{ scope.row.sdk_file_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="URL">
        <template slot-scope="scope">
          <el-link :href="'https://'+scope.row.sdk_file_url" type="primary">下载</el-link>
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
  name: 'PluginSdkTable',
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
