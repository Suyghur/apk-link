<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-input
          v-model="queryMap.channel_alias"
          type="text"
          placeholder="渠道别名"
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
        搜索渠道
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
        @click="createChannel"
      >
        创建渠道
      </el-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'ChannelPanelHead',
  props: {
    queryMap: {
      type: Object,
      default: null
    }
  },
  methods: {
    handleFilter() {
      this.$emit('searchChannel', this.queryMap)
    },
    handleReset() {
      this.queryMap.channel_alias = ''
    },
    createChannel() {
      this.$emit('createChannel')
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

