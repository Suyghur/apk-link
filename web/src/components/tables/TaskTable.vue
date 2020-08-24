<template>
  <div>
    <el-table
      :data="list"
      height="800"
      border
      :fit="true"
      style="width: 100%"
      element-loading-text="Loading"
      empty-text
    >
      <el-table-column align="center" label="日期" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.create_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="任务ID">
        <template slot-scope="scope">
          {{ scope.row.task_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏名">
        <template slot-scope="scope">
          {{ scope.row.game }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="渠道">
        <template slot-scope="scope">
          {{ scope.row.channel }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="聚合SDK版本">
        <template slot-scope="scope">
          {{ scope.row.fuse_version }}
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
      <el-table-column align="center" label="状态" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter" effect="plain" size="medium">{{ scope.row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="400">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">执行</el-button>
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

export default {
  name: 'TaskTable',
  filters: {
    statusFilter(status) {
      const statusMap = {
        未执行: 'warning',
        执行中: 'gray',
        已执行: 'success',
        执行失败: 'danger'
      }
      return statusMap[status]
    }
  },
  props: {
    list: {
      type: Array,
      default: () => []
    }
  },
  methods: {
    handleEdit(index, row) {
      console.log('idnex : ', index)
    },
    handleDelete(index, row) {
      this.list.splice(index, 1)
    }
  }
}
</script>

<style scoped>

</style>
