<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryMap.plugin_alias"
          type="text"
          placeholder="插件别名"
          style="width: auto;"
          class="filter-item"
          clearable
          @keyup.enter.native="handleFilter"
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
        搜索插件
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-delete"
        @click="handleReset"
      >
        重置搜索
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-edit"
        @click="createPlugin"
      >
        创建插件
      </el-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'PluginPanelHead',
  props: {
    queryMap: {
      type: Object,
      default: null
    }
  },
  methods: {
    handleFilter() {
      this.$emit('searchPlugin', this.queryMap)
    },
    handleReset() {
      this.queryMap.plugin_alias = ''
    },
    createPlugin() {
      this.$emit('createPlugin')
    }
  }
}
</script>
<style lang="scss" scoped>
  .keystore-panel-head {
    padding: 20px 20px 0;
  }

  .button_group {
    margin-top: 20px;
  }
</style>

