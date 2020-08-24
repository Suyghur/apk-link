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
      <el-table-column align="center" label="任务ID">
        <template slot-scope="scope">
          {{ scope.row.task_id }}
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
      <el-table-column align="center" label="渠道SDK">
        <template slot-scope="scope">
          {{ scope.row.channel }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="渠道SDK版本">
        <template slot-scope="scope">
          {{ scope.row.channel_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="插件SDK">
        <template slot-scope="scope">
          {{ scope.row.plugin }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="AID">
        <template slot-scope="scope">
          {{ scope.row.aid }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改日期">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.modify_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="分发申请人">
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
  name: 'LinkingBagTable',
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
