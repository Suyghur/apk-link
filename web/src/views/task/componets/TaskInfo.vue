<template>
  <div class="task-info">
    <div class="task-info-form">
      <el-form ref="taskForm" :model="taskForm" :rules="rules" class="form-container">
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏组：" prop="game_group" class="task-info-item">
              <el-select
                v-model="taskForm.game_group"
                filterable
                clearable
                placeholder="请选择"
                :disabled="disabled"
                @change="fetchGameGroupInfo"
              >
                <el-option v-for="item in selectOptions.gameGroupOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label-width="200px" label="GID：" class="task-info-item">
              <el-input
                v-model="taskForm.gid"
                type="text"
                readonly
                disabled
                style="width: auto;"
                class="filter-item"
                @keyup.enter.native="handleFilter"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏母包：" prop="origin_bag" class="task-info-item">
              <el-select
                v-model="taskForm.origin_bag"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
              >
                <el-option v-for="item in originBagOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="聚合SDK版本：" prop="fuse_version" class="task-info-item">
              <el-select
                v-model="taskForm.fuse_version"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
              >
                <el-option v-for="item in fuseVersionOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏版本名：" prop="game_version_name" class="task-info-item">
              <el-input
                v-model="taskForm.game_version_name"
                type="text"
                :disabled="disabled"
                style="width: auto;"
                class="filter-item"
                @keyup.enter.native="handleFilter"
              />
            </el-form-item>
          </el-col>

          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏版本号：" prop="game_version_code" class="task-info-item">
              <el-input
                v-model="taskForm.game_version_code"
                type="text"
                :disabled="disabled"
                style="width: auto;"
                class="filter-item"
                @keyup.enter.native="handleFilter"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="渠道SDK：" prop="channel_name" class="task-info-item">
              <el-select
                v-model="taskForm.channel_name"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
                @change="fetchChannelVersionList"
              >
                <el-option v-for="item in selectOptions.channelNamelOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="渠道SDK版本：" prop="channel_version" class="task-info-item">
              <el-select
                v-model="taskForm.channel_version"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
              >
                <el-option v-for="item in channelVersionOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="插件SDK：" prop="plugin_name" class="task-info-item">
              <el-select
                v-model="taskForm.plugin_name"
                filterable
                clearable
                :disabled="disabled"
                placeholder="请选择"
                @change="fetchPluginVersionList"
              >
                <el-option v-for="(item) in selectOptions.pluginNameOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="插件SDK版本：" prop="plugin_version" class="task-info-item">
              <el-select
                v-model="taskForm.plugin_version"
                filterable
                :disabled="disabled"
                placeholder="请选择"
              >
                <el-option v-for="(item) in pluginVersionOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏icon：" class="task-info-item">
              <img-uploader v-model="taskForm.icon_url" :disabled="disabled" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏闪屏：" class="task-info-item">
              <img-uploader v-model="taskForm.splash_url" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏logo：" class="task-info-item">
              <img-uploader v-model="taskForm.logo_url" :disabled="disabled" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏登录页背景：" class="task-info-item">
              <img-uploader v-model="taskForm.bg_url" :disabled="disabled" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="签名文件：" prop="keystore_name" class="task-info-item">
              <el-select
                v-model="taskForm.keystore_name"
                filterable
                clearable
                :disabled="disabled"
                placeholder="请选择"
              >
                <el-option v-for="item in gameKeystoreOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <param-panel
          :game-name="taskForm.game_group"
          :channel-name="taskForm.channel_name"
          :plugin-name="taskForm.plugin_name"
        />
        <el-divider content-position="left">分发参数</el-divider>
        <el-row>
          <el-form-item label-width="200px" label="AID：">
            <el-transfer
              v-model="taskForm.aids"
              :disable="disabled"
              filterable
              filter-placeholder="请输入AID"
              :data="aidOptions"
            />
          </el-form-item>
        </el-row>
      </el-form>
    </div>
    <div class="button-group">
      <el-button v-loading="loading" type="success" circle icon="el-icon-check" @click="submit" />
      <el-button type="warning" circle icon="el-icon-close" @click="cancel" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getAidList, getGameGroupInfo, getGameKeystore, getOriginBagList } from '@/api/game'
import { getChannelVersionList, getFuseVersionList, getPluginVersionList } from '@/api/sdk'
import ImgUploader from '@/components/ImgUploader'
import ParamPanel from '@/views/task/componets/param/ParamPanel'

