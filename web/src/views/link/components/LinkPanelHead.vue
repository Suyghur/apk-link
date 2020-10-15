<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.game_group"
          placeholder="游戏组"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.game_group=undefined"
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
      <el-col :span="6">
        <el-select
          v-model="queryMap.plugin_name"
          placeholder="插件"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
          @clear="queryMap.plugin_name=undefined"
        >
          <el-option v-for="item in options.pluginNameOptions" :key="item" :label="item" :value="item" />
        </el-select>
      </el-col>
    </el-row>

    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryMap.task_id"
          type="text"
          placeholder="任务ID"
          style="width: auto;"
          class="filter-item"
          clearable
          @clear="queryMap.task_id=undefined"
        />
      </el-col>
      <el-col :span="6">
        <el-input
          v-model="queryMap.aid"
          type="text"
          placeholder="Aid"
          style="width: auto;"
          class="filter-item"
          clearable
          @clear="queryMap.aid=undefined"
        />
      </el-col>
    </el-row>

    <div class="button_group">
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter"
      >
        搜索Link记录
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-delete"
        @click="handleReset"
      >
        重置搜索选项
      </el-button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'LinkPanelHead',
  props: {
    queryMap: {
      type: Object,
      default: null
    }
  },
  data() {
    return {}
  },
  computed: {
    ...mapGetters([
      'options'
    ])
  },

  methods: {
    handleFilter() {
      this.$emit('searchLink', this.queryMap)
    },
    handleReset() {
      this.queryMap.task_id = undefined
      this.queryMap.game_group = undefined
      this.queryMap.channel_name = undefined
      this.queryMap.plugin_name = undefined
      this.queryMap.aid = undefined
    }
  }
}
</script>
<style lang="scss" scoped>
  .keystore-panel-head {
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

