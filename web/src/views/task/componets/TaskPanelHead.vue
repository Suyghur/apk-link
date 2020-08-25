<template>
  <div class="task-panel-head">

    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryList.task_id"
          type="text"
          onkeyup="value=value.replace(/^(0+)|[^\d]+/g,'')"
          placeholder="任务ID"
          style="width: auto;"
          class="filter-item"
          clearable
          @keyup.enter.native="handleFilter"
        />
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="queryList.game_group"
          placeholder="游戏组名"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.gameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

      <el-col :span="6">
        <el-select
          v-model="queryList.channel"
          placeholder="渠道"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.gameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

    </el-row>

    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryList.plugin"
          placeholder="插件SDK"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.pluginOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="queryList.status"
          placeholder="状态"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.statusOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>
    </el-row>
    <div class="button_group">
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索任务
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate"
      >
        创建任务
      </el-button>
    </div>

    <!--    <el-button type="primary" style="margin-left: 20px" @click="dialog = true">设置条件</el-button>-->
    <!--    <el-button type="primary" @click="search">开始筛选</el-button>-->
  </div>
</template>

<script>
export default {
  name: 'TaskPanelHead',
  props: {
    options: {
      type: Object,
      default: () => []
    }
  },
  data() {
    return {
      queryList: {
        task_id: undefined,
        game_group: undefined,
        channel: undefined,
        status: undefined
      }
      // gameOptions: [],
      // channelOptions: [],
      // pluginOptions: [],
      // statusOptions: []

    }
  },
  created() {
    // this.fetchSelectOptions()
  },
  methods: {
    // fetchSelectOptions() {
    //   getSelectOptions().then(response => {
    //     console.log(response)
    //
    //     this.gameOptions = response.data.game_group
    //     this.channelOptions = response.data.channel
    //     this.pluginOptions = response.data.plugin
    //     this.statusOptions = response.data.status
    //     this.listLoading = false
    //   })
    // },
    handleFilter() {
      console.log(this.queryList.task_id)
    },
    handleCreate() {
      this.$router.push({ path: '/task/create' })
    }
  }
}
</script>

<style lang="scss" scoped>
  .task-panel-head {
    padding: 20px 20px 0;
  }

  .el-row {
    margin-bottom: 10px;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .button_group {
    margin-top: 20px;
  }
</style>
