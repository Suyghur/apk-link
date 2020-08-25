<template>
  <div class="task-panel-head">

    <el-input
      v-model="list.task_id"
      placeholder="任务ID"
      style="width: 200px;"
      class="filter-item"
      @keyup.enter.native="handleFilter"
    />
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
  name: 'TaskPanelHead',
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

