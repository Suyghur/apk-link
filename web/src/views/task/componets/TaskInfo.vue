<template>
  <div class="task-info">
    <div class="task-info-form">
      <el-form ref="taskForm" :model="taskForm" :rules="rules" class="form-container">
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏组：" prop="game_site" class="task-info-item">
              <el-select
                v-model="taskForm.game_site"
                filterable
                clearable
                placeholder="请选择"
                :disabled="disabled"
                @change="fetchGameSiteInfo"
              >
                <el-option v-for="item in options.gameSiteOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="Gid：" prop="gid" class="task-info-item">
              <el-select
                v-model="taskForm.gid"
                filterable
                clearable
                placeholder="请选择"
                :disabled="disabled"
              >
                <el-option v-for="item in gidOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏母包：" prop="game_file_name" class="task-info-item">
              <el-select
                v-model="taskForm.game_file_name"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
                @change="setOriginFileAttrs"
              >
                <el-option v-for="item in originOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏包名：" prop="game_package_name" class="task-info-item">
              <el-input
                v-model="taskForm.game_package_name"
                type="text"
                style="width: auto;"
                class="filter-item"
                :disabled="disabled"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏名：" prop="game_name" class="task-info-item">
              <el-input
                v-model="taskForm.game_name"
                type="text"
                style="width: auto;"
                :disabled="disabled"
                class="filter-item"
              />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="签名文件：" prop="keystore_name" class="task-info-item">
              <el-select
                v-model="taskForm.keystore_name"
                filterable
                clearable
                :disabled="disabled"
                placeholder="请选择"
                @change="setKeystoreAttrs"
              >
                <el-option v-for="item in keystoreOptions" :key="item" :label="item" :value="item" />
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
                @input="setGameVersionCode"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="聚合SDK版本：" prop="fuse_sdk_version" class="task-info-item">
              <el-select
                v-model="taskForm.fuse_sdk_version"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
                @change="setFuseSdkAttrs"
              >
                <el-option v-for="item in fuseVersionOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="是否白包：" class="task-info-item">
              <el-switch
                v-model="taskForm.is_white_bag"
                :active-value="1"
                :inactive-value="0"
                :disabled="disabled"
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
                @change="fetchChannelVersionOptions"
              >
                <el-option v-for="item in options.channelNameOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="渠道SDK版本：" prop="channel_sdk_version" class="task-info-item">
              <el-select
                v-model="taskForm.channel_sdk_version"
                filterable
                :disabled="disabled"
                clearable
                placeholder="请选择"
                @change="setChannelSdkAttrs"
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
                @change="fetchPluginVersionOptions"
              >
                <el-option v-for="(item) in options.pluginNameOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="插件SDK版本：" prop="plugin_version" class="task-info-item">
              <el-select
                v-model="taskForm.plugin_sdk_version"
                filterable
                :disabled="disabled"
                placeholder="请选择"
                @change="setPluginSdkAttrs"
              >
                <el-option v-for="(item) in pluginVersionOptions" :key="item" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏icon：" class="task-info-item">
              <img-uploader
                v-model="taskForm.game_icon_url"
                :disabled="disabled"
                :result-url="taskForm.game_icon_url"
              />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏闪屏：" class="task-info-item">
              <img-uploader
                v-model="taskForm.game_splash_url"
                :result-url="taskForm.game_splash_url"
                :disabled="disabled"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏logo：" class="task-info-item">
              <img-uploader
                v-model="taskForm.game_logo_url"
                :disabled="disabled"
                :result-url="taskForm.game_logo_url"
              />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label-width="200px" label="游戏登录页背景：" class="task-info-item">
              <img-uploader
                v-model="taskForm.game_login_bg_url"
                :disabled="disabled"
                :result-url="taskForm.game_login_bg_url"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <param-panel
          :game-name="taskForm.game_site"
          :channel-name="taskForm.channel_name"
          :plugin-name="taskForm.plugin_name"
        />
        <el-divider content-position="left">分发参数</el-divider>
        <el-row>
          <el-col :span="10">
            <el-form-item label-width="200px" prop="aids" label="AID：">
              <el-input
                v-model="taskForm.aids"
                type="textarea"
                :rows="3"
                placeholder="请输入Aid，多个Aid以字母逗号','分隔"
                width="auto"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </div>
    <div class="button-group">
      <el-button v-show="editable" v-loading="loading" type="success" circle icon="el-icon-check" @click="submit" />
      <el-button type="warning" circle icon="el-icon-close" @click="cancel" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getGids } from '@/api/game'
import ImgUploader from '@/components/ImgUploader'
import ParamPanel from '@/views/task/componets/param/ParamPanel'
import { getChannelSdks, getFuseSdks, getGameKeystores, getOrigins, getPluginSdks } from '@/api/options'
import { createTask, modifyTask } from '@/api/task'

