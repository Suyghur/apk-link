<template>
  <div>
    <el-table
      v-loading="listLoading"
      :data="list"
      height="680"
      border
      :fit="true"
      style="width: 100%"
      element-loading-text="Loading"
      empty-text="暂无数据"
    >
      <el-table-column align="center" label="游戏组" width="150">
        <template slot-scope="scope">
          {{ scope.row.game_group }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="签名文件名" width="150">
        <template slot-scope="scope">
          {{ scope.row.keystore_name }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="签名文件密码" width="150">
        <template slot-scope="scope">
          {{ scope.row.keystore_password }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="签名文件别名" width="150">
        <template slot-scope="scope">
          {{ scope.row.keystore_alias }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="签名文件别名密码" width="150">
        <template slot-scope="scope">
          {{ scope.row.keystore_alias_password }}
        </template>
      </el-table-column>

      <el-table-column align="center" label="MD5指纹密码" width="300">
        <template slot-scope="scope">
          <el-tag :type="scope.row.fingerprints_md5" effect="plain" size="medium">{{ scope.row.fingerprints_md5 }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="SHA1指纹密码" width="400">
        <template slot-scope="scope">
          <el-tag :type="scope.row.fingerprints_sha1" effect="plain" size="medium">{{ scope.row.fingerprints_sha1 }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="SHA256指纹密码" width="600">
        <template slot-scope="scope">
          <el-tag :type="scope.row.fingerprints_sha256" effect="plain" size="medium">{{ scope.row.fingerprints_sha256 }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="URL" width="80">
        <template slot-scope="scope">
          <el-link :href="'https://'+scope.row.keystore_file_url" type="primary">下载</el-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="最后修改日期" width="150">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ utc2local(scope.row.UpdatedAt) }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { utc2local } from '@/utils/time'

export default {
  name: 'KeystoreTable',
  props: {
    list: {
      type: Array,
      default: () => []
    },
    listLoading: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    utc2local(dataStr) {
      return utc2local(dataStr)
    }
  }
}
</script>

<style scoped>

</style>
