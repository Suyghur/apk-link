<template>
  <div>
    <el-table
      :data="taskListData"
      height="250"
      border
      :fit="true"
      style="width: 100%;text-align: center"
    >
      <el-table-column
        align="center"
        prop="date"
        label="日期"
      />
      <el-table-column
        align="center"
        prop="taskId"
        label="任务id"
      />
      <el-table-column
        align="center"
        prop="game"
        label="游戏"
      />
      <el-table-column
        align="center"
        prop="aid"
        label="aid"
      />
      <el-table-column
        align="center"
        prop="channel"
        label="渠道"
      />
      <el-table-column
        align="center"
        prop="media"
        label="媒体"
      />
      <el-table-column
        align="center"
        prop="status"
        label="状态"
      />
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'TaskTable',
  data() {
    return {
      taskListData: [
        {
          date: '2020-08-21',
          taskId: '123',
          game: '测试游戏1',
          aid: 'abc123456',
          channel: '后浪',
          media: '百度cps',
          status: '正在执行'
        },
        {
          date: '2020-08-21',
          taskId: '123',
          game: '测试游戏1',
          aid: 'abc123456',
          channel: '后浪',
          media: '百度cps',
          status: '正在执行'
        },
        {
          date: '2020-08-21',
          taskId: '123',
          game: '测试游戏1',
          aid: 'abc123456',
          channel: '后浪',
          media: '百度cps',
          status: '正在执行'
        },
        {
          date: '2020-08-21',
          taskId: '123',
          game: '测试游戏1',
          aid: 'abc123456',
          channel: '后浪',
          media: '百度cps',
          status: '正在执行'
        }
      ]
    }
  }
}
</script>

<style scoped>

</style>