const defaultForm = {
  gid: '',
  game_group: '',
  origin_bag: '',
  fuse_version: '',
  game_version_name: '',
  game_version_code: undefined,
  channel_name: '',
  channel_version: '',
  plugin_name: '',
  plugin_version: '',
  icon_url: '',
  splash_url: '',
  logo_url: '',
  bg_url: '',
  keystore_name: '',
  aids: []
}
export default {
  name: 'TaskInfo',
  components: { ParamPanel, ImgUploader },
  props: {
    disabled: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      taskForm: Object.assign({}, defaultForm),
      originBagOptions: [],
      fuseVersionOptions: [],
      channelVersionOptions: [],
      pluginVersionOptions: [],
      gameKeystoreOptions: [],
      aidOptions: [],
      defaultImgSrc: 'https://sdkfile.hihoulang.com/logo.png',
      loading: false,
      rules: {
        game_group: [{ required: true, message: '请选择游戏组', trigger: 'change' }],
        origin_bag: [{ required: true, message: '请选择母包', trigger: 'change' }],
        fuse_version: [{ required: true, message: '请选择聚合SDK版本', trigger: 'change' }],
        game_version_name: [{ required: true, message: '请输入游戏版本名', trigger: 'blur' }],
        game_version_code: [{ required: true, message: '请输入游戏版本号', trigger: 'blur' }],
        channel_name: [{ required: true, message: '请选择渠道SDK', trigger: 'change' }],
        channel_version: [{ required: true, message: '请选择渠道SDK版本', trigger: 'change' }],
        keystore_name: [{ required: true, message: '请选择签名文件', trigger: 'change' }]
      }
    }
  },
  computed: {
    ...
    mapGetters([
      'selectOptions'
    ])
  },
  created() {
    this.fetchFuseVersionList()
  },

  methods: {
    fetchGameGroupInfo(gameGroup) {
      this.fetchOriginBagList(gameGroup)
      this.fetchGameKeystoreList(gameGroup)
      this.fetchAidList(gameGroup)
      return new Promise((resolve, reject) => {
        getGameGroupInfo({ game_group: gameGroup }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.taskForm.gid = response.data.gid
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    fetchOriginBagList(gameGroup) {
      this.taskForm.origin_bag = ''
      this.originBagOptions = []
      return new Promise((resolve, reject) => {
        getOriginBagList({ game_group: gameGroup }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.originBagOptions = response.data.bag
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    fetchGameKeystoreList(gameGroup) {
      this.taskForm.keystore_name = ''
      this.gameKeystoreOptions = []
      if (gameGroup !== '') {
        return new Promise((resolve, reject) => {
          getGameKeystore({ game_group: gameGroup }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            this.gameKeystoreOptions = response.data.keystore
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    fetchAidList(gameGroup) {
      this.taskForm.aids = []
      this.aidOptions = []
      if (gameGroup !== '') {
        return new Promise((resolve, reject) => {
          getAidList({ game_group: gameGroup }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            const data = []
            response.data.aids.forEach((aid) => {
              data.push({
                label: aid,
                key: aid
              })
            })
            this.aidOptions = data
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    fetchFuseVersionList() {
      this.fuseVersionOptions = []
      this.taskForm.fuse_version = ''

      return new Promise((resolve, reject) => {
        getFuseVersionList().then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.fuseVersionOptions = response.data.version_list
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    fetchChannelVersionList(channelName) {
      this.channelVersionOptions = []
      this.taskForm.channel_version = ''
      if (channelName !== '') {
        return new Promise((resolve, reject) => {
          getChannelVersionList({ channel_name: channelName }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            this.channelVersionOptions = response.data.version_list
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    fetchPluginVersionList(pluginName) {
      this.pluginVersionOptions = []
      this.taskForm.plugin_version = ''
      if (pluginName !== '') {
        return new Promise((resolve, reject) => {
          getPluginVersionList({ plugin_name: pluginName }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            this.pluginVersionOptions = response.data.version_list
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    submit() {
      this.$refs.taskForm.validate(valid => {
        if (this.taskForm.aids.length <= 0) {
          this.$notify({
            title: '失败',
            message: '创建打包任务失败，aid不能为空，请检查',
            type: 'error',
            duration: 3000
          })
          // this.loading = false
          return false
        }
        if (valid) {
          this.loading = true
          this.$notify({
            title: '成功',
            message: '创建打包任务成功',
            type: 'success',
            duration: 3000
          })
          console.log(this.taskForm)
          this.loading = false
          this.$router.replace({ path: '/task/infos' })
          return true
        } else {
          this.$notify({
            title: '失败',
            message: '创建打包任务失败，存在部分参数未填写，请检查',
            type: 'error',
            duration: 3000
          })
          // this.loading = false
          return false
        }
      })
    },
    cancel() {
      this.$router.replace({ path: '/task/infos' })
    }

  }
}
</script>

<style lang="scss" scoped>
  @import "~@/styles/mixin.scss";

  .task-info {
    position: relative;

    .task-info-form {
      padding: 40px 45px 20px 50px;
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .task-info-item {
        float: left;
      }
    }

    .button-group {
      text-align: center;
      margin-bottom: 20px;
      /*position: fixed;*/
    }
  }

</style>
