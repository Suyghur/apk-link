<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryMap.gid"
          type="text"
          onkeyup="value=value.replace(/^(0+)|[^\d]+/g,'')"
          placeholder="游戏组ID"
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
        搜索游戏组
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-delete"
        @click="handleReset"
      >
        重置搜索项
      </el-button>
      <el-button
        class="filter-item"
        type="primary"
        icon="el-icon-edit"
        @click="createGame"
      >
        创建游戏组
      </el-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'GamePanelHead',
  props: {
    queryMap: {
      type: Object,
      default: null
    }
  },
  methods: {
    handleFilter() {
      this.$emit('searchGame', this.queryMap)
    },
    handleReset() {
      this.queryMap.gid = ''
    },
    createGame() {
      this.$emit('createGame')
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

