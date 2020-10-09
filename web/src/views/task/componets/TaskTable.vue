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
          <router-link v-if="scope.row.status==='未执行'" :to="'/task/edit/'+scope.row.task_id" class="link-type">
            <span>{{ scope.row.task_id }}</span>
          </router-link>
          <router-link v-else :to="'/task/detail/'+scope.row.task_id" class="link-type">
            <span>{{ scope.row.task_id }}</span>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="游戏组">
        <template slot-scope="scope">
          {{ scope.row.game_group }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="渠道">
        <template slot-scope="scope">
          {{ scope.row.channel_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="聚合SDK版本">
        <template slot-scope="scope">
          {{ scope.row.fuse_sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="渠道SDK版本">
        <template slot-scope="scope">
          {{ scope.row.channel_sdk_version }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="插件SDK">
        <template slot-scope="scope">
          {{ scope.row.plugin_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="状态" width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status_msg | statusFilter" effect="plain" size="medium">{{ scope.row.status_msg }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改日期" width="200">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ utc2local(scope.row.UpdatedAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            v-if="scope.row.status_msg==='未执行'"
            size="mini"
            @click="handleExecute(scope.row,'执行中')"
          >执行
          </el-button>
          <el-popconfirm
            v-if="scope.row.status_msg==='执行中'"
            confirm-button-text="确定"
            cancel-button-text="点错了"
            icon="el-icon-info"
            icon-color="red"
            title="确定取消该任务吗？"
            @onConfirm="handleExecute(scope.row,'未执行')"
          >
            <el-button slot="reference" size="mini">取消</el-button>
          </el-popconfirm>

          <router-link :to="'/task/detail/'+scope.row.task_id">
            <el-button
              v-if="scope.row.status_msg==='成功'"
              size="mini"
              style="margin-left: 5px"
            >
              查看
            </el-button>
          </router-link>
          <router-link :to="'/task/edit/'+scope.row.task_id">
            <el-button
              v-if="scope.row.status_msg==='未执行'"
              size="mini"
              style="margin-left: 5px"
            >编辑
            </el-button>
          </router-link>
          <el-popconfirm
            v-if="scope.row.status_msg!=='成功'"
            confirm-button-text="确定"
            cancel-button-text="点错了"
            icon="el-icon-info"
            icon-color="red"
            title="确定删除该任务吗？"
            @onConfirm="handleDelete(scope.$index, scope.row)"
          >
            <el-button slot="reference" size="mini" type="danger" style="margin-left: 5px">
              删除
            </el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

import { deleteTask, modifyStatus } from '@/api/task'
import { utc2local } from '@/utils/time'

export default {
  name: 'TaskTable',
  filters: {
    statusFilter(status_msg) {
      const statusMap = {
        未执行: 'warning',
        执行中: 'gray',
        成功: 'success',
        失败: 'danger'
      }
      return statusMap[status_msg]
    }
  },
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
    handleExecute(row, status) {
      let status_code
      switch (status) {
        case '执行中':
          status_code = 1001
          break
        case '未执行':
          status_code = 1000
          break
      }
      return new Promise((resolve, reject) => {
        modifyStatus({ task_id: parseInt(row.task_id), status_code: status_code }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.$message({
            message: response.msg,
            type: 'success'
          })
          row.status_msg = status
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    handleDelete(index, row) {
      return new Promise((resolve, reject) => {
        deleteTask({ task_id: parseInt(row.task_id) }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.popconfirm
          this.$notify({
            title: 'Success',
            message: response.msg,
            type: 'success',
            duration: 1000
          })
          this.list.splice(index, 1)
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    utc2local(dataStr) {
      return utc2local(dataStr)
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
