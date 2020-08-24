<template>
  <div>
    <el-table
      :data="list"
      height="800"
      border
      highlight-current-row
      element-loading-text="Loading"
    >
      <el-table-column align="center" label="SDK版本">
        <template slot-scope="scope">
          {{ scope.row.sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="SDK文件名">
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
  name: 'FuseSdkTable',
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
