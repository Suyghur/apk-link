<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.game_site"
          placeholder="游戏组"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.gameSiteOptions" :key="item" :label="item" :value="item" />
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
        搜索脚本文件
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
        @click="createKeystore"
      >
        创建脚本文件
      </el-button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'GameScriptPanelHead',
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
      this.$emit('searchScript', this.queryMap)
    },
    handleReset() {
      this.queryMap.game_site = ''
    },
    createKeystore() {
      this.$emit('createScript')
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

