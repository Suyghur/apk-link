<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.plugin_name"
          placeholder="插件名"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.pluginNameOptions" :key="item" :label="item" :value="item" />
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
        搜索插件SDK
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-delete"
        @click="handleReset"
      >
        重置搜索选项
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-edit"
        @click="createPluginSdk"
      >
        创建插件SDK
      </el-button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'PluginSdkPanelHead',
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
    handleFilter() {
      this.$emit('searchPluginSdk', this.queryMap)
    },
    handleReset() {
      this.queryMap.plugin_name = ''
    },
    createPluginSdk() {
      this.$emit('createPluginSdk')
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

