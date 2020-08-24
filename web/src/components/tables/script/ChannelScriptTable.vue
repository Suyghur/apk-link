<template>
  <div>
    <el-table
      :data="list"
      height="800"
      border
      highlight-current-row
      element-loading-text="Loading"
    >
      <el-table-column align="center" label="渠道">
        <template slot-scope="scope">
          {{ scope.row.channel }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="脚本文件名">
        <template slot-scope="scope">
          {{ scope.row.file_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改日期">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.modify_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改人">
        <template slot-scope="scope">
          <el-tag :type="scope.row.modify_person" effect="plain" size="medium">{{ scope.row.modify_person }}
          </el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'ChannelScriptTable',
  props: {
    list: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style scoped>

</style>
