<template>
  <div class="task-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryMap.task_id"
          type="text"
          onkeyup="value=value.replace(/^(0+)|[^\d]+/g,'')"
          placeholder="任务ID"
          style="width: auto;"
          class="filter-item"
          clearable
          @clear="queryMap.task_id=undefined"
        />
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="queryMap.game_site"
          placeholder="游戏组"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.game_site=undefined"
        >
          <el-option v-for="item in options.gameSiteOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

      <el-col :span="6">
        <el-select
          v-model="queryMap.channel_name"
          placeholder="渠道"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.channel_name=undefined"
        >
          <el-option v-for="item in options.channelNameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>

    </el-row>

    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.plugin_name"
          placeholder="插件SDK"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.plugin_name=undefined"
        >
          <el-option v-for="item in options.pluginNameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>
      <el-col :span="6">
        <el-select
          v-model="queryMap.status"
          placeholder="状态"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.status=undefined"
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
        @click="searchTask"
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
    queryMap: {
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
    searchTask() {
      this.$emit('searchTask', this.queryMap)
    },
    handleReset() {
      this.queryMap.task_id = undefined
      this.queryMap.game_group = undefined
      this.queryMap.channel_name = undefined
      this.queryMap.plugin_name = undefined
      this.queryMap.task_status = undefined
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
