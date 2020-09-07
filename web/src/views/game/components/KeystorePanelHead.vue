<template>
  <div class="task-panel-head">

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
    <el-button
      class="filter-item"
      style="margin-left: 10px;"
      type="primary"
      icon="el-icon-search"
      @click="handleFilter"
    >
      开始筛选
    </el-button>
    <el-button
      class="filter-item"
      style="margin-left: 10px;"
      type="primary"
      icon="el-icon-edit"
      @click="handleCreate"
    >
      新增任务
    </el-button>
    <!--    <el-button type="primary" style="margin-left: 20px" @click="dialog = true">设置条件</el-button>-->
    <!--    <el-button type="primary" @click="search">开始筛选</el-button>-->
  </div>
</template>

<script>
export default {
  name: 'KeystorePanelHead',
  props: {
    list: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {}
  },

  methods: {
    handleFilter() {

    }
  }
}
</script>

