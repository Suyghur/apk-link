<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.channel_name"
          placeholder="渠道名"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in options.channelNameOptions" :key="item" :label="item" :value="item" />
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
        搜索渠道SDK
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
        @click="createChannelSdk"
      >
        创建渠道SDK
      </el-button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'ChannelSdkPanelHead',
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
      this.$emit('searchChannelSdk', this.queryMap)
    },
    handleReset() {
      this.queryMap.channel_name = ''
    },
    createChannelSdk() {
      this.$emit('createChannelSdk')
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

