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
          <el-option v-for="item in options.gameGroupOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

      <el-col :span="6">
        <el-select
          v-model="queryList.channel_name"
          placeholder="渠道"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.channelNameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

    </el-row>

    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryList.plugin_name"
          placeholder="插件SDK"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.pluginNameOptions" :key="item" :label="item" :value="item" />
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
          <el-option v-for="item in options.taskStatusOptions" :key="item" :label="item" :value="item" />
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
        icon="el-icon-delete"
        @click="handleReset"
      >
        重置选项
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
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'TaskPanelHead',
  props: {
    queryList: {
      type: Object,
      default: null
    }
  },
  computed: {
    ...mapGetters([
      'options'
    ])
  },
  methods: {
    handleFilter() {
      this.queryList.task_id = parseInt(this.queryList.task_id)
      this.$emit('searchTask', this.queryList)
    },
    handleReset() {
      this.queryList.task_id = undefined
      this.queryList.game_group = undefined
      this.queryList.channel_name = undefined
      this.queryList.plugin_name = undefined
      this.queryList.task_status = undefined
    },
    handleCreate() {
      this.$emit('createTask')
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
