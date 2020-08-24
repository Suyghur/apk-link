<template>
  <div>
    <el-table
      :data="list"
      height="800"
      border
      fit
      highlight-current-row
      element-loading-text="Loading"
    >
      <el-table-column align="center" label="游戏icon">
        <template slot-scope="scope">
          {{ scope.row.icon }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏">
        <template slot-scope="scope">
          {{ scope.row.game }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏版本">
        <template slot-scope="scope">
          {{ scope.row.game_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="聚合SDK版本">
        <template slot-scope="scope">
          {{ scope.row.fuse_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="文件名">
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
      <el-table-column align="center" label="URL" width="500">
        <template slot-scope="scope">
          <a href="http://www.hihoulang.com" target="view_window"> {{ scope.row.url }}</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'OriginBagTable',
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
