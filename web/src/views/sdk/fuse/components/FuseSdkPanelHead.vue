<template>
  <div class="keystore-panel-head">
    <el-row :gutter="10">
      <el-col :span="6">
        <el-select
          v-model="queryMap.sdk_version"
          placeholder="聚合SDK版本"
          clearable
          style="width: auto"
          class="filter-item"
          filterable
        >
          <el-option v-for="item in fuseVersionOptions" :key="item" :label="item" :value="item" />
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
        搜索聚合SDK
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
        @click="createFuseSdk"
      >
        创建聚合SDK
      </el-button>
    </div>
  </div>
</template>

<script>
import { getFuseSdks } from '@/api/options'

export default {
  name: 'FuseSdkPanelHead',
  props: {
    queryMap: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      fuseVersionOptions: []
    }
  },
  created() {
    this.fetchSdkVersion()
  },
  methods: {
    fetchSdkVersion() {
      this.fuseVersionOptions = []
      this.queryMap.sdk_version = ''
      return new Promise((resolve, reject) => {
        getFuseSdks().then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          const fuseSdks = response.data
          for (let i = 0; i < fuseSdks.length; i++) {
            this.fuseVersionOptions.push(fuseSdks[i].sdk_version)
            // this.fuseSdkInfoMap[fuseSdks[i].sdk_version] = {
            //   'sdk_file_name': fuseSdks[i].sdk_file_name,
            //   'sdk_file_url': fuseSdks[i].sdk_file_url
            // }
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },

    handleFilter() {
      this.$emit('searchFuseSdk', this.queryMap)
    },
    handleReset() {
      this.queryMap.sdk_version = ''
    },
    createFuseSdk() {
      this.$emit('createFuseSdk')
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