const defaultForm = {
  is_white_bag: 0,
  is_plugin_sdk: 0,
  game_site: '',
  gid: '',
  aids: '',
  channel_params: '',
  plugin_params: '',
  game_name: '',
  game_package_name: '',
  game_file_name: '',
  game_file_url: '',
  game_version_code: undefined,
  game_version_name: '',
  game_orientation: undefined,
  game_icon_url: '',
  game_logo_url: '',
  game_splash_url: '',
  game_login_bg_url: '',
  fuse_sdk_file_name: '',
  fuse_sdk_version: '',
  fuse_sdk_file_url: '',
  channel_name: '',
  channel_sdk_file_name: '',
  channel_sdk_version: '',
  channel_sdk_file_url: '',
  plugin_name: '',
  plugin_sdk_file_name: '',
  plugin_sdk_version: '',
  plugin_sdk_file_url: '',
  keystore_name: '',
  keystore_password: '',
  keystore_alias: '',
  keystore_alias_password: '',
  keystore_file_url: '',
  status_code: 1000,
  status_msg: '未执行',
  script_file_name: '',
  script_file_url: '',
  ext: ''
}
export default {
  name: 'TaskInfo',
  components: { ParamPanel, ImgUploader },
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    taskForm: {
      type: Object,
      default: () => defaultForm
    },
    editable: {
      type: Boolean,
      default: false
    },
    isCreate: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      originOptions: [],
      fuseVersionOptions: [],
      channelVersionOptions: [],
      pluginVersionOptions: [],
      keystoreOptions: [],
      gidOptions: [],
      originInfoMap: {},
      keystoreInfoMap: {},
      fuseSdkInfoMap: {},
      channelSdkInfoMap: {},
      pluginSdkInfoMap: {},
      defaultImgSrc: 'https://sdkfile.hihoulang.com/logo.png',
      loading: false,
      rules: {
        game_site: [{ required: true, message: '请选择游戏组', trigger: 'change' }],
        gid: [{ required: true, message: '请选择游戏ID', trigger: 'change' }],
        game_file_name: [{ required: true, message: '请选择母包', trigger: 'change' }],
        game_package_name: [{ required: true, message: '请输入游戏包名', trigger: 'blur' }],
        game_name: [{ required: true, message: '请输入游戏名', trigger: 'blur' }],
        game_version_name: [{ required: true, message: '请输入游戏版本名', trigger: 'blur' }],
        game_version_code: [{ required: true, message: '请输入游戏版本号', trigger: 'blur' }],
        fuse_sdk_version: [{ required: true, message: '请选择聚合SDK版本', trigger: 'change' }],
        channel_name: [{ required: true, message: '请选择渠道SDK', trigger: 'change' }],
        channel_sdk_version: [{ required: true, message: '请选择渠道SDK版本', trigger: 'change' }],
        keystore_name: [{ required: true, message: '请选择签名文件', trigger: 'change' }],
        aids: [{ required: true, message: '请输入Aid', trigger: 'blur' }]
      }
    }
  },
  computed: {
    ...mapGetters([
      'options'
    ])
  },
  watch: {
    'taskForm.aids': function(val) {
      // this.handleAids(val)
    }
  },
  created() {
    this.fetchFuseVersionOptions()
  },
  mounted() {
    // this.handleAids()
  },

  methods: {
    fetchGameSiteInfo(gameSite) {
      this.fetchGidOptions(gameSite)
      this.fetchOrigins(gameSite)
      this.fetchKeystoreOptions(gameSite)
    },
    fetchGidOptions(gameSite) {
      this.taskForm.gid = ''
      this.gidOptions = []
      return new Promise((resolve, reject) => {
        getGids({ game_site: gameSite }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          this.gidOptions = response.data
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    fetchOrigins(gameSite) {
      this.taskForm.game_file_name = ''
      this.originOptions = []
      return new Promise((resolve, reject) => {
        getOrigins({ game_site: gameSite }).then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          const origins = response.data.list
          for (let i = 0; i < origins.length; i++) {
            this.originOptions.push(origins[i].game_file_name)
            this.originInfoMap[origins[i].game_file_name] = {
              'game_orientation': origins[i].game_orientation,
              'apk_url': origins[i].apk_url
            }
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    setOriginFileAttrs(item) {
      this.taskForm.game_orientation = this.originInfoMap[item].game_orientation
      this.taskForm.game_file_url = this.originInfoMap[item].apk_url
    },
    setGameVersionCode(code) {
      this.taskForm.game_version_code = parseInt(code)
    },
    fetchKeystoreOptions(gameSite) {
      this.taskForm.keystore_name = ''
      this.keystoreOptions = []
      if (gameSite !== '') {
        return new Promise((resolve, reject) => {
          getGameKeystores({ game_site: gameSite }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            const keystores = response.data
            for (let i = 0; i < keystores.length; i++) {
              this.keystoreOptions.push(keystores[i].keystore_name)
              this.keystoreInfoMap[keystores[i].keystore_name] = {
                'keystore_password': keystores[i].keystore_password,
                'keystore_alias': keystores[i].keystore_alias,
                'keystore_alias_password': keystores[i].keystore_alias_password,
                'keystore_file_url': keystores[i].keystore_file_url
              }
            }
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    setKeystoreAttrs(item) {
      this.taskForm.keystore_password = this.keystoreInfoMap[item].keystore_password
      this.taskForm.keystore_alias = this.keystoreInfoMap[item].keystore_alias
      this.taskForm.keystore_alias_password = this.keystoreInfoMap[item].keystore_alias_password
      this.taskForm.keystore_file_url = this.keystoreInfoMap[item].keystore_file_url
    },
    fetchFuseVersionOptions() {
      this.fuseVersionOptions = []
      this.taskForm.fuse_sdk_version = ''
      return new Promise((resolve, reject) => {
        getFuseSdks().then(response => {
          if (!response) {
            return reject('数据加载异常')
          }
          const fuseSdks = response.data.list
          for (let i = 0; i < fuseSdks.length; i++) {
            this.fuseVersionOptions.push(fuseSdks[i].sdk_version)
            this.fuseSdkInfoMap[fuseSdks[i].sdk_version] = {
              'sdk_file_name': fuseSdks[i].sdk_file_name,
              'sdk_file_url': fuseSdks[i].sdk_file_url
            }
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    setFuseSdkAttrs(item) {
      this.taskForm.fuse_sdk_file_name = this.fuseSdkInfoMap[item].sdk_file_name
      this.taskForm.fuse_sdk_file_url = this.fuseSdkInfoMap[item].sdk_file_url
    },
    fetchChannelVersionOptions(channelName) {
      this.channelVersionOptions = []
      this.taskForm.channel_sdk_version = ''
      if (channelName !== '') {
        return new Promise((resolve, reject) => {
          getChannelSdks({ channel_name: channelName }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            const channelSdks = response.data
            for (let i = 0; i < channelSdks.length; i++) {
              this.channelVersionOptions.push(channelSdks[i].sdk_version)
              this.channelSdkInfoMap[channelSdks[i].sdk_version] = {
                'sdk_file_name': channelSdks[i].sdk_file_name,
                'sdk_file_url': channelSdks[i].sdk_file_url
              }
            }
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    setChannelSdkAttrs(item) {
      this.taskForm.channel_sdk_file_name = this.channelSdkInfoMap[item].sdk_file_name
      this.taskForm.channel_sdk_file_url = this.channelSdkInfoMap[item].sdk_file_url
    },
    fetchPluginVersionOptions(pluginName) {
      this.pluginVersionOptions = []
      this.taskForm.plugin_sdk_version = ''
      if (pluginName !== '') {
        return new Promise((resolve, reject) => {
          getPluginSdks({ plugin_name: pluginName }).then(response => {
            if (!response) {
              return reject('数据加载异常')
            }
            const pluginSdks = response.data.plugin_sdks
            for (let i = 0; i < pluginSdks.length; i++) {
              this.pluginVersionOptions.push(pluginSdks[i].sdk_version)
              this.pluginSdkInfoMap[pluginSdks[i].sdk_version] = {
                'sdk_file_name': pluginSdks[i].sdk_file_name,
                'sdk_file_url': pluginSdks[i].sdk_file_url
              }
            }
            resolve(response)
          }).catch(error => {
            reject(error)
          })
        })
      }
    },
    setPluginSdkAttrs(item) {
      this.taskForm.plugin_sdk_file_name = this.pluginSdkInfoMap[item].sdk_file_name
      this.taskForm.plugin_sdk_file_url = this.pluginSdkInfoMap[item].sdk_file_url
    },
    submit() {
      this.$refs.taskForm.validate(valid => {
        if (valid) {
          return new Promise((resolve, reject) => {
            this.taskForm.aids = this.tempAids.toString()
            if (!this.isCreate) {
              modifyTask(this.taskForm).then(response => {
                if (!response) {
                  this.$notify({
                    title: '失败',
                    message: '编辑打包任务失败',
                    type: 'error',
                    duration: 3000
                  })
                  return reject('编辑打包任务失败')
                }
                this.loading = true
                this.$notify({
                  title: '成功',
                  message: '编辑打包任务成功',
                  type: 'success',
                  duration: 3000
                })
                this.loading = false
                this.$router.replace({ path: '/task/infos' })
                resolve(response)
              }).catch(error => {
                reject(error)
              })
            } else {
              createTask(this.taskForm).then(response => {
                if (!response) {
                  this.$notify({
                    title: '失败',
                    message: '创建打包任务失败',
                    type: 'error',
                    duration: 3000
                  })
                  return reject('创建打包任务失败')
                }
                this.loading = true
                this.$notify({
                  title: '成功',
                  message: '创建打包任务成功',
                  type: 'success',
                  duration: 3000
                })
                this.loading = false
                this.$router.replace({ path: '/task/infos' })
                resolve(response)
              }).catch(error => {
                reject(error)
              })
            }
          }
          )
        } else {
          this.$notify({
            title: '失败',
            message: '创建或编辑打包任务失败，存在部分参数未填写，请检查',
            type: 'error',
            duration: 3000
          })
          return false
        }
      }
      )
    },
    cancel() {
      this.$router.replace({ path: '/task/index' })
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
