<template>
  <div>
    <el-button type="primary" @click="dialog = true">设置条件</el-button>
    <el-button type="primary">开始筛选</el-button>
    <el-drawer
      ref="drawer"
      :before-close="handleClose"
      :visible.sync="dialog"
      direction="rtl"
      custom-class="demo-drawer"
    >
      <div class="demo-drawer__content" style="text-align: left">
        <el-form :model="form" :inline="true">
          <el-form-item label="任务id：" :label-width="formLabelWidth">
            <el-input v-model="form.task_id" autocomplete="off" />
          </el-form-item>
          <el-form-item label="aid：" :label-width="formLabelWidth">
            <el-input v-model="form.aid" autocomplete="off" />
          </el-form-item>
          <el-form-item label="游戏：" :label-width="formLabelWidth">
            <el-select v-model="form.game" placeholder="请选择游戏">
              <el-option label="测试游戏一" value="shanghai" />
              <el-option label="测试游戏二" value="beijing" />
            </el-select>
          </el-form-item>
          <el-form-item label="渠道：" :label-width="formLabelWidth">
            <el-select v-model="form.channel" placeholder="请选择渠道">
              <el-option label="后浪(houlang)" value="shanghai" />
              <el-option label="应用宝(yyb)" value="beijing" />
              <el-option label="华为(huawei)" value="beijing" />
              <el-option label="Oppo(oppo)" value="beijing" />
              <el-option label="Vivo(vivo)" value="beijing" />
              <el-option label="百度多酷(duoku)" value="beijing" />
              <el-option label="小米(mi)" value="beijing" />
              <el-option label="U九游(uc)" value="beijing" />
            </el-select>
          </el-form-item>
          <el-form-item label="媒体：" :label-width="formLabelWidth">
            <el-select v-model="form.media" placeholder="请选择媒体">
              <el-option label="头条" value="shanghai" />
              <el-option label="广点通" value="shanghai" />
              <el-option label="快手" value="shanghai" />
              <el-option label="百度cps" value="shanghai" />
              <el-option label="穿山甲" value="beijing" />
            </el-select>
          </el-form-item>
          <el-form-item label="状态：" :label-width="formLabelWidth">
            <el-select v-model="form.status" placeholder="请选择打包状态">
              <el-option label="未执行" value="shanghai" />
              <el-option label="执行中" value="shanghai" />
              <el-option label="已执行" value="beijing" />
              <el-option label="执行失败" value="beijing" />
            </el-select>
          </el-form-item>
        </el-form>
        <div class="demo-drawer__footer" style="text-align: center;">
          <el-button size="large" @click="cancelForm">取 消</el-button>
          <el-button
            type="primary"
            size="large"
            :loading="loading"
            @click="$refs.drawer.closeDrawer()"
          >
            {{ loading ? '提交中 ...' : '确定' }}
          </el-button>
        </div>
      </div>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'TaskSearchDrawer',
  data() {
    return {
      dialog: false,
      loading: false,
      timer: null,
      formLabelWidth: '100px',
      form: {
        task_id: '',
        aid: '',
        game: '',
        channel: '',
        media: '',
        status: ''
      }
    }
  },
  methods: {
    handleClose(done) {
      if (this.loading) {
        return
      }
      this.$confirm('确定提交筛选条件吗？')
        .then(_ => {
          this.loading = true
          this.timer = setTimeout(() => {
            done()
            setTimeout(() => {
              this.loading = false
            }, 400)
          }, 2000)
        })
        .catch(_ => {
        })
    },
    cancelForm() {
      this.loading = false
      this.dialog = false
      clearTimeout(this.timer)
    }
  }
}
</script>

